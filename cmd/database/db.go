package database

import (
	"github.com/cloudquery/sqlite"
	"github.com/teamssix/cf/pkg/util"
	"github.com/teamssix/cf/pkg/util/cmdutil"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

var GlobalDataBase *GlobalDB

type GlobalDB struct {
	MainDB *gorm.DB
	// if you need memcache or other data storage service. We can easily add at there.
}

func init() {
	InitService()
}

func InitService() {
	dbList := new(GlobalDB)
	dbList.MainDB = InitDB()
	//
	GlobalDataBase = dbList
}

// InitDB initializes the database for user.
func InitDB() *gorm.DB {
	home, err := cmdutil.GetCFHomeDir()
	util.HandleErr(err)
	configHomeFile := filepath.Join(home, "cache")
	if cmdutil.FileExists(configHomeFile) == false {
		err = os.MkdirAll(configHomeFile, 0700)
		util.HandleErr(err)
	}
	db, err := gorm.Open(sqlite.Open(filepath.Join(home, "local.db")), &gorm.Config{})
	return db
}
