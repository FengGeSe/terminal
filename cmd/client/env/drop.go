package env

import (
	"github.com/spf13/cobra"

	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/util"
)

// Cmd
func NewDropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop [env]",
		Short: "移除环境",
		Long:  `移除环境`,
		Run:   dropRun,
	}
	return cmd
}

// Run
func dropRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Println(util.WrapRed("请指定要移除的环境!"))
		cmd.Help()
		return
	}
	cfg, err := conf.GetConfig()
	if err != nil {
		cmd.Print(util.WrapRed(err))
		return
	}

	envName := args[0]
	if err := cfg.Del(envName); err != nil {
		cmd.Print(util.WrapRed(err))
		cmd.Help()
		return
	}

	if err := cfg.Save(); err != nil {
		cmd.Print(util.WrapRed(err))
		cmd.Help()
		return
	}
	cmd.Println(util.WrapGreen("success"))
	cmd.Println(conf.GetConfigFilePath())
	cmd.Println()
	cmd.Println(cfg.Show())
}
