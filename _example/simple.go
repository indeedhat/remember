package main

import (
	"github.com/indeedhat/remember"
	_ "github.com/mattn/go-sqlite3"
	"path/filepath"
	"fmt"
)

func main() {
	dns, err := filepath.Abs("simple.db")
	if nil != err {
		panic(err)
	}

	driver, err := remember.LoadDriver("../driver/sqlite/sqlite.so", dns)
	if nil != err {
		panic(err)
	}
	fmt.Println(driver)

	fmt.Println(driver.Close())
}

