package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/yusakusekine/ghrepo/internal/client"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "ユーザー操作",
}

var userReposCmd = &cobra.Command{
	Use:   "repos <username>",
	Short: "ユーザーのpublicリポジトリ一覧を取得する",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sort, _ := cmd.Flags().GetString("sort")
		limit, _ := cmd.Flags().GetInt("limit")

		c := client.NewClient()
		repos, err := c.ListUserRepos(args[0], sort, limit)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %s\n", err)
			os.Exit(1)
		}

		if len(repos) == 0 {
			fmt.Println("リポジトリが見つかりません")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tDESCRIPTION\tLANGUAGE\tSTARS")
		for _, r := range repos {
			desc := r.Description
			if len(desc) > 40 {
				desc = desc[:37] + "..."
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%d\n", r.Name, desc, r.Language, r.Stars)
		}
		w.Flush()
	},
}

func init() {
	userReposCmd.Flags().String("sort", "updated", "ソート基準 (stars, updated, name)")
	userReposCmd.Flags().Int("limit", 10, "表示件数")
	userCmd.AddCommand(userReposCmd)
	rootCmd.AddCommand(userCmd)
}
