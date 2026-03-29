package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yusakusekine/ghrepo/internal/model"
)

const defaultBaseURL = "https://api.github.com"

// Client はGitHub APIと通信するクライアント。
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient は新しいClientを生成する。
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    defaultBaseURL,
	}
}

// NewClientWithBaseURL はベースURLを指定してClientを生成する(テスト用)。
func NewClientWithBaseURL(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    baseURL,
	}
}

// GetRepository は指定されたリポジトリの詳細情報を取得する。
func (c *Client) GetRepository(owner, repo string) (*model.Repository, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.baseURL, owner, repo)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("APIリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("リポジトリ %s/%s が見つかりません", owner, repo)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("APIエラー: ステータスコード %d", resp.StatusCode)
	}

	var repository model.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repository); err != nil {
		return nil, fmt.Errorf("レスポンスの解析に失敗しました: %w", err)
	}

	return &repository, nil
}

// ListUserRepos は指定ユーザーのpublicリポジトリ一覧を取得する。
func (c *Client) ListUserRepos(username, sort string, limit int) ([]model.Repository, error) {
	apiSort := sort
	direction := "desc"
	if sort == "stars" {
		apiSort = "updated" // GitHub APIはstarsソートを直接サポートしないため後でソート
	}
	if sort == "name" {
		apiSort = "full_name"
		direction = "asc"
	}

	url := fmt.Sprintf("%s/users/%s/repos?sort=%s&direction=%s&per_page=%d",
		c.baseURL, username, apiSort, direction, limit)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("APIリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("ユーザー %s が見つかりません", username)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("APIエラー: ステータスコード %d", resp.StatusCode)
	}

	var repos []model.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, fmt.Errorf("レスポンスの解析に失敗しました: %w", err)
	}

	if sort == "stars" {
		sortByStars(repos)
	}

	return repos, nil
}

func sortByStars(repos []model.Repository) {
	for i := 0; i < len(repos)-1; i++ {
		for j := i + 1; j < len(repos); j++ {
			if repos[j].Stars > repos[i].Stars {
				repos[i], repos[j] = repos[j], repos[i]
			}
		}
	}
}
