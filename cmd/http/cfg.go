package main

type Config struct {
	DB       string
	ApiScope string
}

func NewConfig() Config {
	return Config{
		DB: "dbname=complex_archive_test user=i.samoylov",
	}
}
