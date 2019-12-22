package main

import (
	"github.com/spf13/cobra"

	"github.com/mkzilla/koala/pkg/initial"
	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/utils"
)

type EnvOption struct {
	DB DB
}

var (
	envOption EnvOption
	EnvCmd    = &cobra.Command{
		Use:     "env",
		Short:   ``,
		Long:    ``,
		Example: `koala env`,
		Run: func(cmd *cobra.Command, args []string) {
			initial.InitDB(envOption.DB.DriverName, envOption.DB.DataSource, true)
			config.DBEngine.Sync2(new(types.User))
			config.DBEngine.Sync2(new(types.Password))
			config.DBEngine.Sync2(new(types.Role))
			config.DBEngine.Sync2(new(types.Group))
			config.DBEngine.Sync2(new(types.GroupUser))
			config.DBEngine.Sync2(new(types.Behavior))
			config.DBEngine.Sync2(new(types.Label))
			config.DBEngine.Sync2(new(types.Comment))
			config.DBEngine.Sync2(new(types.Task))
			config.DBEngine.Sync2(new(types.TaskLog))
			config.DBEngine.Sync2(new(types.TaskVote))
			config.DBEngine.Sync2(new(types.TaskWatch))
			config.DBEngine.Sync2(new(types.Notification))
			config.DBEngine.Sync2(new(types.Objective))

			config.DBEngine.InsertOne(types.User{
				ID:       1,
				Username: "admin",
				Nickname: "管理员",
				Email:    "admin@qq.com",
				Phone:    "132-317-44585",
				Enable:   true,
				IsAdmin:  true,
			})
			config.DBEngine.InsertOne(types.User{
				ID:       2,
				Username: "messi",
				Nickname: "梅西",
				Email:    "792400644@qq.com",
				Phone:    "800-810-3099",
				Enable:   true,
			})
			config.DBEngine.InsertOne(types.Password{
				ID:       1,
				UserID:   1,
				Password: utils.Cipher("admin"),
			})
			config.DBEngine.InsertOne(types.Password{
				ID:       2,
				UserID:   2,
				Password: utils.Cipher("goat"),
			})
		},
	}
)

func init() {
	EnvCmd.Flags().SortFlags = false
	EnvCmd.Flags().StringVarP(&envOption.DB.DriverName, "driver", "d", "mysql", "")
	EnvCmd.Flags().StringVarP(&envOption.DB.DataSource, "source", "s", "", "")
}
