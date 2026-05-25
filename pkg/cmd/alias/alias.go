package alias

import (
	"github.com/MakeNowJust/heredoc"
	deleteCmd "github.com/cli/cli/v2/pkg/cmd/alias/delete"
	importCmd "github.com/cli/cli/v2/pkg/cmd/alias/imports"
	listCmd "github.com/cli/cli/v2/pkg/cmd/alias/list"
	setCmd "github.com/cli/cli/v2/pkg/cmd/alias/set"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAlias(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias <command>",
		Short: "创建命令快捷方式",
		Long: heredoc.Docf(`
				别名可用于为 gh 命令创建快捷方式，或组合多个命令。

				运行 %[1]sgh help alias set%[1]s 以了解更多信息。
			`, "`"),
	}

	cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(deleteCmd.NewCmdDelete(f, nil))
	cmd.AddCommand(importCmd.NewCmdImport(f, nil))
	cmd.AddCommand(listCmd.NewCmdList(f, nil))
	cmd.AddCommand(setCmd.NewCmdSet(f, nil))

	return cmd
}
