package db

import (
	"database/sql"
	"fmt"

	grpc "Danial-Movahed.github.io/clientServerGrpc"
	_ "github.com/lib/pq"
)

type DB struct {
	connection *sql.DB
}

type noHolidaysError struct {
	Date string
}

func (e *noHolidaysError) Error() string {
	return fmt.Sprintf("No holidays on %s", e.Date)
}

func (db *DB) Connect() error {
	connStr := fmt.Sprintf("host=%s:%v dbname=%s user=%s password=%s sslmode=disable", host, port, user, password, dbname)
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

func (db *DB) Find(req *grpc.HolidayRequest) (grpc.Holiday, error) {
	date := fmt.Sprintf("%s-%s-%s", req.Year, req.Month, req.Day)

	selectionQuery := "SELECT FROM Holidays WHERE date = '$1'"
	rows, err := db.connection.Query(selectionQuery, date)

	if err != nil {
		return grpc.Holiday{
			Name:        "",
			Description: "",
		}, err
	}

	defer rows.Close()

	if rows.Next() {
		var name, description string
		if err := rows.Scan(&name, &description); err != nil {
			return grpc.Holiday{
				Name:        "",
				Description: "",
			}, err
		}
		if name == "" && description == "" {
			return grpc.Holiday{
					Name:        name,
					Description: description,
				}, &noHolidaysError{
					Date: date,
				}
		}
		return grpc.Holiday{
			Name:        name,
			Description: description,
		}, err

	} else {
		// If no holiday is found reports to client and asks for new information
		// Then saves to Database and returns results to gRtc server
		return grpc.Holiday{
			Name:        "",
			Description: "",
		}, err
	}

}

// func (db *DB) Add() {

// }
