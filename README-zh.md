# Member Service
[![auto test](https://github.com/iamsad5566/member_service_frame/actions/workflows/test.yml/badge.svg)](https://github.com/iamsad5566/member_service_frame/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamsad5566/member_service_frame)](https://goreportcard.com/report/github.com/iamsad5566/member_service_frame)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/iamsad5566/member_service_frame/graph/badge.svg?token=NTFKVYJH4K)](https://codecov.io/gh/iamsad5566/member_service_frame)


[中文版](/README-zh.md) , [English](/README.md)

會員服務是一個使用Go語言開發的開源微服務，處理使用者驗證、註冊和登入等功能,整合了各種與使用者管理相關的使用案例和情況。

## 功能
- 使用者註冊
- 使用者登入
- 使用者驗證
- 密碼更新
- OAuth2.0整合(Google)
- 基於gRPC的授權服務
- 自動建立資料庫
- 容器化部署

## 開始使用
### 先決條件
- Go (版本1.16或更新)
- PostgreSQL
- Redis

## 安裝
1. 複製 Repo：
```
https://github.com/iamsad5566/member_service_frame.git
```
2. 切換到專案目錄：
```
cd member_service_frame
```
3. 建置專案：
```
go build
```

## 設定
應用程式使用配置檔(`example_config.yml`)來存儲各種設定，以下是設定選項詳細說明：
- **valid_login:** 使用者登入有效天數 (預設：14)。
- **member_service:**
    - **host:** 會員服務主機(範例中已隱藏)。
    - **port:** 會員服務埠號 (預設：888)。
    - **RESTfulPort:**  RESTful API 埠號 (預設：8112)。
    - **gRPCPort:**  gRPC 服務埠號 (預設：8113)。

- **jwt:**
    - **secret_key:** JSON Web Token 密鑰 (範例中已隱藏)。
    - **token_expire:** Token 過期時間(秒) (預設：86400, 即24小時)。
- **db:**
    - **psql:**
        - **account:** PostgreSQL 資料庫帳號 (範例中已隱藏)。
        - **password:** PostgreSQL 資料庫密碼 (範例中已隱藏)。
        - **host:** PostgreSQL 資料庫主機 (範例中已隱藏)。
        - **port:** PostgreSQL 資料庫埠號 (預設：5433)。
        - **maxIdleConns:** PostgreSQL 連線池最大空閒連線數 (預設：20)。
        - **maxOpenConns:** PostgreSQL 連線池最大開啟連線數 (預設：20)。
        - **maxLifeMinute:** PostgreSQL 連線池連線最長存活時間(分鐘) (預設：10)。
    - **redis:**
        - **password:** Redis 密碼 (範例中已隱藏)。
        - **host:** Redis 主機 (範例中已隱藏)。
        - **port:** Redis 埠號 (預設：6379)。
- **logConfig:**
    - **level:** 日誌級別 (預設：info)。
    - **filename:** 日誌檔案名稱 (預設：logs/viper_zap_gin.log)。
    - **maxsize:**  日誌檔案大小上限(MB) (預設：1)。
    - **max_age:**  日誌檔案最長保存天數 (預設：30)。
    - **max_backups:** 最多保留的舊日誌檔案數量 (預設：5)。
- **oauth2**
    - **google:** 
        - **client_id:** Google OAuth2.0 客戶端ID (範例中已隱藏)。
        - **client_secret:** Google OAuth2.0 客戶端秘鑰 (範例中已隱藏)。

> [!NOTE]
> 請注意，出於安全考慮，範例中的敏感值(如密碼、秘鑰)已被隱藏。

#### 使用開源版設定實作
如果您偏好使用開源版本的設定實作，可以遵循以下步驟:
1. 將`example_config.yml`重新命名為`config.yml`。
2. 根據您的環境更新配置值。

### 使用私有設定實作(setconf)
或者，您可以使用私有的 `setconf` 存放庫（setconf 是一個 private repo，只有作者本人可以使用）來管理設定，在這種情況下，您需要用自己的實現替換`config.Setting`對象。

### 運行應用程式
配置好設定後，可以用以下指令運行應用程式:
```
./member_service_frame
```
這將啟動會員服務，包括RESTful API、gRPC服務，並自動創建資料庫(如果資料庫不存在)。

### API文件
會員服務提供Swagger UI的API文件，您可以在`http://localhost:8080/swagger/index.html`訪問(如果配置的埠號不是8080，請替換為對應值)。

### 部署
可以使用Docker容器部署會員服務，此專案提供了一個Dockerfile：
1. **Dockerfile:** 用於建置和運行應用程式。

### 建置Docker映像檔
若要建置Docker映像檔,請運行以下指令：
```
docker build --build-arg GITHUB_TOKEN=<your_github_token> --build-arg LATEST_SETCONF_VERSION=<setconf_version> -t member_service .
```

將<your_github_token>替換為您的GitHub個人訪問令牌,<setconf_version>替換為setconf存放庫的最新版本。順帶一提，若你決定使用開源版設定實作，請將 `Dockerfile`、`./github/workflows/ci-cd.yml`，和 `./github/workflows/test.yml` 中所有與 `setconf` 和 `GOPRIVATE` 有關的內容刪除。

### 運行Docker容器
建置好映像檔後,可以用以下指令運行容器:
```
docker run -d --name member_service -p 8080:8080 member_service
```
這將啟動會員服務容器並將主機的8080埠號對映到容器。


### CI/CD
會員服務包含用於持續整合(CI)和持續部署(CD)的GitHub Actions工作流程。

#### CI工作流程
CI工作流程(`auto test`)在每次推送到main分支時觸發,執行以下步驟:
1. 檢測原始碼
2. 設置`GOPRIVATE`環境變數
3. 登錄GitHub套件庫
4. 建置測試Docker映像檔(傳入GitHub Token和`setconf`版本作為建置參數)
5. 在Docker容器中運行測試

#### CD工作流程
CD工作流程(`CI/CD`)在發佈新版本時觸發,執行以下步驟:
1. 檢測原始碼
2. 設置`GOPRIVATE`環境變數
3. 登錄GitHub套件庫
4. 建置測試Docker映像檔(傳入GitHub Token和`setconf`版本作為建置參數)
5. 為映像檔添加標籤並推送到Docker Hub
6. 通過SSH將Docker映像檔部署到遠端伺服器

> [!NOTE]
> 請注意,您需要配置所需的GitHub Secrets(TOKEN、LATEST_SETCONF_VERSION、REMOTE_HOST、SSH_PRIVATE_KEY、DOCKER_USERNAME和DOCKER_PASSWORD),以使工作流程正常運行。

### 貢獻
歡迎為會員服務做出貢獻!如果您發現任何問題或想要添加新功能,請開啟issue或提交pull request。

### 授權
會員服務是根據MIT授權發佈的開源軟體。