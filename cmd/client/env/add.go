package env

import (
	"github.com/spf13/cobra"

	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/util"
)

// flags
type ConfigOpts struct {
	Server string `flag:"server" shorthand:"s" default:"http://127.0.0.1:8080" desc:"服务器地址,例如http://127.0.0.1:8080"`
}

// Cmd
func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [env]",
		Short: "add a new env",
		Long:  `add a new env`,
		Run:   addRun,
	}

	// flags
	util.SetFlagsByStruct(cmd.Flags(), ConfigOpts{})
	return cmd
}

// Run
func addRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Println(util.WrapRed("请指定要添加的环境名!"))
		cmd.Help()
		return
	}

	// options
	opts := &ConfigOpts{}
	if err := util.SetValuesFromFlags(cmd.Flags(), opts); err != nil {
		cmd.Print(util.WrapRed(err))
		return
	}

	if opts.Server == "" {
		cmd.Println(util.WrapRed("--server(服务器地址)是必填字段"))
		cmd.Help()
		return
	}

	// set config
	cfg, err := conf.GetConfig()
	if err != nil {
		cmd.Print(util.WrapRed(err))
		return
	}
	envName := args[0]
	env := &conf.Env{
		Server: opts.Server,
	}
	cfg.Add(envName, env)
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
