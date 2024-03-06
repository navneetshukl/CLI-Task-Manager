package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/navneetshukl/task/cmd"
	"github.com/navneetshukl/task/db"
)

func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "task.db")
	must(db.DBInit(dbPath))

	must(cmd.RootCmd.Execute())

	//fmt.Println(home)
}

// must is a helper function which will handle if any error occured is right way
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
