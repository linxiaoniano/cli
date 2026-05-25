package licenses

import (
	"fmt"

	"github.com/cli/cli/v2/internal/licenses"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdLicenses(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "licenses",
		Short: "查看第三方许可证信息",
		Long:  "查看 GitHub CLI 此版本使用的第三方库的许可证信息。",
		RunE: func(cmd *cobra.Command, args []string) error {
			io := f.IOStreams
			if err := io.StartPager(); err == nil {
				defer io.StopPager()
			}
			_, err := fmt.Fprint(io.Out, licenses.Content())
			return err
		},
	}

	cmdutil.DisableAuthCheck(cmd)

	return cmd
}
