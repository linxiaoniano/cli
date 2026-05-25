package gitignore

import (
	cmdList "github.com/cli/cli/v2/pkg/cmd/repo/gitignore/list"
	cmdView "github.com/cli/cli/v2/pkg/cmd/repo/gitignore/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdGitIgnore(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gitignore <command>",
		Short: "列出和查看可用的仓库 gitignore 模板",
	}

	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdView.NewCmdView(f, nil))

	return cmd
}
