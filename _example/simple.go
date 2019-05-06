package main

import "github.com/indeedhat/remember"

func main() {
	driver, err := remember.LoadDriver("../driver/sqlite/sqlite.so", "./simple.db")
	if nil != err {
		panic(err)
	}

	driver.Close()
}
