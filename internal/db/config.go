package db

import "github.com/go-pg/pg/v9"

type Config struct {
	DB  *pg.Options
	TLS struct {
		Enabled            bool
		CAFilePath         string
		InsecureSkipVerify bool
		ServerName         string
	}
}
