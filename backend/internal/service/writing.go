package service

import (
  "context"
  "encoding/json"
  "encoding/xml"
  "errors"
  "fmt"
  "io"
  "net/http"
  "net/url"
  "os"
  "path/filepath"
  "strings"
  "sync"
  "time"
)

type WritingRequest struct {
  Title          string `json:"title"`
  Topic          string `json:"topic"`
  Locale         string `json:"locale"`
  PubmedQuery    string `json:"pubmedQuery"`
  IncludeFilters bool   `json:"includeFilters"`
}

type WritingResult struct {
  ReviewID    string `json:"reviewId"`
  Status      string `json:"status"`
  DownloadURL string `json:"downloadUrl"`
  Message     string `json:"message"`
}

type writingTask struct {
  ID           string
  Status       string
  Message      string
  DownloadPath string
  UpdatedAt    time.Time
}

type writingTaskStore struct {
  mu    sync.RWMutex
  tasks map[string]*writingTask
}

func newWritingTaskStore() *writingTaskStore {
  return &writingTaskStore{tasks: make(map[string]*writingTask)}
}

func (store *writingTaskStore) save(task *writingTask) {
  store.mu.Lock()
  defer store.mu.Unlock()
  store.tasks[task.ID] = task
}

func (store *writingTaskStore) get(id string) (*writingTask, bool) {
  store.mu.RLock()
  defer store.mu.RUnlock()
  task, ok := store.tasks[id]
  return task, ok
}

func (service *ReviewService) StartWritingWorkflow(ctx context.Context, payload WritingRequest) (*WritingResult, error) {
  if strings.TrimSpace(payload.Title) == "" || strings.TrimSpace(payload.Topic) == "" {
    return nil, errors.New("title and topic are required")
  }

  taskID := fmt.Sprintf("review-%d", time.Now().UnixNano())
  task := &writingTask{
    ID:        taskID,
    Status:    "processing",
    Message:   "writing workflow started",
    UpdatedAt: time.Now(),
  }
  service.tasks.save(task)

  query := strings.TrimSpace(payload.PubmedQuery)
  if query == "" {
    query = payload.Topic
  }

  articles, err := service.fetchPubMedArticles(ctx, query, 50)
  if err != nil {
    task.Status = "failed"
    task.Message = err.Error()
    service.tasks.save(task)
    return service.taskToResult(task), nil
  }

  docxPath, err := service.buildDocx(taskID, payload, articles)
  if err != nil {
    task.Status = "failed"
    task.Message = err.Error()
    service.tasks.save(task)
    return service.taskToResult(task), nil
  }

  task.Status = "completed"
  task.Message = "review document generated"
  task.DownloadPath = docxPath
  task.UpdatedAt = time.Now()
  service.tasks.save(task)

  return service.taskToResult(task), nil
}

func (service *ReviewService) GetWritingStatus(id string) (*WritingResult, error) {
  task, ok := service.tasks.get(id)
  if !ok {
    return nil, errors.New("writing task not found")
  }
  return service.taskToResult(task), nil
}

func (service *ReviewService) GetWritingDocumentPath(id string) (string, error) {
  task, ok := service.tasks.get(id)
  if !ok {
    return "", errors.New("writing task not found")
  }
  if task.DownloadPath == "" {
    return "", errors.New("document not available")
  }
  return task.DownloadPath, nil
}

func (service *ReviewService) taskToResult(task *writingTask) *WritingResult {
  downloadURL := ""
  if task.DownloadPath != "" {
    downloadURL = fmt.Sprintf("/api/writing/%s/download", task.ID)
  }
  return &WritingResult{
    ReviewID:    task.ID,
    Status:      task.Status,
    DownloadURL: downloadURL,
    Message:     task.Message,
  }
}

type esearchResponse struct {
  ESearchResult struct {
    IDList []string `json:"idlist"`
  } `json:"esearchresult"`
}

type pubmedArticleSet struct {
  Articles []pubmedArticle `xml:"PubmedArticle"`
}

type pubmedArticle struct {
  MedlineCitation struct {
    PMID    string `xml:"PMID"`
    Article struct {
      ArticleTitle string `xml:"ArticleTitle"`
      Abstract     struct {
        AbstractText []string `xml:"AbstractText"`
      } `xml:"Abstract"`
      AuthorList struct {
        Authors []pubmedAuthor `xml:"Author"`
      } `xml:"AuthorList"`
      Journal struct {
        Title        string `xml:"Title"`
        JournalIssue struct {
          PubDate struct {
            Year  string `xml:"Year"`
            Month string `xml:"Month"`
            Day   string `xml:"Day"`
          } `xml:"PubDate"`
        } `xml:"JournalIssue"`
      } `xml:"Journal"`
    } `xml:"Article"`
  } `xml:"MedlineCitation"`
}

type pubmedAuthor struct {
  ForeName string `xml:"ForeName"`
  LastName string `xml:"LastName"`
}

type articleInfo struct {
  PMID    string
  Title   string
  Abstract string
  Journal string
  Authors string
  PubDate string
}

