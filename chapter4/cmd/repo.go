package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yusakusekine/ghrepo/internal/client"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "リポジトリ操作",
}

var repoGetCmd = &cobra.Command{
	Use:   "get <owner/repo>",
	Short: "リポジトリの詳細情報を取得する",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.SplitN(args[0], "/", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			fmt.Fprintln(os.Stderr, "エラー: <owner/repo> の形式で指定してください")
			os.Exit(1)
		}

		c := client.NewClient()
		repo, err := c.GetRepository(parts[0], parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Name:        %s\n", repo.FullName)
		fmt.Printf("Description: %s\n", repo.Description)
		fmt.Printf("Language:    %s\n", repo.Language)
		fmt.Printf("Stars:       %d\n", repo.Stars)
		fmt.Printf("Forks:       %d\n", repo.Forks)
		fmt.Printf("Created:     %s\n", repo.CreatedAt.Format("2006-01-02"))
		fmt.Printf("Updated:     %s\n", repo.UpdatedAt.Format("2006-01-02"))
		fmt.Printf("URL:         %s\n", repo.HTMLURL)
	},
}

func init() {
	repoCmd.AddCommand(repoGetCmd)
	rootCmd.AddCommand(repoCmd)
}
