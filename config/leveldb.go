package config

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func LevelDB() *leveldb.DB {
	// Open the LevelDB database
	dbIPaddress, err := leveldb.OpenFile("./leveldb/", nil)
	if err != nil {
		panic(err)
	}
	return dbIPaddress
}