# 設計: 初期実装

## アーキテクチャ

```
main.go
  └─ cmd/ (CLI定義層: cobra)
       ├─ root.go
       ├─ repo.go
       └─ user.go
  └─ internal/
       ├─ client/
       │   └─ github.go (API通信層)
       └─ model/
           └─ repo.go (データモデル)
```

## 各レイヤーの責務

### cmd層
- cobraによるコマンド定義・引数/フラグのパース
- clientを呼び出し、結果を整形して標準出力に表示

### internal/client層
- GitHub REST API v3 へのHTTPリクエスト送信
- JSONレスポンスのパースとモデルへの変換
- エラーハンドリング (タイムアウト、404等)

### internal/model層
- APIレスポンスに対応する構造体定義

## API設計

### GitHub APIクライアント
```go
type Client struct {
    httpClient *http.Client
    baseURL    string
}

func NewClient() *Client
func (c *Client) GetRepository(owner, repo string) (*model.Repository, error)
func (c *Client) ListUserRepos(username, sort string, limit int) ([]model.Repository, error)
```

### データモデル
```go
type Repository struct {
    Name        string
    Description string
    Language    string
    Stars       int
    Forks       int
    HTMLURL     string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## 出力フォーマット

### repo get
```
Name:        owner/repo
Description: A sample repository
Language:    Go
Stars:       150
Forks:       30
Created:     2024-01-15
Updated:     2025-12-01
URL:         https://github.com/owner/repo
```

### user repos
```
NAME                DESCRIPTION          LANGUAGE   STARS
repo-1              A sample project     Go         150
repo-2              Another project      Python      42
```

## 依存パッケージ
- `github.com/spf13/cobra` — CLIフレームワーク
