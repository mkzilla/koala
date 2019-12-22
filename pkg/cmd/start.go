package main

import (
	"github.com/robfig/cron"
	"github.com/spf13/cobra"

	"github.com/mkzilla/koala/pkg/initial"
	//"github.com/mkzilla/koala/pkg/jobs"
	"github.com/mkzilla/koala/pkg/types/config"
)

type DB struct {
	DriverName string
	DataSource string
}

type StartOption struct {
	Debug  bool
	Config string
}

var (
	startOption StartOption
	StartCmd    = &cobra.Command{
		Use:     "start",
		Short:   ``,
		Long:    ``,
		Example: `koala start`,
		Run: func(cmd *cobra.Command, args []string) {
			initial.InitConfig(startOption.Config)
			initial.InitDB()
			initial.InitOAuth()
			initial.InitRsaKey()
			initial.InitRouter()

			c := cron.New()
			c.Start()
			config.GinEngine.Run(config.Configs.Addr)
		},
	}
)

func init() {
	StartCmd.Flags().SortFlags = false
	StartCmd.Flags().StringVarP(&startOption.Config, "config", "c", "", "")
}
