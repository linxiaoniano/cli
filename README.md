# GitHub CLI

`gh` 是 GitHub 的命令行工具。它将 Pull Request、Issue 及其他 GitHub 概念带到终端中，与你正在使用的 `git` 和代码并排工作。

![gh pr status 截图](https://user-images.githubusercontent.com/98482/84171218-327e7a80-aa40-11ea-8cd1-5177fc2d0e72.png)

GitHub CLI 支持 GitHub.com、GitHub Enterprise Cloud 和 GitHub Enterprise Server 2.20+ 用户，支持 macOS、Windows 和 Linux。

## 文档

[安装选项见下文](#安装)，使用说明请[查阅手册](https://cli.github.com/manual/)。

## 贡献

如果你觉得有任何不妥或缺少某些功能，请查看[贡献页面](.github/CONTRIBUTING.md)。那里有关于分享反馈、本地构建工具以及向项目提交 Pull Request 的说明。

如果你是 GitHub 内部员工并有兴趣为 CLI 添加新命令，请查看我们的[内部贡献文档](docs/working-with-us.md)。

<!-- 此锚点被其他位置引用，请避免重命名 -->
## 安装

### [macOS](docs/install_macos.md)

- [Homebrew](docs/install_macos.md#homebrew)
- [预编译二进制](docs/install_macos.md#precompiled-binaries) 见[发布页面][]

更多 macOS 包和安装程序请查看[社区支持的文档](docs/install_macos.md#community-unofficial)

### [Linux 与 Unix](docs/install_linux.md)

- [Debian、Raspberry Pi、Ubuntu](docs/install_linux.md#debian)
- [Amazon Linux、CentOS、Fedora、openSUSE、RHEL、SUSE](docs/install_linux.md#rpm)
- [预编译二进制](docs/install_linux.md#precompiled-binaries) 见[发布页面][]

更多 Linux 与 Unix 包和安装程序请查看[社区支持的文档](docs/install_linux.md#community-unofficial)

### [Windows](docs/install_windows.md)

- [WinGet](docs/install_windows.md#winget)
- [预编译二进制](docs/install_windows.md#precompiled-binaries) 见[发布页面][]

更多 Windows 包和安装程序请查看[社区支持的文档](docs/install_windows.md#community-unofficial)

### 从源码构建

请参阅如何[从源码构建 GitHub CLI](docs/install_source.md)。

### GitHub Codespaces

要将 GitHub CLI 添加到你的 codespace，请在 [devcontainer 文件](https://docs.github.com/en/codespaces/setting-up-your-project-for-codespaces/adding-features-to-a-devcontainer-file)中添加以下内容：

```json
"features": {
  "ghcr.io/devcontainers/features/github-cli:1": {}
}
```

### GitHub Actions

[GitHub 托管的运行器](https://docs.github.com/en/actions/using-github-hosted-runners/about-github-hosted-runners)已预装 GitHub CLI，并每周更新。

如果需要特定版本，你的 GitHub Actions 工作流需要根据上述 [macOS](#macos)、[Linux 与 Unix](#linux--unix) 或 [Windows](#windows) 说明进行安装。

有关所有预装工具的信息，请参阅 [`actions/runner-images`](https://github.com/actions/runner-images)

### 二进制验证

自 2.50.0 版本起，`gh` 开始生成[构建出处证明](https://github.blog/changelog/2024-06-25-artifact-attestations-is-generally-available/)，提供可加密验证的追踪链，追溯到源 GitHub 仓库、git 修订版和使用的构建指令。构建出处证明已签名，并依赖于 Public Good [Sigstore](https://www.sigstore.dev/) 进行 PKI。

根据是否已安装 `gh`，有两种常见的验证下载版本的方法。如果已安装 `gh`，验证新版本非常简单：

- **选项 1：使用已安装的 `gh` 进行验证：**

  ```shell
  $ gh at verify -R cli/cli gh_2.62.0_macOS_arm64.zip
  Loaded digest sha256:fdb77f31b8a6dd23c3fd858758d692a45f7fc76383e37d475bdcae038df92afc for file://gh_2.62.0_macOS_arm64.zip
  Loaded 1 attestation from GitHub API
  ✓ Verification succeeded!

  sha256:fdb77f31b8a6dd23c3fd858758d692a45f7fc76383e37d475bdcae038df92afc was attested by:
  REPO     PREDICATE_TYPE                  WORKFLOW
  cli/cli  https://slsa.dev/provenance/v1  .github/workflows/deployment.yml@refs/heads/trunk
  ```

- **选项 2：使用 Sigstore [`cosign`](https://github.com/sigstore/cosign)：**

  执行此操作时，请下载所下载版本的[证明](https://github.com/cli/cli/attestations)，并使用 cosign 验证下载版本的真实性：

  ```shell
  $ cosign verify-blob-attestation --bundle cli-cli-attestation-3120304.sigstore.json \
        --new-bundle-format \
        --certificate-oidc-issuer="https://token.actions.githubusercontent.com" \
        --certificate-identity="https://github.com/cli/cli/.github/workflows/deployment.yml@refs/heads/trunk" \
        gh_2.62.0_macOS_arm64.zip
  Verified OK
  ```

## 与 hub 的对比

多年来，[hub](https://github.com/github/hub) 一直是非官方的 GitHub CLI 工具。`gh` 是一个新项目，帮助我们探索官方 GitHub CLI 工具在根本不同设计下的可能性。虽然这两个工具都将 GitHub 带到终端，但 `hub` 充当 `git` 的代理，而 `gh` 是一个独立的工具。查看我们的[详细说明](docs/gh-vs-hub.md)以了解更多。

[发布页面]: https://github.com/cli/cli/releases/latest
