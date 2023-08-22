package dbMan

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type DbManager struct {
	DB *sql.DB
}

var Con *DbManager = nil

func New(path string) (*DbManager, error) {
	if Con != nil {
		return Con, nil
	}
	open, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	//defer func(open *sql.DB) {
	//	println("closing dbMan")
	//	err := open.Close()
	//	if err != nil {
	//		return
	//	}
	//}(open)
	manager := DbManager{DB: open}
	Con = &manager
	return &manager, nil
}
