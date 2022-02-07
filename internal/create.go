package internal

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func getDbName(uri string) (string, error) {

	begining, err := regexp.Compile(`(mongodb([+]srv:|)\/\/(\S*):(\S*)@(\S*)\/)`)
	if err != nil {
		return "", err
	}

	match := begining.FindAllString(uri, -1)
	if len(match) > 0 {

		dbName := strings.Replace(uri, match[0], "", 1)
		ending, err := regexp.Compile(`\?(\S*)`)
		if err != nil {
			return "", err
		}
		match = ending.FindAllString(dbName, -1)
		if len(match) > 0 {
			dbName = strings.Replace(string(dbName), match[0], "", 1)
		}

		return string(dbName), nil
	}
	return "", errors.New("cannot extract db name")
}

func CreateNewUser(u, p, uri string) error {
	hashedPassword, err := hashAndSalt(p)
	if err != nil {
		return err
	}
	dbName, err := getDbName(uri)
	if err != nil {
		return errors.New("parsing database from uri failed")
	}
	user := bson.D{{"_id", u}, {"username", u}, {"password", hashedPassword}, {"approved", true}, {"createdAt", time.Now()}, {"updatedAt", time.Now()}}
	db, client := ConnectOrDie(uri, dbName)
	defer client.Disconnect(context.Background())
	results, err := db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		if strings.Contains(err.Error(), "E11000 duplicate key error") {
			return errors.New("username already exists")
		}
		return err
	}
	fmt.Printf("User created succesfully: %v", results)
	return nil
}
