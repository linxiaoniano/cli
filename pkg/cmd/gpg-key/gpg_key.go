package key

import (
	cmdAdd "github.com/cli/cli/v2/pkg/cmd/gpg-key/add"
	cmdDelete "github.com/cli/cli/v2/pkg/cmd/gpg-key/delete"
	cmdList "github.com/cli/cli/v2/pkg/cmd/gpg-key/list"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdGPGKey(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gpg-key <command>",
		Short: "管理 GPG 密钥",
		Long:  "管理已在 GitHub 账户中注册的 GPG 密钥。",
	}

	cmd.AddCommand(cmdAdd.NewCmdAdd(f, nil))
	cmd.AddCommand(cmdDelete.NewCmdDelete(f, nil))
	cmd.AddCommand(cmdList.NewCmdList(f, nil))

	return cmd
}
