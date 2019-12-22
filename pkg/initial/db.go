package initial

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/mkzilla/koala/pkg/types"
	"github.com/mkzilla/koala/pkg/types/config"
	"github.com/mkzilla/koala/pkg/utils"
)

func InitDB() {
	var err error
	config.DBEngine, err = xorm.NewEngine(config.Configs.DataSource.Driver, config.Configs.DataSource.Source)
	if err != nil {
		panic(err)
	}
	config.DBEngine.ShowSQL(config.Configs.DataSource.Debug)
	if config.Configs.DataSource.Init {
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
			ID:           1,
			Username:     "admin",
			Nickname:     "管理员",
			Email:        "792400644@qq.com",
			GlobalUnique: "koala-1",
			Phone:        "132-317-44585",
			Enable:       true,
			IsAdmin:      true,
		})
		config.DBEngine.InsertOne(types.User{
			ID:           2,
			Username:     "messi",
			Nickname:     "梅西",
			GlobalUnique: "koala-2",
			Email:        "792400644@qq.com",
			Phone:        "800-810-3099",
			Enable:       true,
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
	}
}
