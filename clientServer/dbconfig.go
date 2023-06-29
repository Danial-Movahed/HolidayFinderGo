package main

type dbConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	Name         string
	SSLmode      string
	maxTableSize int
}

var dbconfig dbConfig

var (
	DBhost         = dbconfig.Host
	DBport         = dbconfig.Port
	DBuser         = dbconfig.Username
	DBpassword     = dbconfig.Password
	DBname         = dbconfig.Name
	SSLMode        = dbconfig.SSLmode
	DBmaxTableSize = dbconfig.maxTableSize
)
