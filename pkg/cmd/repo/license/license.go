package license

import (
	cmdList "github.com/cli/cli/v2/pkg/cmd/repo/license/list"
	cmdView "github.com/cli/cli/v2/pkg/cmd/repo/license/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdLicense(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "license <command>",
		Short: "浏览仓库许可证",
	}

	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdView.NewCmdView(f, nil))

	return cmd
}
