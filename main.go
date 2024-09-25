package main

import (
	"fmt"

	"dbBackupUtility/MySQL"
	"dbBackupUtility/globals"
)

func main() {
	fmt.Println("Hi Hello")
	var args globals.DbArgs
	args.Backup = "backup"
	args.Port = "27567"
	args.Username = "avnadmin"
	args.Password = "AVNS_-BQQMrmhinZOmbMjSdX"
	args.DbName = "test"
	
	// host := "mysql-80ede4b-srikaryelamanchili12-d747.l.aivencloud.com"
	// ip := net.ParseIP(host) // Try parsing as an IP address
    // args.Host = net.IPAddr{IP: ip} // Assign parsed IP address
	fmt.Println(args)
	mysql.Backup(args)
}