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
	DB    DB
	Addr  string
	Debug bool
}

var (
	startOption StartOption
	StartCmd    = &cobra.Command{
		Use:     "start",
		Short:   ``,
		Long:    ``,
		Example: `koala start`,
		Run: func(cmd *cobra.Command, args []string) {
			initial.InitDB(startOption.DB.DriverName, startOption.DB.DataSource, startOption.Debug)
			initial.InitOAuth()
			initial.InitRsaKey()
			initial.InitRouter()

			c := cron.New()
			c.Start()
			config.GinEngine.Run(startOption.Addr)
		},
	}
)

func init() {
	StartCmd.Flags().SortFlags = false
	StartCmd.Flags().StringVarP(&startOption.DB.DriverName, "driver", "d", "mysql", "")
	StartCmd.Flags().StringVarP(&startOption.DB.DataSource, "source", "s", "", "")

	StartCmd.Flags().StringVarP(&startOption.Addr, "addr", "a", ":9826", "")
	StartCmd.Flags().BoolVarP(&startOption.Debug, "debug", "", true, "")
}
