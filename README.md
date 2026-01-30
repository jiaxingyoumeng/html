# 小智科研系统（Vue 3 + Go）

本仓库已调整为 **Vue 3 + TypeScript 前端** 与 **Go 后端** 的统一架构示例，用于支持科学文献综述助手的真实开发与部署。

## 项目结构

```
frontend/
├── public/
├── src/
│   ├── assets/
│   ├── components/
│   ├── layouts/
│   ├── pages/
│   ├── router/
│   ├── store/
│   ├── services/
│   ├── utils/
│   ├── App.vue
│   └── main.ts
├── index.html
├── package.json
├── tsconfig.json
└── vite.config.ts

backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   ├── config/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   └── service/
├── migrations/
├── pkg/
├── go.mod
└── go.sum
```

## 前端说明

- 使用 Vue 3 + Vite + Pinia + Vue Router。
- `services/apiService.ts` 集中管理 API 请求与拦截器。
- `store/auth.ts` 负责鉴权状态与本地持久化。
- `/writing` 写作服务直接调用 PubMed 接口生成综述草稿与 Word 文档。

## 后端说明

- 基于 Gin + MongoDB，预置 JWT 认证链路与文献综述相关路由。
- `internal/service` 中预留业务实现入口，可逐步补充实际逻辑。
- `docs/writing-service.md` 记录写作服务的代码实现与接口说明。

### PubMed 配置

- `PUBMED_API_KEY`：可选，提升 PubMed 限流额度。
- `PUBMED_BASE_URL`：可选，默认 `https://eutils.ncbi.nlm.nih.gov/entrez/eutils`。

## 本地开发

```bash
# 前端
cd frontend
npm install
npm run dev

# 后端
cd ../backend
go run ./cmd/server
```

## 快速启动（静态前端 + Go 后端）

无需安装前端依赖时，可使用脚本一键启动静态页面与后端服务：

```bash
./scripts/start-dev.sh
```

脚本默认会读取 `PUBMED_API_KEY`，未设置时会使用内置的示例 Key（可在启动前通过环境变量覆盖）。

停止服务：

```bash
./scripts/stop-dev.sh
```

### Windows 用户（PowerShell / CMD）

Windows 不支持直接运行 `./scripts/*.sh`，请使用对应的批处理脚本：

```bat
scripts\start-dev.bat
```

停止服务：

```bat
scripts\stop-dev.bat
```

如果你已安装 Git Bash，也可以在 Git Bash 里运行 `./scripts/start-dev.sh`。

## Docker 运行

```bash
docker-compose up --build
```
