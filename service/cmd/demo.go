package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/FengGeSe/terminal/model"
	"github.com/FengGeSe/terminal/util"
)

type TestOpts struct {
	Name string `flag:"name" shorthand:"n" default:"" desc:"名字"`
}

// Cmd
func NewTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "测试用命令",
		Long:  `测试用命令`,
		Run:   testRun,
	}

	// flags
	util.SetFlagsByStruct(cmd.Flags(), TestOpts{})
	return cmd
}

// Run
func testRun(cmd *cobra.Command, args []string) {
	opts := &TestOpts{}
	err := util.SetValuesFromFlags(cmd.Flags(), opts)
	if err != nil {
		cmd.Println(util.WrapRed(err))
		return
	}

	if opts.Name == "" {
		cmd.Println(util.WrapRed("姓名(--name)不能为空"))
		cmd.Help()
		return
	}

	rst := model.NewExecuteResult()

	rst.Text = fmt.Sprintf("Hello %s\n", opts.Name)
	rst.Data["name"] = opts.Name
	cmd.Println(rst)
}
