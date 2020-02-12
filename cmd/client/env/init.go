package env

import (
	"github.com/spf13/cobra"

	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/util"
)

// flags
type InitConfigOpts struct {
	Server string `flag:"server" shorthand:"s" desc:"服务器地址,例如http://127.0.0.1:8080"`
}

// Cmd
func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init config",
		Long:  `init config`,
		Run:   initRun,
	}

	util.SetFlagsByStruct(cmd.Flags(), InitConfigOpts{})
	return cmd
}

// Run
func initRun(cmd *cobra.Command, args []string) {
	opts := &InitConfigOpts{}
	if err := util.SetValuesFromFlags(cmd.Flags(), opts); err != nil {
		cmd.Print(util.WrapRed(err))
		cmd.Help()
		return
	}

	if opts.Server == "" {
		cmd.Println(util.WrapRed("--server(服务器地址)是必填字段"))
		cmd.Help()
		return
	}

	env := &conf.Env{
		Server: opts.Server,
	}
	cfg := conf.NewDefaultConfig(env)
	if err := cfg.Save(); err != nil {
		cmd.Print(util.WrapRed(err))
		cmd.Help()
		return
	}
	if err := cfg.Save(); err != nil {
		cmd.Println(util.WrapRed(err))
		return
	}

	cmd.Println(util.WrapGreen("success"))
	cmd.Println(conf.GetConfigFilePath())
	cmd.Println()
	cmd.Println(cfg.Show())
}
