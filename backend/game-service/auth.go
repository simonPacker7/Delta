package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func extractSessionCookie(req *http.Request) (string, error) {
	data, err := req.Cookie("session_id")
	if err != nil {
		return "", err
	}
	return data.Value, nil
}

func getSessionFromStore(rdb *redis.Client, session_id string) (string, error) {
	data, err := rdb.Get(context.Background(), session_id).Bytes()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	sessionData := make(map[string]interface{})

	if err := dec.Decode(&sessionData); err != nil {
		log.Fatal(err)
		return "", err
	}

	userId := sessionData["id"].(string)
	userName := sessionData["name"].(string)
	log.Println("getSessionFromStore() userId:", userId, "userName:", userName)

	return userId, nil
}

func authenticateRequest(rdb *redis.Client, req *http.Request) (string, error) {
	log.Println("authenticateRequest called")
	session_id, err := extractSessionCookie(req)
	if err != nil {
		return "", err
	}

	userId, err := getSessionFromStore(rdb, session_id)
	if err != nil {
		return "", err
	}

	return userId, nil
}
