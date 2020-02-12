package env

import (
	"github.com/spf13/cobra"
)

// Cmd
// 环境配置相关命令
func NewEnvCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "manage env config",
		Long:  `manage env config`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// 添加子命令
	childCommands := []*cobra.Command{
		NewInitCmd(), // init
		NewSetCmd(),  // set
		NewAddCmd(),  // add
		NewShowCmd(), // show
		NewDropCmd(), // drop
	}
	cmd.AddCommand(childCommands...)
	return cmd
}
