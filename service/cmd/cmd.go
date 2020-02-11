package cmd

import (
	"bytes"
	"context"

	"github.com/FengGeSe/terminal/model"
)

// interface
type CmdService interface {
	Execute(context.Context, *model.ExecuteParams) (*model.ExecuteResult, error)
}

func NewCmdService() CmdService {
	var svc CmdService = cmdService{}
	return svc
}

// implement
type cmdService struct{}

var _ CmdService = cmdService{}

func (c cmdService) Execute(ctx context.Context, params *model.ExecuteParams) (*model.ExecuteResult, error) {
	// 1. 生成命令
	rootCmd := NewRootCmd()
	rootCmd.SetArgs(params.Args)

	// 2. 输出定向
	var output []byte
	buf := bytes.NewBuffer(output)
	rootCmd.SetOutput(buf)
	rootCmd.SetErr(buf)

	// 3. 执行命令
	if err := rootCmd.Execute(); err != nil {
		buf.WriteString(err.Error())
		return nil, err
	}

	// 4. 解析结果
	resp := model.LoadExecuteResult(buf.Bytes())
	return resp, nil
}
