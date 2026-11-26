package cmd

import (
	"github.com/spf13/cobra"
)

func createCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "completion [bash|zsh|fish|powershell]",
		Aliases: []string{"shell-completion"},
		Short:   "Generate the autocompletion script for the specified shell",
		Long: `To load completions:

Bash:

$ source <(yq completion bash)

# To load completions for each session, execute once:
Linux:
  $ yq completion bash > /etc/bash_completion.d/yq
MacOS:
  $ yq completion bash > /usr/local/etc/bash_completion.d/yq

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ yq completion zsh > "${fpath[1]}/_yq"

# You will need to start a new shell for this setup to take effect.

Fish:

$ yq completion fish | source

# To load completions for each session, execute once:
$ yq completion fish > ~/.config/fish/completions/yq.fish
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error = nil
			switch args[0] {
			case "bash":
				err = cmd.Root().GenBashCompletionV2(cmd.OutOrStdout(), true)
			case "zsh":
				err = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			case "fish":
				err = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			case "powershell":
				err = cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
			}
			return err

		},
	}
}
