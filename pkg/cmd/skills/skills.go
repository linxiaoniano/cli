package skills

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/v2/internal/gh/ghtelemetry"
	"github.com/cli/cli/v2/pkg/cmd/skills/install"
	"github.com/cli/cli/v2/pkg/cmd/skills/preview"
	"github.com/cli/cli/v2/pkg/cmd/skills/publish"
	"github.com/cli/cli/v2/pkg/cmd/skills/search"
	"github.com/cli/cli/v2/pkg/cmd/skills/update"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

// NewCmdSkills returns the top-level "skill" command.
func NewCmdSkills(f *cmdutil.Factory, telemetry ghtelemetry.CommandRecorder) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "skill <command>",
		Short: "安装和管理 Agent 技能（预览）",
		Long: heredoc.Doc(`
			从 GitHub 仓库安装和管理 Agent 技能。

			在 GitHub CLI 中使用 Agent 技能处于预览阶段，
			可能随时更改，恕不另行通知。
		`),
		Aliases: []string{"skills"},
		GroupID: "core",
		Example: heredoc.Doc(`
			# Search for skills
			$ gh skill search terraform

			# Install a skill
			$ gh skill install github/awesome-copilot documentation-writer

			# Preview a skill before installing
			$ gh skill preview github/awesome-copilot documentation-writer

			# Update all installed skills
			$ gh skill update --all

			# Validate skills for publishing
			$ gh skill publish --dry-run
		`),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			telemetry.SetSampleRate(ghtelemetry.SAMPLE_ALL)
			return nil
		},
	}

	cmd.AddCommand(install.NewCmdInstall(f, telemetry, nil))
	cmd.AddCommand(preview.NewCmdPreview(f, telemetry, nil))
	cmd.AddCommand(publish.NewCmdPublish(f, nil))
	cmd.AddCommand(search.NewCmdSearch(f, telemetry, nil))
	cmd.AddCommand(update.NewCmdUpdate(f, nil))

	return cmd
}
