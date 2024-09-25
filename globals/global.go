package globals

import (
	"net"
)

type DbArgs struct {
	Backup string
	Port string
	Host net.IPAddr
	Username string
	Password string
	DbName string
	TableNames []string
}