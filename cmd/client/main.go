package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"

	"github.com/FengGeSe/terminal/cmd/client/env"
	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/model"
	"github.com/FengGeSe/terminal/util"
)

func main() {
	// 1. 创建root cmd
	rootCmd := NewRootCmd()
	// 2. 输出重定向
	var output []byte
	buf := bytes.NewBuffer(output)
	rootCmd.SetOutput(buf)

	// 3. 执行命令
	if len(os.Args) == 1 {
		RemoteExecute()
		return
	}
	err := rootCmd.Execute()
	if err != nil {
		if IsUnknownCommand(err) {
			// 不是本地命令，发往服务器
			RemoteExecute()
			return
		} else {
			panic(err)
		}
	}

	fmt.Println(buf.String())
}

// 发送参数到服务器执行命令
func RemoteExecute() {
	// 获得当前配置
	cfg, err := conf.GetConfig()
	if err != nil {
		panic(err)
		return
	}

	env := cfg.CurrentEnv()
	u := env.Server + "/v1/execute"

	args := os.Args
	args[0] = env.ToYaml()

	vals := url.Values{}
	vals["args"] = args

	data, err := util.Post(u, vals)
	if err != nil {
		panic(err)
	}

	rstStr := gjson.GetBytes(data, "data").String()
	result := model.LoadExecuteResult([]byte(rstStr))
	fmt.Print(result.Text)
}

// 是否是unknown command 报错
func IsUnknownCommand(err error) bool {
	pattern := `unknown command "[a-zA-Z]*" for "[a-zA-Z]*"`
	r, _ := regexp.Compile(pattern)
	s := err.Error()
	find := r.FindString(s)
	return len(find) == len(s)
}

// Root Cmd
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "client terminal",
		Long:  `client terminal`,
	}
	// 设置全局命令参数
	cmd.PersistentFlags().BoolP("verbose", "v", false, "show more info")

	// 添加子命令
	childCommands := []*cobra.Command{
		env.NewEnvCmd(), // env
	}
	cmd.AddCommand(childCommands...)

	return cmd
}
