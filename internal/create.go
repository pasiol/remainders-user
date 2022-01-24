package internal

import (
	"context"
	"errors"
	"fmt"
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

func CreateNewUser(u, p, uri string) error {
	hashedPassword, err := hashAndSalt(p)
	if err != nil {
		return err
	}
	user := bson.D{{"_id", u}, {"username", u}, {"password", hashedPassword}, {"approved", true}, {"createdAt", time.Now()}, {"updatedAt", time.Now()}}
	db, client := ConnectOrDie(uri)
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
