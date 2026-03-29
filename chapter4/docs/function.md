# 機能一覧

## 1. リポジトリ情報取得 (`repo get`)
指定したオーナー/リポジトリ名からリポジトリの詳細情報を取得する。

### コマンド
```
ghrepo repo get <owner>/<repo>
```

### 表示項目
- リポジトリ名
- 説明 (description)
- 言語 (language)
- スター数
- フォーク数
- 作成日
- 最終更新日
- リポジトリURL

## 2. ユーザーリポジトリ一覧取得 (`user repos`)
指定したユーザーのpublicリポジトリ一覧を取得する。

### コマンド
```
ghrepo user repos <username>
```

### 表示項目
- リポジトリ名
- 説明 (description)
- 言語 (language)
- スター数

### オプション
- `--sort`: ソート基準 (stars, updated, name) デフォルト: updated
- `--limit`: 表示件数 デフォルト: 10
