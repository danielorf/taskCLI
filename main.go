package main

import (
	"fmt"
	"os"

	"github.com/danielorf/task/cmd"
	"github.com/danielorf/task/db"
)

func main() {
	// home, _ := homedir.Dir()
	// dbPath := filepath.Join(home, "tasks.db")
	// must(db.Init(dbPath))
	must(db.Init("/tmp/tasks.db"))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
