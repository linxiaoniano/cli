package root

import (
	"fmt"
	"io"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/v2/internal/text"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/spf13/cobra"
)

type helpTopic struct {
	name    string
	short   string
	long    string
	example string
}

var HelpTopics = []helpTopic{
	{
		name:  "mintty",
		short: "关于在 MinTTY 中使用 gh 的信息",
		long: heredoc.Docf(`
			MinTTY 是 Git for Windows 默认附带的终端模拟器。它存在已知问题，
			会影响 gh 提示用户输入的能力。

			以下是使 gh 能在 MinTTY 中工作的一些解决方法：

			- 重新安装 Git for Windows，并勾选"启用伪控制台的实验性支持"。

			- 使用其他终端模拟器，如 Windows Terminal。
			  你可以从任意终端模拟器中运行 %[1]sC:\Program Files\Git\bin\bash.exe%[1]s，
			  以继续使用 Git for Windows 中的所有工具，而无需 MinTTY。

			- 在 gh 命令前加上 %[1]swinpty%[1]s 前缀，例如：%[1]swinpty gh auth login%[1]s。
			  注意：这可能会导致一些 UI 问题。
		`, "`"),
	},
	{
		name:  "environment",
		short: "可与 gh 一起使用的环境变量",
		long: heredoc.Docf(`
			%[1]sGH_TOKEN%[1]s、%[1]sGITHUB_TOKEN%[1]s（按优先级顺序）：当命令的目标为 %[1]sgithub.com%[1]s
			或其子域 %[1]sghe.com%[1]s 时使用的身份验证令牌。设置此项可避免被提示进行身份验证，
			并优先于之前存储的凭据。

			%[1]sGH_ENTERPRISE_TOKEN%[1]s、%[1]sGITHUB_ENTERPRISE_TOKEN%[1]s（按优先级顺序）：当命令的目标是
			GitHub Enterprise Server 主机时使用的身份验证令牌。

			%[1]sGH_HOST%[1]s：为未提供主机名或无法从本地 Git 仓库上下文中推断主机名的命令指定
			GitHub 主机名。如果此主机之前已通过身份验证，则将使用存储的凭据。
			否则，根据目标主机需要设置 %[1]sGH_TOKEN%[1]s 或 %[1]sGH_ENTERPRISE_TOKEN%[1]s。

			%[1]sGH_REPO%[1]s：以 %[1]s[HOST/]OWNER/REPO%[1]s 格式为操作本地仓库的命令指定 GitHub 仓库。

			%[1]sGH_EDITOR%[1]s、%[1]sGIT_EDITOR%[1]s、%[1]sVISUAL%[1]s、%[1]sEDITOR%[1]s（按优先级顺序）：用于编辑文本的编辑器工具。

			%[1]sGH_BROWSER%[1]s、%[1]sBROWSER%[1]s（按优先级顺序）：用于打开链接的网页浏览器。

			%[1]sGH_DEBUG%[1]s：设置为真值以在标准错误上输出详细信息。设置为 %[1]sapi%[1]s
			可额外记录 HTTP 流量的详细信息。

			%[1]sDEBUG%[1]s（已弃用）：设置为 %[1]s1%[1]s、%[1]strue%[1]s 或 %[1]syes%[1]s 以在标准错误上启用详细输出。

			%[1]sGH_PAGER%[1]s、%[1]sPAGER%[1]s（按优先级顺序）：用于分页标准输出的终端分页程序，
			例如 %[1]sless%[1]s。

			%[1]sGLAMOUR_STYLE%[1]s：用于渲染 Markdown 的样式。请参见
			<https://github.com/charmbracelet/glamour#styles>

			%[1]sNO_COLOR%[1]s：设置为任意值以避免输出 ANSI 转义序列的颜色。

			%[1]sCLICOLOR%[1]s：设置为 %[1]s0%[1]s 以禁用输出中的 ANSI 颜色。

			%[1]sCLICOLOR_FORCE%[1]s：设置为非 %[1]s0%[1]s 的值以在输出被重定向时仍保留 ANSI 颜色。

			%[1]sGH_COLOR_LABELS%[1]s：设置为任意值以在支持真彩色的终端中
			使用标签的 RGB 十六进制颜色代码显示。

			%[1]sGH_ACCESSIBLE_COLORS%[1]s（预览）：设置为真值以使用可自定义的 4 位可访问颜色。

			%[1]sGH_FORCE_TTY%[1]s：设置为任意值以在输出被重定向时强制使用终端样式输出。
			如果值为数字，则解释为视口中可用的列数。如果值为百分比，
			则将其应用于当前视口中可用的列数。

			%[1]sGH_NO_UPDATE_NOTIFIER%[1]s：设置为任意值以禁用 GitHub CLI 更新通知。
			当执行任何命令时，gh 每 24 小时检查一次新版本。
			如果发现新版本，将在标准错误上显示升级通知。

			%[1]sGH_NO_EXTENSION_UPDATE_NOTIFIER%[1]s：设置为任意值以禁用 GitHub CLI 扩展更新通知。
			当执行扩展时，gh 每 24 小时检查一次所执行扩展的新版本。
			如果发现新版本，将在标准错误上显示升级通知。

			%[1]sGH_CONFIG_DIR%[1]s：gh 存储配置文件的目录。如果未指定，
			默认值将是以下路径之一（按优先级顺序）：
			  - %[1]s$XDG_CONFIG_HOME/gh%[1]s（如果设置了 %[1]s$XDG_CONFIG_HOME%[1]s）
			  - %[1]s$AppData/GitHub CLI%[1]s（在 Windows 上如果设置了 %[1]s$AppData%[1]s）
			  - %[1]s$HOME/.config/gh%[1]s

			%[1]sGH_PROMPT_DISABLED%[1]s：设置为任意值以禁用终端中的交互式提示。

			%[1]sGH_PATH%[1]s：设置 gh 可执行文件的路径，当 gh 无法正确确定自身路径时使用，
			例如在 cygwin 终端中。

			%[1]sGH_MDWIDTH%[1]s：markdown 渲染换行的默认最大宽度。终端上换行的最大宽度
			取终端宽度、此值和 120（如果未指定）中的较小值。
			此值例如与 %[1]spr view%[1]s 子命令一起使用。

			%[1]sGH_ACCESSIBLE_PROMPTER%[1]s（预览）：设置为真值以启用与语音合成和
			盲文屏幕阅读器更兼容的提示。

			%[1]sGH_TELEMETRY%[1]s：设置为 %[1]slog%[1]s 以将遥测数据打印到标准错误而不是发送。
			设置为 %[1]sfalse%[1]s 或 %[1]s0%[1]s 以禁用遥测。优先于 %[1]sDO_NOT_TRACK%[1]s。

			%[1]sDO_NOT_TRACK%[1]s：设置为 %[1]strue%[1]s 或 %[1]s1%[1]s 以禁用遥测。当设置了
			%[1]sGH_TELEMETRY%[1]s 时会被忽略。

			%[1]sGH_SPINNER_DISABLED%[1]s：设置为真值以用文本进度指示器替换旋转动画。
		`, "`"),
	},
	{
		name:  "telemetry",
		short: "关于 gh 中遥测的信息",
		long: heredoc.Doc(`
			gh 收集遥测数据以帮助我们了解 CLI 的使用情况并进行改进。

			要了解收集了哪些数据、如何使用以及如何选择退出，请访问：
			<https://cli.github.com/telemetry>
		`),
	},
	{
		name:  "reference",
		short: "所有 gh 命令的完整参考",
	},
	{
		name:  "formatting",
		short: "从 gh 导出的 JSON 数据的格式化选项",
		long: heredoc.Docf(`
			默认情况下，%[1]sgh%[1]s 命令的结果以基于行的纯文本格式输出。
			某些命令支持传递 %[1]s--json%[1]s 标志，将输出转换为 JSON 格式。
			转换为 JSON 后，可以通过添加 %[1]s--jq%[1]s 或 %[1]s--template%[1]s 标志，
			根据所需的格式化字符串进一步格式化输出。这对于选择数据子集、
			创建新数据结构、以不同格式显示数据或作为另一个命令行脚本的输入非常有用。

			%[1]s--json%[1]s 标志需要逗号分隔的字段列表。要查看命令可能的 JSON 字段名称，
			可以在运行命令时省略 %[1]s--json%[1]s 标志的字符串参数。
			请注意，要使用 %[1]s--jq%[1]s 或 %[1]s--template%[1]s 标志，
			必须先传递 %[1]s--json%[1]s 标志和字段名称。

			%[1]s--jq%[1]s 标志需要一个 jq 查询语法的字符串参数，并只打印与查询匹配的 JSON 值。
			jq 查询可用于从数组中选择元素、从对象中选择字段、创建新数组等。
			使用此格式化指令时，系统中无需安装 %[1]sjq%[1]s 工具。
			当连接到终端时，输出会自动进行美化打印。要了解 jq 查询语法，请访问：
			<https://jqlang.github.io/jq/manual/>

			%[1]s--template%[1]s 标志需要一个 Go 模板语法的字符串参数，并只打印与查询匹配的 JSON 值。

			除了标准库中的 Go 模板函数外，还可以使用以下函数与格式化指令配合使用：
			- %[1]sautocolor%[1]s：类似 %[1]scolor%[1]s，但仅在终端中输出颜色
			- %[1]scolor <样式> <输入>%[1]s：使用 <https://github.com/mgutz/ansi> 为输入着色
			- %[1]sjoin <分隔符> <列表>%[1]s：使用分隔符连接列表中的值
			- %[1]spluck <字段> <列表>%[1]s：从输入的所有项目中收集指定字段的值
			- %[1]stablerow <字段>...%[1]s：将字段在输出中垂直对齐为表格
			- %[1]stablerender%[1]s：在适当位置渲染由 tablerow 添加的字段
			- %[1]stimeago <时间>%[1]s：将时间戳渲染为相对于当前时间
			- %[1]stimefmt <格式> <时间>%[1]s：使用 Go 的 %[1]sTime.Format%[1]s 函数格式化时间戳
			- %[1]struncate <长度> <输入>%[1]s：确保输入不超过指定长度
			- %[1]shyperlink <URL> <文本>%[1]s：渲染终端超链接

			以下 Sprig 模板库函数也可与此格式化指令配合使用：
			- %[1]scontains <参数> <字符串>%[1]s：检查 %[1]s字符串%[1]s 是否包含 %[1]s参数%[1]s
			- %[1]shasPrefix <前缀> <字符串>%[1]s：检查 %[1]s字符串%[1]s 是否以 %[1]s前缀%[1]s 开头
			- %[1]shasSuffix <后缀> <字符串>%[1]s：检查 %[1]s字符串%[1]s 是否以 %[1]s后缀%[1]s 结尾
			- %[1]sregexMatch <正则> <字符串>%[1]s：检查 %[1]s字符串%[1]s 是否匹配 %[1]s正则%[1]s

			有关 Sprig 库的更多信息，请访问 <https://masterminds.github.io/sprig/>。

			要了解有关 Go 模板的更多信息，请访问：<https://golang.org/pkg/text/template/>。
		`, "`"),
		example: heredoc.Doc(`
			# 默认输出格式
			$ gh pr list
			Showing 23 of 23 open pull requests in cli/cli

			#123  A helpful contribution          contribution-branch              about 1 day ago
			#124  Improve the docs                docs-branch                      about 2 days ago
			#125  An exciting new feature         feature-branch                   about 2 days ago


			# 添加 --json 标志和字段名称列表
			$ gh pr list --json number,title,author
			[
			  {
			    "author": {
			      "login": "monalisa"
			    },
			    "number": 123,
			    "title": "A helpful contribution"
			  },
			  {
			    "author": {
			      "login": "codercat"
			    },
			    "number": 124,
			    "title": "Improve the docs"
			  },
			  {
			    "author": {
			      "login": "cli-maintainer"
			    },
			    "number": 125,
			    "title": "An exciting new feature"
			  }
			]


			# 添加 --jq 标志并从数组中选择字段
			$ gh pr list --json author --jq '.[].author.login'
			monalisa
			codercat
			cli-maintainer


			# --jq 可用于实现更复杂的过滤和输出变换
			$ gh issue list --json number,title,labels --jq \
			  'map(select((.labels | length) > 0))    # 必须有标签
			  | map(.labels = (.labels | map(.name))) # 只显示标签名称
			  | .[:3]                                 # 选择前 3 个结果'
			  [
			    {
			      "labels": [
			        "enhancement",
			        "needs triage"
			      ],
			      "number": 123,
			      "title": "A helpful contribution"
			    },
			    {
			      "labels": [
			        "help wanted",
			        "docs",
			        "good first issue"
			      ],
			      "number": 125,
			      "title": "Improve the docs"
			    },
			    {
			      "labels": [
			        "enhancement",
			      ],
			      "number": 7221,
			      "title": "An exciting new feature"
			    }
			  ]


			# 使用 --template 标志和超链接辅助函数
			$ gh issue list --json title,url --template '{{range .}}{{hyperlink .url .title}}{{"\n"}}{{end}}'


			# 添加 --template 标志并修改显示格式
			$ gh pr list --json number,title,headRefName,updatedAt --template \
				'{{range .}}{{tablerow (printf "#%v" .number | autocolor "green") .title .headRefName (timeago .updatedAt)}}{{end}}'

			#123  A helpful contribution      contribution-branch       about 1 day ago
			#124  Improve the docs            docs-branch               about 2 days ago
			#125  An exciting new feature     feature-branch            about 2 days ago


			# 使用 --template 标志的更复杂示例，使用多个带标题的表格格式化 pull request
			$ gh pr view 3519 --json number,title,body,reviews,assignees --template \
			'{{printf "#%v" .number}} {{.title}}

			{{.body}}

			{{tablerow "ASSIGNEE" "NAME"}}{{range .assignees}}{{tablerow .login .name}}{{end}}{{tablerender}}
			{{tablerow "REVIEWER" "STATE" "COMMENT"}}{{range .reviews}}{{tablerow .author.login .state .body}}{{end}}
			'

			#3519 Add table and helper template functions

			Resolves #3488

			ASSIGNEE  NAME
			mislav    Mislav Marohnić


			REVIEWER  STATE              COMMENT
			mislav    COMMENTED          This is going along great! Thanks for working on this ❤️
		`),
	},
	{
		name:  "exit-codes",
		short: "gh 使用的退出码",
		long: heredoc.Doc(`
			gh 遵循关于退出码的常规约定。

			- 如果命令成功完成，退出码为 0

			- 如果命令因任何原因失败，退出码为 1

			- 如果命令正在运行但被取消，退出码为 2

			- 如果命令需要身份验证，退出码为 4

			注意：特定命令可能有更多的退出码，因此如果你依赖退出码来控制某些行为，
			建议查阅该命令的文档。
		`),
	},
}

func NewCmdHelpTopic(ios *iostreams.IOStreams, ht helpTopic) *cobra.Command {
	cmd := &cobra.Command{
		Use:     ht.name,
		Short:   ht.short,
		Long:    ht.long,
		Example: ht.example,
		Hidden:  true,
		Annotations: map[string]string{
			"markdown:generate": "true",
			"markdown:basename": "gh_help_" + ht.name,
		},
	}

	cmd.SetUsageFunc(func(c *cobra.Command) error {
		return helpTopicUsageFunc(ios.ErrOut, c)
	})

	cmd.SetHelpFunc(func(c *cobra.Command, _ []string) {
		helpTopicHelpFunc(ios.Out, c)
	})

	return cmd
}

func helpTopicHelpFunc(w io.Writer, command *cobra.Command) {
	fmt.Fprint(w, command.Long)
	if command.Example != "" {
		fmt.Fprintf(w, "\n\nEXAMPLES\n")
		fmt.Fprint(w, text.Indent(command.Example, "  "))
	}
}

func helpTopicUsageFunc(w io.Writer, command *cobra.Command) error {
	fmt.Fprintf(w, "Usage: gh help %s", command.Use)
	return nil
}
