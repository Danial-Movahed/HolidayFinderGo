package main

import (
	"database/sql"
	"fmt"

	grpc "Danial-Movahed.github.io/clientServerGrpc"
	_ "github.com/lib/pq"
)

type DB struct {
	connection *sql.DB
	DBhost     string
	DBport     int
	DBuser     string
	DBpassword string
	DBname     string
	SSLMode    string
}

func (db *DB) Connect() error {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", db.DBhost, DBport, db.DBname, db.DBuser, db.DBpassword, db.SSLMode)
	connection, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = connection.Ping()
	if err != nil {
		return err
	}

	db.connection = connection
	return nil
}

func (db *DB) Close() error {
	return db.connection.Close()
}

func (db *DB) GetHoliday(req *grpc.HolidayRequest) (grpc.Holiday, error) {
	date := fmt.Sprintf("%s-%s-%s", req.GetYear(), req.GetMonth(), req.GetDay())

	selectionQuery := `SELECT FROM Holidays WHERE date = "$1"`
	rows, err := db.connection.Query(selectionQuery, date)

	if err != nil {
		return grpc.Holiday{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var name, description string
		if err := rows.Scan(&name, &description); err != nil {
			return grpc.Holiday{}, err
		}
		if name == "" && description == "" {
			return grpc.Holiday{}, fmt.Errorf("no holidays on %s", date)
		}
		fmt.Printf("Holiday name: %s\nHoliday description: %s", name, description)
		return grpc.Holiday{
			Name:        name,
			Description: description,
		}, err

	} else {
		// If no holiday is found reports to client and asks for new information
		// Then saves to Database and returns results to gRtc server
		return grpc.Holiday{}, err
	}

}

var DBConnection = DB{DBhost: DBhost, DBport: DBport, DBuser: DBuser, DBpassword: DBpassword, DBname: DBname}

// func (db *DB) registerHoliday() {

// }
