package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Schema default to model scavenger
type ScanvegerSchema struct {
	Name    string `db:"user_name"`
	Address struct {
		City     string `db:"city"`
		District string `db:"district"`
		Street   string `db:"street"`
		Number   string `db:"number_residential"`
		ZipCode  string `db:"zip_code"`
	} `db:"address"`
	Phones []struct {
		PhoneMobile string `db:"mobile_phone"`
		PhoneHome   string `db:"home_phone"`
	} `db:"phones"`
	BirthDate     string `db:"birth_date"`
	DoumentNumber string `db:"doument_number"`
	Email         string `db:"email"`
}

const DATABASE_URL = "postgresql://localhost:54320/?user=root"

func Pool() {
	dbpool, err := pgxpool.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
