package main

import (
	"fmt"

	"dbBackupUtility/MySQL"
	"dbBackupUtility/globals"
)

func main() {
	fmt.Println("Hi Hello")
	var args globals.DbArgs
	fmt.Println(args)
	mysql.Backup(args)
}