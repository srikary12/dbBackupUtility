package main

import (
	"dbBackupUtility/MySQL"
	"dbBackupUtility/globals"
)

func main() {
	var args globals.DbArgs
	mysql.Backup(args)
}