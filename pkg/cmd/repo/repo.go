package repo

import (
	"github.com/MakeNowJust/heredoc"
	repoArchiveCmd "github.com/cli/cli/v2/pkg/cmd/repo/archive"
	repoAutolinkCmd "github.com/cli/cli/v2/pkg/cmd/repo/autolink"
	repoCloneCmd "github.com/cli/cli/v2/pkg/cmd/repo/clone"
	repoCreateCmd "github.com/cli/cli/v2/pkg/cmd/repo/create"
	creditsCmd "github.com/cli/cli/v2/pkg/cmd/repo/credits"
	repoDeleteCmd "github.com/cli/cli/v2/pkg/cmd/repo/delete"
	deployKeyCmd "github.com/cli/cli/v2/pkg/cmd/repo/deploy-key"
	repoEditCmd "github.com/cli/cli/v2/pkg/cmd/repo/edit"
	repoForkCmd "github.com/cli/cli/v2/pkg/cmd/repo/fork"
	gardenCmd "github.com/cli/cli/v2/pkg/cmd/repo/garden"
	gitIgnoreCmd "github.com/cli/cli/v2/pkg/cmd/repo/gitignore"
	licenseCmd "github.com/cli/cli/v2/pkg/cmd/repo/license"
	repoListCmd "github.com/cli/cli/v2/pkg/cmd/repo/list"
	repoRenameCmd "github.com/cli/cli/v2/pkg/cmd/repo/rename"
	repoDefaultCmd "github.com/cli/cli/v2/pkg/cmd/repo/setdefault"
	repoSyncCmd "github.com/cli/cli/v2/pkg/cmd/repo/sync"
	repoUnarchiveCmd "github.com/cli/cli/v2/pkg/cmd/repo/unarchive"
	repoViewCmd "github.com/cli/cli/v2/pkg/cmd/repo/view"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRepo(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo <command>",
		Short: "管理仓库",
		Long:  `处理 GitHub 仓库。`,
		Example: heredoc.Doc(`
			$ gh repo create
			$ gh repo clone cli/cli
			$ gh repo view --web
		`),
		Annotations: map[string]string{
			"help:arguments": heredoc.Doc(`
				仓库可以通过以下任意格式作为参数提供：
				- "OWNER/REPO"
				- 按 URL，例如 "https://github.com/OWNER/REPO"
			`),
		},
		GroupID: "core",
	}

	cmdutil.AddGroup(cmd, "常用命令",
		repoListCmd.NewCmdList(f, nil),
		repoCreateCmd.NewCmdCreate(f, nil),
	)

	cmdutil.AddGroup(cmd, "目标命令",
		repoViewCmd.NewCmdView(f, nil),
		repoCloneCmd.NewCmdClone(f, nil),
		repoForkCmd.NewCmdFork(f, nil),
		repoDefaultCmd.NewCmdSetDefault(f, nil),
		repoSyncCmd.NewCmdSync(f, nil),
		repoEditCmd.NewCmdEdit(f, nil),
		deployKeyCmd.NewCmdDeployKey(f),
		licenseCmd.NewCmdLicense(f),
		gitIgnoreCmd.NewCmdGitIgnore(f),
		repoRenameCmd.NewCmdRename(f, nil),
		repoArchiveCmd.NewCmdArchive(f, nil),
		repoUnarchiveCmd.NewCmdUnarchive(f, nil),
		repoDeleteCmd.NewCmdDelete(f, nil),
		creditsCmd.NewCmdRepoCredits(f, nil),
		gardenCmd.NewCmdGarden(f, nil),
		repoAutolinkCmd.NewCmdAutolink(f),
	)

	return cmd
}
