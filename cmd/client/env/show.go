package env

import (
	"github.com/spf13/cobra"

	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/util"
)

// Cmd
func NewShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "查看环境",
		Long:  `查看环境`,
		Run:   showRun,
	}

	return cmd
}

// Run
func showRun(cmd *cobra.Command, args []string) {
	// 读取环境配置
	cfg, err := conf.GetConfig()
	if err != nil {
		cmd.Println(util.WrapRed(err.Error()))
		return
	}

	cmd.Println(conf.GetConfigFilePath())
	cmd.Println()
	cmd.Println(cfg.Show())
}
