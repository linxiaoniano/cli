package pr

import (
	"github.com/MakeNowJust/heredoc"
	cmdLock "github.com/cli/cli/v2/pkg/cmd/issue/lock"
	cmdCheckout "github.com/cli/cli/v2/pkg/cmd/pr/checkout"
	cmdChecks "github.com/cli/cli/v2/pkg/cmd/pr/checks"
	cmdClose "github.com/cli/cli/v2/pkg/cmd/pr/close"
	cmdComment "github.com/cli/cli/v2/pkg/cmd/pr/comment"
	cmdCreate "github.com/cli/cli/v2/pkg/cmd/pr/create"
	cmdDiff "github.com/cli/cli/v2/pkg/cmd/pr/diff"
	cmdEdit "github.com/cli/cli/v2/pkg/cmd/pr/edit"
	cmdList "github.com/cli/cli/v2/pkg/cmd/pr/list"
	cmdMerge "github.com/cli/cli/v2/pkg/cmd/pr/merge"
	cmdReady "github.com/cli/cli/v2/pkg/cmd/pr/ready"
	cmdReopen "github.com/cli/cli/v2/pkg/cmd/pr/reopen"
	cmdRevert "github.com/cli/cli/v2/pkg/cmd/pr/revert"
	cmdReview "github.com/cli/cli/v2/pkg/cmd/pr/review"
	cmdStatus "github.com/cli/cli/v2/pkg/cmd/pr/status"
	cmdUpdateBranch "github.com/cli/cli/v2/pkg/cmd/pr/update-branch"
	cmdView "github.com/cli/cli/v2/pkg/cmd/pr/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdPR(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pr <command>",
		Short: "管理 Pull Request",
		Long:  "处理 GitHub Pull Request。",
		Example: heredoc.Doc(`
			$ gh pr checkout 353
			$ gh pr create --fill
			$ gh pr view --web
		`),
		Annotations: map[string]string{
			"help:arguments": heredoc.Doc(`
				Pull Request 可以通过以下任意格式作为参数提供：
				- 按编号，例如 "123"；
				- 按 URL，例如 "https://github.com/OWNER/REPO/pull/123"；或者
				- 按 head 分支名称，例如 "patch-1" 或 "OWNER:patch-1"。
			`),
		},
		GroupID: "core",
	}

	cmdutil.EnableRepoOverride(cmd, f)

	cmdutil.AddGroup(cmd, "常用命令",
		cmdList.NewCmdList(f, nil),
		cmdCreate.NewCmdCreate(f, nil),
		cmdStatus.NewCmdStatus(f, nil),
	)

	cmdutil.AddGroup(cmd, "目标命令",
		cmdView.NewCmdView(f, nil),
		cmdDiff.NewCmdDiff(f, nil),
		cmdCheckout.NewCmdCheckout(f, nil),
		cmdChecks.NewCmdChecks(f, nil),
		cmdReview.NewCmdReview(f, nil),
		cmdMerge.NewCmdMerge(f, nil),
		cmdUpdateBranch.NewCmdUpdateBranch(f, nil),
		cmdReady.NewCmdReady(f, nil),
		cmdComment.NewCmdComment(f, nil),
		cmdClose.NewCmdClose(f, nil),
		cmdReopen.NewCmdReopen(f, nil),
		cmdRevert.NewCmdRevert(f, nil),
		cmdEdit.NewCmdEdit(f, nil),
		cmdLock.NewCmdLock(f, cmd.Name(), nil),
		cmdLock.NewCmdUnlock(f, cmd.Name(), nil),
	)

	return cmd
}
