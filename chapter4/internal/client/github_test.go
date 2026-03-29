package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yusakusekine/ghrepo/internal/model"
)

func setupTestServer(handler http.HandlerFunc) (*httptest.Server, *Client) {
	server := httptest.NewServer(handler)
	client := NewClientWithBaseURL(server.URL)
	return server, client
}

func TestGetRepository_Success(t *testing.T) {
	repo := model.Repository{
		Name:        "test-repo",
		FullName:    "testuser/test-repo",
		Description: "A test repository",
		Language:    "Go",
		Stars:       100,
		Forks:       10,
		HTMLURL:     "https://github.com/testuser/test-repo",
	}

	server, client := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/testuser/test-repo" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repo)
	})
	defer server.Close()

	result, err := client.GetRepository("testuser", "test-repo")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.FullName != "testuser/test-repo" {
		t.Errorf("expected FullName 'testuser/test-repo', got '%s'", result.FullName)
	}
	if result.Stars != 100 {
		t.Errorf("expected Stars 100, got %d", result.Stars)
	}
}

func TestGetRepository_NotFound(t *testing.T) {
	server, client := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	defer server.Close()

	_, err := client.GetRepository("testuser", "nonexistent")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestListUserRepos_Success(t *testing.T) {
	repos := []model.Repository{
		{Name: "repo-a", Stars: 50, Language: "Go"},
		{Name: "repo-b", Stars: 200, Language: "Python"},
		{Name: "repo-c", Stars: 10, Language: "Rust"},
	}

	server, client := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/testuser/repos" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repos)
	})
	defer server.Close()

	result, err := client.ListUserRepos("testuser", "updated", 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 3 {
		t.Fatalf("expected 3 repos, got %d", len(result))
	}
}

func TestListUserRepos_SortByStars(t *testing.T) {
	repos := []model.Repository{
		{Name: "repo-a", Stars: 50},
		{Name: "repo-b", Stars: 200},
		{Name: "repo-c", Stars: 10},
	}

	server, client := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repos)
	})
	defer server.Close()

	result, err := client.ListUserRepos("testuser", "stars", 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result[0].Name != "repo-b" {
		t.Errorf("expected first repo 'repo-b', got '%s'", result[0].Name)
	}
	if result[2].Name != "repo-c" {
		t.Errorf("expected last repo 'repo-c', got '%s'", result[2].Name)
	}
}

func TestListUserRepos_NotFound(t *testing.T) {
	server, client := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	defer server.Close()

	_, err := client.ListUserRepos("nonexistent", "updated", 10)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
