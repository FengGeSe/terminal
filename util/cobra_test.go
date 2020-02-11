package util

import (
	"testing"

	"github.com/spf13/cobra"
)

type UserOptions struct {
	Name  string `flag:"name" shorthand:"n" default:"wushuaishuai" desc:"姓名"`
	Grade string `flag:"grade" shorthand:"" default:"一年级" desc:"年级"`
	Path  string `flag:"path" shorthand:"p" default:"${PWD}" desc:"路径"`

	Other OtherInfo `flag:"other"`
}

type OtherInfo struct {
	Age int `flag:"age" default:"24" shorthand:"a" desc:"年龄"`
}

func TestSetFlagsByStruct(t *testing.T) {
	var testCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	err := SetFlagsByStruct(testCmd.Flags(), UserOptions{})
	if err != nil {
		t.Error(err)
	}

	testCmd.Execute()
}

func TestSetValuesFromFlags(t *testing.T) {
	var testCmd = &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()

			options := &UserOptions{}
			if err := SetValuesFromFlags(cmd.Flags(), options); err != nil {
				return err
			}
			t.Logf("%#v \n", options)

			return nil
		},
	}
	err := SetFlagsByStruct(testCmd.Flags(), UserOptions{})
	if err != nil {
		t.Error(err)
	}

	testCmd.Execute()
}
