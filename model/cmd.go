package model

import (
	"encoding/json"

	"github.com/FengGeSe/terminal/conf"
)

type ExecuteParams struct {
	Args []string  `json:"args"  desc:"命令参数"`
	Env  *conf.Env `json:"env" desc:"客户端环境配置"`
}

func LoadExecuteParams(args []string) (*ExecuteParams, error) {
	envStr := args[0]
	env, err := conf.LoadEnv(envStr)
	if err != nil {
		return nil, err
	}
	return &ExecuteParams{
		Args: args[1:],
		Env:  env,
	}, nil
}

func NewExecuteParams() *ExecuteParams {
	return &ExecuteParams{
		Args: []string{},
		Env:  &conf.Env{},
	}
}

type ExecuteResult struct {
	Text string            `json:"text" desc:"显示文本"`
	Data map[string]string `json:"data,omitempty" desc:"数据"`
}

func NewExecuteResult() *ExecuteResult {
	return &ExecuteResult{
		Text: "",
		Data: map[string]string{},
	}
}

func (r *ExecuteResult) ToJson() string {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (r *ExecuteResult) String() string {
	return r.ToJson()
}

// 根据data转换成ExecuteResult对象
func LoadExecuteResult(data []byte) *ExecuteResult {
	var result ExecuteResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		// data并不是ExecuteResult的时候，转换成ExecuteResult
		return &ExecuteResult{
			Text: string(data),
			Data: map[string]string{},
		}
	}
	return &result
}
