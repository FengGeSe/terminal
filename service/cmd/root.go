package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "client terminal",
		Long:  `client terminal`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			return
		},
	}
	// 设置命令参数
	cmd.PersistentFlags().BoolP("verbose", "v", false, "show more info")

	// 添加子命令
	childCommands := []*cobra.Command{
		NewTestCmd(),   // test
		NewEnvCmd(),    // env
		NewUpdateCmd(), // update
	}
	cmd.AddCommand(childCommands...)

	// SetHelpCmd(cmd)
	return cmd
}

// 用来占位的
func NewEnvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "manage env config",
		Long:  `manage env config`,
		Run:   func(c *cobra.Command, args []string) {},
	}
}
func NewUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "update client self",
		Long:  `update client self`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}

// 设置帮助命令
func SetHelpCmd(c *cobra.Command) {
	var helpCommand = &cobra.Command{
		Use:   "help [command]",
		Short: "show help info",
		Long: `Help provides help for any command in the application.
Simply type ` + c.Name() + ` help [path to command] for full details.`,

		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Printf("Unknown help topic %#q\n", args)
				c.Root().Usage()
			} else {
				cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
				cmd.Help()
			}
		},
	}
	c.SetHelpCommand(helpCommand)
}
