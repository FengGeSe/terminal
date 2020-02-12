package env

import (
	"github.com/spf13/cobra"

	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/util"
)

// Cmd
func NewSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [env]",
		Short: "switch env",
		Long:  `switch env`,
		Run:   setRun,
	}
	return cmd
}

// Run
func setRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Println(util.WrapRed("请指定要切换的环境!"))
		cmd.Help()
		return
	}

	envName := args[0]

	cfg, err := conf.GetConfig()
	if err != nil {
		cmd.Print(util.WrapRed(err))
		return
	}

	// change
	if err := cfg.Set(envName); err != nil {
		cmd.Println(util.WrapRed(err))
		cmd.Help()
		return
	}

	// save
	err = cfg.Save()
	if err != nil {
		cmd.Print(util.WrapRed(err))
		cmd.Help()
		return
	}

	cmd.Println(util.WrapGreen("success"))
	cmd.Println(conf.GetConfigFilePath())
	cmd.Println()
	cmd.Println(cfg.Show())
}
