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
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", db.DBhost, db.DBport, db.DBname, db.DBuser, db.DBpassword, db.SSLMode)
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
	selectionQuery := "SELECT * FROM holidays WHERE date = $1"
	rows, err := db.connection.Query(selectionQuery, date)

	if err != nil {
		return grpc.Holiday{}, err
	}
	defer rows.Close()
	if rows.Next() {
		var tmp, name, description string
		if err := rows.Scan(&tmp, &name, &description); err != nil {
			return grpc.Holiday{}, err
		}
		if name == "" && description == "" {
			return grpc.Holiday{Name: "Nothing", Description: "No holidays on this date!"}, err
		}
		return grpc.Holiday{
			Name:        name,
			Description: description,
		}, err

	} else {
		holiday := get_holiday_request(HolidayRequest{
			Day:   req.GetDay(),
			Month: req.GetMonth(),
			Year:  req.GetYear(),
		})
		tmp, err := db.registerHoliday(&date, holiday)
		if err != nil {
			return grpc.Holiday{}, err
		} else {
			return grpc.Holiday{
				Name:        tmp.Name,
				Description: tmp.Description,
			}, err
		}
	}

}

func (db *DB) registerHoliday(date *string, hol Holiday) (grpc.Holiday, error) {
	registerQuery := "INSERT INTO holidays(date, name, description) VALUES ($1, $2, $3)"
	res, err := db.connection.Exec(registerQuery, date, hol.Name, hol.Description)
	if err != nil {
		return grpc.Holiday{}, err
	} else {
		_, err := res.RowsAffected()
		if err != nil {
			return grpc.Holiday{}, err
		}
		return grpc.Holiday{
			Name:        hol.Name,
			Description: hol.Description,
		}, err
	}
}

var DBConnection = DB{DBhost: DBhost, DBport: DBport, DBuser: DBuser, DBpassword: DBpassword, DBname: DBname}
