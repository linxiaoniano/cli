package workflow

import (
	cmdDisable "github.com/cli/cli/v2/pkg/cmd/workflow/disable"
	cmdEnable "github.com/cli/cli/v2/pkg/cmd/workflow/enable"
	cmdList "github.com/cli/cli/v2/pkg/cmd/workflow/list"
	cmdRun "github.com/cli/cli/v2/pkg/cmd/workflow/run"
	cmdView "github.com/cli/cli/v2/pkg/cmd/workflow/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdWorkflow(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workflow <command>",
		Short:   "查看 GitHub Actions 工作流详情",
		Long:    "列出、查看和运行 GitHub Actions 中的工作流。",
		GroupID: "actions",
	}
	cmdutil.EnableRepoOverride(cmd, f)

	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdEnable.NewCmdEnable(f, nil))
	cmd.AddCommand(cmdDisable.NewCmdDisable(f, nil))
	cmd.AddCommand(cmdView.NewCmdView(f, nil))
	cmd.AddCommand(cmdRun.NewCmdRun(f, nil))

	return cmd
}
