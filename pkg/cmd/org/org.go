package org

import (
	"github.com/MakeNowJust/heredoc"
	orgListCmd "github.com/cli/cli/v2/pkg/cmd/org/list"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdOrg(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "org <command>",
		Short: "管理组织",
		Long:  "处理 GitHub 组织。",
		Example: heredoc.Doc(`
			$ gh org list
		`),
		GroupID: "core",
	}

	cmdutil.AddGroup(cmd, "常用命令", orgListCmd.NewCmdList(f, nil))

	return cmd
}
