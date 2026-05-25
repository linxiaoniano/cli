package label

import (
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdLabel(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "label <command>",
		Short: "管理标签",
		Long:  `处理 GitHub 标签。`,
	}
	cmdutil.EnableRepoOverride(cmd, f)

	cmd.AddCommand(newCmdList(f, nil))
	cmd.AddCommand(newCmdCreate(f, nil))
	cmd.AddCommand(newCmdClone(f, nil))
	cmd.AddCommand(newCmdEdit(f, nil))
	cmd.AddCommand(newCmdDelete(f, nil))

	return cmd
}
