package models

import (
	"database/sql"

	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	DbUser   string `yaml:"dbUser"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DbName   string `yaml:"dbName"`
}

func MatchKaro(password string, oldhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldhash), []byte(password))
	return err == nil
}

func HashKaro(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	return string(hashedPasswordBytes), err
}

func dsn() (string, error) {

	creds, err := os.Open("config.yaml")
	if err != nil {
		fmt.Printf("Error: '%s' while trying to open config.yaml\n", err)
		return "", err
	}
	defer creds.Close()

	var config YamlConfig
	decoder := yaml.NewDecoder(creds)

	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("Error: '%s' while decoding config.yaml\n", err)
		return "", err
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DbUser,
		config.Password,
		config.Host,
		config.DbName,
	), nil
}

func Connection() (*sql.DB, error) {
	dsn, err := dsn()

	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	
	log.Println("Connected to DB successfully\n")
	return db, err
}
