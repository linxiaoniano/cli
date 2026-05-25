package secret

import (
	"github.com/MakeNowJust/heredoc"
	cmdDelete "github.com/cli/cli/v2/pkg/cmd/secret/delete"
	cmdList "github.com/cli/cli/v2/pkg/cmd/secret/list"
	cmdSet "github.com/cli/cli/v2/pkg/cmd/secret/set"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdSecret(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret <command>",
		Short: "管理 GitHub 密钥",
		Long: heredoc.Docf(`
			密钥可以在仓库或组织级别设置，用于 GitHub Actions、Agent 或 Dependabot。
			用户、组织和仓库密钥可以设置用于 GitHub Codespaces。环境密钥可以设置用于
			GitHub Actions。运行 %[1]sgh help secret set%[1]s 以了解如何开始。
		`, "`"),
	}

	cmdutil.EnableRepoOverride(cmd, f)

	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdSet.NewCmdSet(f, nil))
	cmd.AddCommand(cmdDelete.NewCmdDelete(f, nil))

	return cmd
}
