package gist

import (
	"github.com/MakeNowJust/heredoc"
	gistCloneCmd "github.com/cli/cli/v2/pkg/cmd/gist/clone"
	gistCreateCmd "github.com/cli/cli/v2/pkg/cmd/gist/create"
	gistDeleteCmd "github.com/cli/cli/v2/pkg/cmd/gist/delete"
	gistEditCmd "github.com/cli/cli/v2/pkg/cmd/gist/edit"
	gistListCmd "github.com/cli/cli/v2/pkg/cmd/gist/list"
	gistRenameCmd "github.com/cli/cli/v2/pkg/cmd/gist/rename"
	gistViewCmd "github.com/cli/cli/v2/pkg/cmd/gist/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdGist(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gist <command>",
		Short: "管理 Gist",
		Long:  `处理 GitHub Gist。`,
		Annotations: map[string]string{
			"help:arguments": heredoc.Doc(`
				Gist 可以通过以下任意格式作为参数提供：
				- 按 ID，例如 5b0e0062eb8e9654adad7bb1d81cc75f
				- 按 URL，例如 "https://gist.github.com/OWNER/5b0e0062eb8e9654adad7bb1d81cc75f"
			`),
		},
		GroupID: "core",
	}

	cmd.AddCommand(gistCloneCmd.NewCmdClone(f, nil))
	cmd.AddCommand(gistCreateCmd.NewCmdCreate(f, nil))
	cmd.AddCommand(gistListCmd.NewCmdList(f, nil))
	cmd.AddCommand(gistViewCmd.NewCmdView(f, nil))
	cmd.AddCommand(gistEditCmd.NewCmdEdit(f, nil))
	cmd.AddCommand(gistDeleteCmd.NewCmdDelete(f, nil))
	cmd.AddCommand(gistRenameCmd.NewCmdRename(f, nil))

	return cmd
}
