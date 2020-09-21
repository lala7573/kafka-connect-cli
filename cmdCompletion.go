package main

import (
	"os"

	"github.com/spf13/cobra"
)

var completionTmpl = `To load completions:
  Bash:
  $ source <(kafka-connect-cli completion bash)
  # To load completions for each session, execute once:
  Linux:
    $ kafka-connect-cli completion bash > /etc/bash_completion.d/kafka-connect-cli
  MacOS:
    $ kafka-connect-cli completion bash > /usr/local/etc/bash_completion.d/kafka-connect-cli
  Zsh:
  # If shell completion is not already enabled in your environment you will need
  # to enable it.  You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  # To load completions for each session, execute once:
  $ kafka-connect-cli completion zsh > "${fpath[1]}/_kafka-connect-cli"
  # You will need to start a new shell for this setup to take effect.
  Fish:
  $ kafka-connect-cli completion fish | source
  # To load completions for each session, execute once:
  $ kafka-connect-cli completion fish > ~/.config/fish/completions/kafka-connect-cli.fish
  `

var cmdCompletion = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  completionTmpl,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}