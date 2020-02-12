package update

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/FengGeSe/terminal/conf"
	"github.com/FengGeSe/terminal/util"
	"github.com/spf13/cobra"
)

// Cmd
// 环境配置相关命令
func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update client self",
		Long:  `update client self`,
		Run: func(cmd *cobra.Command, args []string) {
			config, err := conf.GetConfig()
			if err != nil {
				cmd.Print(util.WrapRed(err))
				return
			}
			url := config.CurrentEnv().Server + "/file/cli"
			for _, p := range FindYourself() {
				cmd.Println(p)
				if err := SaveFile(url, p); err != nil {
					cmd.Print(util.WrapRed(err))
					return
				}
			}
		},
	}
	return cmd
}

func SaveFile(url, path string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0731)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, res.Body)
	return err
}

func FindYourself() []string {
	firstArg := os.Args[0]
	lastIndex := strings.LastIndex(firstArg, string(os.PathSeparator))
	result := []string{}
	if lastIndex > 1 {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err)
			return result
		}
		yourself := path.Join(dir, firstArg[lastIndex+1:])
		result = append(result, yourself)
	} else {
		// find cli from $PATH
		pwd := os.Getenv("PATH")
		dirs := strings.Split(pwd, ":")
		for _, dirname := range dirs {
			fileInfos, err := ioutil.ReadDir(dirname)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, info := range fileInfos {
				if info.Name() == firstArg {
					yourself := path.Join(dirname, firstArg)
					result = append(result, yourself)
					break
				}
			}
		}
	}
	return result
}
