package config

import (
	"fmt"
	"strconv"
)

type MySQL struct {
	Database       string
	Host           string
	MaxConnections int
	Password       string
	Port           int
	Username       string
}

func (this *MySQL) ConnectionString() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		this.Username,
		this.Password,
		this.Host,
		strconv.Itoa(this.Port),
		this.Database,
	)
}