func (service *ReviewService) fetchPubMedArticles(ctx context.Context, query string, limit int) ([]articleInfo, error) {
  if strings.TrimSpace(query) == "" {
    return nil, errors.New("pubmed query is empty")
  }

  searchURL := fmt.Sprintf("%s/esearch.fcgi", strings.TrimSuffix(service.pubmedBaseURL, "/"))
  form := url.Values{}
  form.Set("db", "pubmed")
  form.Set("retmode", "json")
  form.Set("retmax", fmt.Sprintf("%d", limit))
  form.Set("term", query)
  if service.pubmedAPIKey != "" {
    form.Set("api_key", service.pubmedAPIKey)
  }

  req, err := http.NewRequestWithContext(ctx, http.MethodPost, searchURL, strings.NewReader(form.Encode()))
  if err != nil {
    return nil, err
  }
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  client := &http.Client{Timeout: 30 * time.Second}
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("pubmed esearch failed: %s", resp.Status)
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  var searchResp esearchResponse
  if err := json.Unmarshal(body, &searchResp); err != nil {
    return nil, err
  }

  if len(searchResp.ESearchResult.IDList) == 0 {
    return nil, errors.New("no pubmed articles found")
  }

  fetchURL := fmt.Sprintf("%s/efetch.fcgi", strings.TrimSuffix(service.pubmedBaseURL, "/"))
  fetchQuery := url.Values{}
  fetchQuery.Set("db", "pubmed")
  fetchQuery.Set("retmode", "xml")
  fetchQuery.Set("rettype", "abstract")
  fetchQuery.Set("id", strings.Join(searchResp.ESearchResult.IDList, ","))
  if service.pubmedAPIKey != "" {
    fetchQuery.Set("api_key", service.pubmedAPIKey)
  }

  fetchReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fetchURL+"?"+fetchQuery.Encode(), nil)
  if err != nil {
    return nil, err
  }
  fetchResp, err := client.Do(fetchReq)
  if err != nil {
    return nil, err
  }
  defer fetchResp.Body.Close()

  if fetchResp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("pubmed efetch failed: %s", fetchResp.Status)
  }

  fetchBody, err := io.ReadAll(fetchResp.Body)
  if err != nil {
    return nil, err
  }

  var articleSet pubmedArticleSet
  if err := xml.Unmarshal(fetchBody, &articleSet); err != nil {
    return nil, err
  }

  results := make([]articleInfo, 0, len(articleSet.Articles))
  for _, item := range articleSet.Articles {
    info := articleInfo{
      PMID:    strings.TrimSpace(item.MedlineCitation.PMID),
      Title:   strings.TrimSpace(item.MedlineCitation.Article.ArticleTitle),
      Journal: strings.TrimSpace(item.MedlineCitation.Article.Journal.Title),
    }
    info.Abstract = strings.TrimSpace(strings.Join(item.MedlineCitation.Article.Abstract.AbstractText, "\n"))
    info.Authors = formatAuthors(item.MedlineCitation.Article.AuthorList.Authors)
    info.PubDate = formatPubDate(
      item.MedlineCitation.Article.Journal.JournalIssue.PubDate.Year,
      item.MedlineCitation.Article.Journal.JournalIssue.PubDate.Month,
      item.MedlineCitation.Article.Journal.JournalIssue.PubDate.Day,
    )
    if info.Abstract == "" {
      continue
    }
    results = append(results, info)
  }

  if len(results) == 0 {
    return nil, errors.New("no usable abstracts found")
  }

  return results, nil
}

func formatAuthors(authors []pubmedAuthor) string {
  if len(authors) == 0 {
    return "N/A"
  }
  var names []string
  for _, author := range authors {
    name := strings.TrimSpace(strings.Join([]string{author.ForeName, author.LastName}, " "))
    if name != "" {
      names = append(names, name)
    }
  }
  if len(names) == 0 {
    return "N/A"
  }
  if len(names) > 3 {
    return strings.Join(names[:3], ", ") + ", et al."
  }
  return strings.Join(names, ", ")
}

func formatPubDate(year string, month string, day string) string {
  parts := []string{}
  if year != "" {
    parts = append(parts, year)
  }
  if month != "" {
    parts = append(parts, month)
  }
  if day != "" {
    parts = append(parts, day)
  }
  if len(parts) == 0 {
    return "N/A"
  }
  return strings.Join(parts, "-")
}

func (service *ReviewService) buildDocx(taskID string, payload WritingRequest, articles []articleInfo) (string, error) {
  if err := os.MkdirAll(service.storageDir, 0o755); err != nil {
    return "", err
  }

  filename := fmt.Sprintf("%s.docx", taskID)
  filePath := filepath.Join(service.storageDir, filename)

  content := buildDocxText(payload, articles)
  if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
    return "", err
  }

  return filePath, nil
}

func buildDocxText(payload WritingRequest, articles []articleInfo) string {
  builder := strings.Builder{}
  builder.WriteString(payload.Title)
  builder.WriteString("\n")
  builder.WriteString(fmt.Sprintf("主题：%s\n", payload.Topic))
  builder.WriteString(fmt.Sprintf("生成时间：%s\n\n", time.Now().Format("2006-01-02 15:04")))
  builder.WriteString("引言\n")
  builder.WriteString("本文基于 PubMed 检索结果整理相关文献摘要，快速生成综述草稿，供后续人工润色与完善。\n\n")
  builder.WriteString("文献摘要与要点\n")
  for index, article := range articles {
    builder.WriteString(fmt.Sprintf("文献 %d\n", index+1))
    builder.WriteString(fmt.Sprintf("标题：%s\n", article.Title))
    builder.WriteString(fmt.Sprintf("作者：%s\n", article.Authors))
    builder.WriteString(fmt.Sprintf("期刊：%s\n", article.Journal))
    builder.WriteString(fmt.Sprintf("发表时间：%s\n", article.PubDate))
    builder.WriteString(fmt.Sprintf("PMID：%s\n", article.PMID))
    builder.WriteString(fmt.Sprintf("摘要：%s\n\n", article.Abstract))
  }
  builder.WriteString("结论\n")
  builder.WriteString("本综述为自动生成的初稿，建议结合原始文献与研究目标进行进一步完善。\n")
  return builder.String()
}

 
