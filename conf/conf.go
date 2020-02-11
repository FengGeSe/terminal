package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/FengGeSe/terminal/util"
	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

const (
	HTTPAddr = "0.0.0.0:8080"
)

// 客户端配置文件
type Config struct {
	Current string `yaml:"current" desc:"当前环境"`
	Envs    map[string]*Env
}

// 生成默认配置
func NewDefaultConfig(env *Env) *Config {
	return &Config{
		Current: "default",
		Envs: map[string]*Env{
			"default": env,
		},
	}
}

// 环境配置
// 一个终端可以有多个环境配置，并随意切换
type Env struct {
	Server string `yaml:"server" desc:"服务器地址,例如http://127.0.0.1:8080"`
}

func (e *Env) ToYaml() string {
	data, err := yaml.Marshal(e)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

// 根据env yaml字符串加载一个env对象
func LoadEnv(s string) (*Env, error) {
	env := &Env{}
	err := yaml.Unmarshal([]byte(s), env)
	return env, err
}

// 环境配置是否正确
func (e *Env) IsValidate() error {
	if e.Server == "" {
		return fmt.Errorf("服务器地址(--server)不能为空!")
	}
	return nil
}

// 获取当前环境
func (c *Config) CurrentEnv() *Env {
	return c.Envs[c.Current]
}

// 添加一个环境的配置
func (c *Config) Add(name string, env *Env) {
	if name == "" {
		return
	}
	c.Envs[name] = env
}

// 删除一个环境的配置
func (c *Config) Del(name string) error {
	if name == "" {
		return nil
	}
	if name == "default" {
		return fmt.Errorf("不能删除default环境")
	}
	if _, ok := c.Envs[name]; ok {
		delete(c.Envs, name)
	}
	return nil
}

// 切换环境
func (c *Config) Set(name string) error {
	if name == "" {
		return fmt.Errorf("环境名称不能为空!")
	}
	if _, ok := c.Envs[name]; !ok {
		return fmt.Errorf("环境(%s)不存在!", name)
	}
	c.Current = name
	return nil
}

// 序列化 yaml
func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (c *Config) Show() string {
	raw := c.String()
	underline := util.WrapBlue("*" + c.Current)
	middle := strings.Replace(raw, c.Current, underline, 2)
	return strings.Replace(middle, underline, c.Current, 1)
}

var configName = "config.yml"

// 持久化配置文件
func (c *Config) Save() error {
	configPath := GetConfigPath()
	configFile := GetConfigFilePath()

	_, err := os.Stat(configPath)
	if err != nil {
		// 目录不存在，则创建
		if os.IsNotExist(err) {
			if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// 强制覆盖文件,不存在则创建
	if err := ioutil.WriteFile(configFile, []byte(c.String()), 0664); err != nil {
		return err
	}
	return nil
}

func GetConfigPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".terminal")
}

func GetConfigFilePath() string {
	return path.Join(GetConfigPath(), configName)
}

// 获取配置信息
func GetConfig() (*Config, error) {
	file := GetConfigFilePath()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件出错！%v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("配置文件解析出错！请检查配置文件语法。%v", err)
	}

	return &cfg, nil
}
