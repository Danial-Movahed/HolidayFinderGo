package main

import (
	"database/sql"
	"fmt"

	grpc "Danial-Movahed.github.io/clientServerGrpc"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DB struct {
	connection   *sql.DB
	DBhost       string
	DBport       int
	DBuser       string
	DBpassword   string
	DBname       string
	SSLMode      string
	maxTableSize int
}

func (db *DB) ReadFromConfig() {
	viper.SetConfigName("dbconfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file: %s", err))
	}

	db.DBhost = viper.GetString("host")
	db.DBport = viper.GetInt("port")
	db.DBuser = viper.GetString("username")
	db.DBpassword = viper.GetString("password")
	db.DBname = viper.GetString("name")
	db.SSLMode = viper.GetString("SSLmode")
	db.maxTableSize = viper.GetInt("maxTableSize")

}

func (db *DB) Connect() error {
	db.ReadFromConfig()
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
		return grpc.Holiday{
			Name:        name,
			Description: description,
		}, err

	}
	holiday, err := get_holiday_request(HolidayRequest{
		Day:   req.GetDay(),
		Month: req.GetMonth(),
		Year:  req.GetYear()})
	if err != nil {
		return grpc.Holiday{}, err
	}
	tmp, err := db.registerHoliday(&date, holiday)
	if err != nil {
		return grpc.Holiday{}, err
	}
	return grpc.Holiday{
		Name:        tmp.Name,
		Description: tmp.Description,
	}, nil
}

func (db *DB) registerHoliday(date *string, hol Holiday) (grpc.Holiday, error) {
	registerQuery := "INSERT INTO holidays(date, name, description) VALUES ($1, $2, $3)"
	res, err := db.connection.Exec(registerQuery, date, hol.Name, hol.Description)
	if err != nil {
		return grpc.Holiday{}, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return grpc.Holiday{}, err
	}
	numOfRows, err := db.checkNumberOfHolidays()
	if err != nil {
		return grpc.Holiday{}, err
	}
	fmt.Println(numOfRows)
	fmt.Println("found!")
	return grpc.Holiday{
		Name:        hol.Name,
		Description: hol.Description,
	}, err
}

func (db *DB) checkNumberOfHolidays() (int, error) {
	selectionQuery := "SELECT COUNT(*) FROM holidays"
	var count int
	err := db.connection.QueryRow(selectionQuery).Scan(&count)
	if err != nil {
		return count, nil
	}
	if count == db.maxTableSize {
		deletionQuery := "DELETE FROM holidays LIMIT 1"
		_, err := db.connection.Exec(deletionQuery)
		if err != nil {
			return count, err
		}
		err = db.connection.QueryRow(selectionQuery).Scan(&count)
		if err != nil {
			return count, nil
		}

		return count, err
	}
	return count, err
}

var DBConnection = DB{}
