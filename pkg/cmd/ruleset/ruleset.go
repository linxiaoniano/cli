package ruleset

import (
	"github.com/MakeNowJust/heredoc"
	cmdCheck "github.com/cli/cli/v2/pkg/cmd/ruleset/check"
	cmdList "github.com/cli/cli/v2/pkg/cmd/ruleset/list"
	cmdView "github.com/cli/cli/v2/pkg/cmd/ruleset/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRuleset(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ruleset <command>",
		Short: "查看仓库规则集信息",
		Long: heredoc.Doc(`
			仓库规则集是一种定义应用于仓库的一组规则的方式。
			这些命令允许你查看相关信息。
		`),
		Aliases: []string{"rs"},
		Example: heredoc.Doc(`
			$ gh ruleset list
			$ gh ruleset view --repo OWNER/REPO --web
			$ gh ruleset check branch-name
		`),
	}

	cmdutil.EnableRepoOverride(cmd, f)
	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdView.NewCmdView(f, nil))
	cmd.AddCommand(cmdCheck.NewCmdCheck(f, nil))

	return cmd
}
