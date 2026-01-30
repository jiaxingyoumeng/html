# 写作服务（代码实现）

该服务在后端直接调用 PubMed 接口完成检索、摘要聚合与 Word 文档生成，不依赖 n8n 工作流。

## 接口

- `POST /api/writing`：提交写作任务，返回任务 ID 与下载地址（完成后可直接下载）。
- `GET /api/writing/:id`：查询任务状态。
- `GET /api/writing/:id/download`：下载生成的 Word 文档。

请求示例：

```json
{
  "title": "综述标题",
  "topic": "检索主题",
  "locale": "zh-CN"
}
```

## 处理流程

1. 根据 `topic` 生成 PubMed 查询（可直接传入 `pubmedQuery` 覆盖默认查询）。
2. 调用 PubMed ESearch/EFetch 获取摘要、作者、期刊与发表日期。
3. 生成综述草稿并写入 Word（docx）。
4. 保存到 `backend/storage/` 并提供下载链接。

## 环境变量

- `PUBMED_API_KEY`：可选，提升 PubMed 限流额度。
- `PUBMED_BASE_URL`：可选，默认 `https://eutils.ncbi.nlm.nih.gov/entrez/eutils`。
