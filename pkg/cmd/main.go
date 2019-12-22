package main

import (
	"github.com/spf13/cobra"

	"github.com/mkzilla/koala/pkg/types/config"
)

var (
	Version = "0.1.0"
)

// 是否输出 version 信息
var v bool

// RootCmd 是命令行的入口
var RootCmd = &cobra.Command{
	Use: "koala",
	Long: `
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			versionE(cmd, args)
		}
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd, StartCmd)
	cobra.EnableCommandSorting = false
	RootCmd.Flags().BoolVarP(&v, "version", "v", false, "show version")
}

func main() {
	config.Version = Version
	err := RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
