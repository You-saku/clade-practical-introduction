# 設計書

## 技術スタック
- 言語: Go
- CLIフレームワーク: cobra
- HTTPクライアント: net/http (標準ライブラリ)
- JSON処理: encoding/json (標準ライブラリ)
- 外部API: GitHub REST API v3 (認証なし)

## アプリケーション設計
- CLIアーキテクチャ: cobraによるサブコマンド構成
- レイヤー構成: cmd(CLI定義) → service(ビジネスロジック) → client(API通信)

## フォルダ構成
```
.
├── main.go
├── go.mod
├── go.sum
├── cmd/
│   ├── root.go        # ルートコマンド定義
│   ├── repo.go        # リポジトリ情報取得コマンド
│   └── user.go        # ユーザーリポジトリ一覧取得コマンド
├── internal/
│   ├── client/
│   │   └── github.go  # GitHub API クライアント
│   └── model/
│       └── repo.go    # リポジトリモデル定義
└── docs/
```

## 技術制約
- GitHub API のレートリミット: 認証なしで60リクエスト/時間
- publicリポジトリのみ対象

## パフォーマンス要件
- API レスポンスのタイムアウト: 10秒
