package main

import (
	. "controllers"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	. "golanger.com/database/activerecord"
	. "golanger.com/middleware"
	"os"
	"path/filepath"
	"runtime"
	_ "templateFunc"
)

var (
	addr       = flag.String("addr", ":80", "Server port")
	configPath = flag.String("config", "./config/site", "site filepath of config")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()*2 + 1)

	flag.Parse()
	os.Chdir(filepath.Dir(os.Args[0]))
	fmt.Println("Listen server address: " + *addr)
	fmt.Println("Read configuration file success, fithpath: " + filepath.Join(filepath.Dir(os.Args[0]), *configPath))

	App.Load(*configPath)

	if sqliteDns, ok := App.Database["Sqlite"]; ok && sqliteDns != "" {
		sqlite, err := sql.Open("sqlite3", sqliteDns)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		defer sqlite.Close()
		orm := NewActiveRecord(sqlite)
		Middleware.Add("orm", orm)
		Middleware.Add("db", sqlite)
	}

	App.AddHeader("Content-Type", "text/html; charset=utf-8")
	App.HandleFavicon()
	App.HandleStatic()
	App.ListenAndServe(*addr, App)
}
