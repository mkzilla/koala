package initial

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/mkzilla/koala/pkg/types/config"
)

func InitDB(driverName, dataSourceName string, debug bool) {
	var err error
	config.DBEngine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	config.DBEngine.ShowSQL(debug)
}
