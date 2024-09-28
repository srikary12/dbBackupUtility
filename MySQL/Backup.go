package mysql

import (
	"database/sql"
	"dbBackupUtility/globals"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-sql-driver/mysql"
)

func Backup(args globals.DbArgs) {
	cfg := mysql.Config{
		User:   args.Username,
		Passwd: args.Password,
		Net:    "tcp",
		Addr:   args.Host.IP.String() + ":" + args.Port,
		DBName: args.DbName,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	var cmd *exec.Cmd
	if len(args.TableNames) == 0 {
		cmd = exec.Command("mysqldump",
			"--host", args.Host.IP.String(),
			"--user", args.Username,
			"--password="+args.Password,
			"-P", string(args.Port),
			args.DbName)
	} else {
		cmd = exec.Command("mysqldump", "-host", args.Host.IP.String(), "-u", args.Username, "--port", string(args.Port), "-p", args.Password)
		cmd.Args = append(cmd.Args, args.TableNames...)
	}
	//todo: add option to select the folder and compression
	outfile, err := os.Create("backup.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	cmd.Stdout = outfile
	err = cmd.Start(); if err != nil {
		log.Fatal(err)
	}

	cmd.Wait()
	fmt.Println("Backup Completed")
}

//Restore
