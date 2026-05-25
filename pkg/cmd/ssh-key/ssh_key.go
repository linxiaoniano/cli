package key

import (
	cmdAdd "github.com/cli/cli/v2/pkg/cmd/ssh-key/add"
	cmdDelete "github.com/cli/cli/v2/pkg/cmd/ssh-key/delete"
	cmdList "github.com/cli/cli/v2/pkg/cmd/ssh-key/list"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdSSHKey(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssh-key <command>",
		Short: "管理 SSH 密钥",
		Long:  "管理已在 GitHub 账户中注册的 SSH 密钥。",
	}

	cmd.AddCommand(cmdAdd.NewCmdAdd(f, nil))
	cmd.AddCommand(cmdDelete.NewCmdDelete(f, nil))
	cmd.AddCommand(cmdList.NewCmdList(f, nil))

	return cmd
}
