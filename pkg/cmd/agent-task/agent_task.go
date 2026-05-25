package agent

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	cmdCreate "github.com/cli/cli/v2/pkg/cmd/agent-task/create"
	cmdList "github.com/cli/cli/v2/pkg/cmd/agent-task/list"
	cmdView "github.com/cli/cli/v2/pkg/cmd/agent-task/view"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/spf13/cobra"
)

// NewCmdAgentTask creates the base `agent-task` command.
func NewCmdAgentTask(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "agent-task <command>",
		Aliases: []string{"agent-tasks", "agent", "agents"},
		Short:   "处理 Agent 任务（预览）",
		Long: heredoc.Doc(`
			在 GitHub CLI 中使用 Agent 任务处于预览阶段，
			可能随时更改，恕不另行通知。
		`),
		Annotations: map[string]string{
			"help:arguments": heredoc.Doc(`
				任务可以通过以下任意格式作为参数指定：
				- 按 Pull Request 编号，例如 "123"；或者
				- 按会话 ID，例如 "12345abc-12345-12345-12345-12345abc"；或者
				- 按 URL，例如 "https://github.com/OWNER/REPO/pull/123/agent-sessions/12345abc-12345-12345-12345-12345abc"；

				不建议在非交互式用例中通过 Pull Request 来标识任务，
				因为给定的 Pull Request 可能有多个任务需要消歧。
			`),
		},
		Example: heredoc.Doc(`
			# List your most recent agent tasks
			$ gh agent-task list
			
			# Create a new agent task on the current repository
			$ gh agent-task create "Improve the performance of the data processing pipeline"
			
			# View details about agent tasks associated with a pull request
			$ gh agent-task view 123

			# View details about a specific agent task
			$ gh agent-task view 12345abc-12345-12345-12345-12345abc
		`),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return requireOAuthToken(f)
		},
		// This is required to run this root command. We want to
		// run it to test PersistentPreRunE behavior.
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// register subcommands
	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdCreate.NewCmdCreate(f, nil))
	cmd.AddCommand(cmdView.NewCmdView(f, nil))

	return cmd
}

// requireOAuthToken ensures an OAuth (device flow) token is present and valid.
// agent-task subcommands inherit this check via PersistentPreRunE.
func requireOAuthToken(f *cmdutil.Factory) error {
	cfg, err := f.Config()
	if err != nil {
		return err
	}

	authCfg := cfg.Authentication()
	host, _ := authCfg.DefaultHost()
	if host == "" {
		return errors.New("no default host configured; run 'gh auth login'")
	}

	if auth.IsEnterprise(host) {
		return errors.New("agent tasks are not supported on this host")
	}

	token, source := authCfg.ActiveToken(host)

	// Tokens from sources "oauth_token" and "keyring" are likely
	// minted through our device flow.
	tokenSourceIsDeviceFlow := source == "oauth_token" || source == "keyring"
	// Tokens with "gho_" prefix are OAuth tokens.
	tokenIsOAuth := strings.HasPrefix(token, "gho_")

	// Reject if the token is not from a device flow source or is not an OAuth token
	if !tokenSourceIsDeviceFlow || !tokenIsOAuth {
		return fmt.Errorf("this command requires an OAuth token. Re-authenticate with: gh auth login")
	}
	return nil
}
