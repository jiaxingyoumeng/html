# n8n 工作流 → 平台写作服务映射

本说明将现有 n8n 工作流拆解为平台内的写作服务步骤，便于后端与前端集成：输入标题后即可检索 PubMed、生成综述并输出 Word。

## 1. 请求入口（平台）

接口：`POST /api/writing`

```json
{
  "title": "综述标题",
  "topic": "检索主题",
  "locale": "zh-CN",
  "pubmedQuery": "",
  "includeFilters": false
}
```

前端在用户填写标题与主题后调用该接口，后端将请求转发至工作流执行器或队列。

## 2. 工作流步骤映射

| n8n 节点 | 平台服务步骤 | 输出 |
| --- | --- | --- |
| 计算当前日期 | 生成当前日期上下文 | `currentDateInfo` |
| 生成 PubMed 检索式 | LLM 生成检索式 | `pubmedQuery` |
| Esearch PMID / Efetch Abstract | PubMed 检索+拉取摘要 | XML / Medline |
| Parse / Split / Process Article Data | 解析摘要 + 作者 + 期刊 | 结构化摘要列表 |
| 整合摘要 | 生成合并摘要与检索要求 | `combined_summary` |
| 综述框架 agent | 输出综述大纲 | `outline` |
| 综述写作 agent | 生成 Markdown 正文 | `reviewMarkdown` |
| 生成 AMA 参考文献 | 参考文献列表 | `referencesSection` |
| 合并参考文献 | 最终 Markdown | `finalMarkdown` |
| MD 转 docx | Word 文件 | `review.docx` |

## 3. 关键业务规则

1. PubMed 检索式必须忽略 JCR/中科院分区/影响因子等指标，仅使用标准 PubMed 字段。
2. 当用户要求分区或影响因子筛选时，二次检索/过滤逻辑由独立脚本处理，不应拼接到 PubMed 检索式。
3. 输出文档需保留 Markdown → Word 转换流程，最终产生可下载的 docx。

## 4. 平台内服务接口建议

- `POST /api/writing`：提交写作任务，返回任务状态。
- `GET /api/writing/:id`：查询任务进度（排队/生成中/完成/失败）。
- `GET /api/writing/:id/download`：下载 Word 文件。

## 5. 前端对接建议

- 标题 + 主题输入后，调用写作接口。
- 完成后展示下载按钮，链接至 `/api/writing/:id/download`。

