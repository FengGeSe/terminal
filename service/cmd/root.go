package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "客户端",
		Long:  `客户端`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			return
		},
	}
	// 设置命令参数
	cmd.PersistentFlags().BoolP("verbose", "v", false, "打印请求参数")

	// 添加子命令
	childCommands := []*cobra.Command{
		NewEnvCmd(),  // env
		NewTestCmd(), // test
	}
	cmd.AddCommand(childCommands...)

	SetHelpCmd(cmd)
	return cmd
}

// 用来占位的
func NewEnvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "环境配置",
		Long:  `环境配置`,
	}
}

// 设置帮助命令
func SetHelpCmd(c *cobra.Command) {
	var helpCommand = &cobra.Command{
		Use:   "help [command]",
		Short: "查看帮助",
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
