package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghrepo",
	Short: "GitHub publicリポジトリ情報取得CLIツール",
}

// Execute はルートコマンドを実行する。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
