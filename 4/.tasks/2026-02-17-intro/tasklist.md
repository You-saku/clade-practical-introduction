# タスクリスト: 初期実装

## 1. プロジェクト初期化
- [x] `go mod init` でモジュール初期化
- [x] cobra 依存パッケージの追加
- [x] フォルダ構成の作成 (`cmd/`, `internal/client/`, `internal/model/`)

## 2. データモデル定義
- [x] `internal/model/repo.go` に Repository 構造体を定義

## 3. GitHub APIクライアント実装
- [x] `internal/client/github.go` に Client 構造体と NewClient を実装
- [x] `GetRepository` メソッドを実装
- [x] `ListUserRepos` メソッドを実装
- [x] エラーハンドリング (404、タイムアウト等)

## 4. CLIコマンド実装
- [x] `cmd/root.go` にルートコマンドを定義
- [x] `cmd/repo.go` に `repo get` サブコマンドを実装
- [x] `cmd/user.go` に `user repos` サブコマンドを実装 (--sort, --limit フラグ)
- [x] `main.go` にエントリーポイントを作成

## 5. テスト実装
- [x] `internal/client/github_test.go` — APIクライアントのテスト (httptest利用)
- [x] `internal/model/repo_test.go` — モデルのテスト (不要と判断しスキップ)

## 6. 動作確認
- [x] `go build` でビルド確認
- [x] `ghrepo repo get` の動作確認
- [x] `ghrepo user repos` の動作確認
- [x] エラーケースの動作確認 (テストコードでカバー)
