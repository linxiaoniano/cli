package preview

import (
	"github.com/MakeNowJust/heredoc"
	cmdPrompter "github.com/cli/cli/v2/pkg/cmd/preview/prompter"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdPreview(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preview <command>",
		Short: "执行 gh 功能的预览",
		Long: heredoc.Doc(`
			预览命令仅用于测试、演示和开发目的。
			它们应被视为不稳定的，随时可能更改。
		`),
	}

	cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(cmdPrompter.NewCmdPrompter(f, nil))

	return cmd
}
