# ghrepo

GitHub の public リポジトリ情報をターミナルから取得する CLI ツールです。

## インストール

```bash
go install github.com/yusakusekine/ghrepo@latest
```

またはソースからビルド:

```bash
git clone https://github.com/yusakusekine/ghrepo.git
cd ghrepo
go build -o ghrepo ./main.go
```

## 使い方

### リポジトリの詳細情報を取得

```bash
ghrepo repo get <owner>/<repo>
```

例:

```bash
$ ghrepo repo get golang/go
Name:        golang/go
Description: The Go programming language
Language:    Go
Stars:       132504
Forks:       18820
Created:     2014-08-19
Updated:     2026-02-16
URL:         https://github.com/golang/go
```

### ユーザーの public リポジトリ一覧を取得

```bash
ghrepo user repos <username> [flags]
```

| フラグ | 説明 | デフォルト |
|---|---|---|
| `--sort` | ソート基準 (stars, updated, name) | updated |
| `--limit` | 表示件数 | 10 |

例:

```bash
$ ghrepo user repos torvalds --limit 5 --sort stars
NAME              DESCRIPTION                               LANGUAGE  STARS
linux             Linux kernel source tree                  C         217594
AudioNoise        Random digital audio effects              C         4222
uemacs            Random version of microemacs with my ...  C         1878
GuitarPedal       Linus learns analog circuits              C         1750
HunspellColorize  Wrapper around 'less' to colorize spe...  C         302
```

## 注意事項

- 認証なしで GitHub REST API v3 を利用するため、レートリミットは 60 リクエスト/時間です
- public リポジトリのみ取得できます
