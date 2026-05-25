package variable

import (
	"github.com/MakeNowJust/heredoc"
	cmdDelete "github.com/cli/cli/v2/pkg/cmd/variable/delete"
	cmdGet "github.com/cli/cli/v2/pkg/cmd/variable/get"
	cmdList "github.com/cli/cli/v2/pkg/cmd/variable/list"
	cmdSet "github.com/cli/cli/v2/pkg/cmd/variable/set"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdVariable(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "variable <command>",
		Short: "管理 GitHub Actions 变量",
		Long: heredoc.Docf(`
			变量可以在仓库、环境或组织级别设置，用于 GitHub Actions 或 Dependabot。
			运行 %[1]sgh help variable set%[1]s 以了解如何开始。
		`, "`"),
	}

	cmdutil.EnableRepoOverride(cmd, f)

	cmd.AddCommand(cmdGet.NewCmdGet(f, nil))
	cmd.AddCommand(cmdSet.NewCmdSet(f, nil))
	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdDelete.NewCmdDelete(f, nil))

	return cmd
}
