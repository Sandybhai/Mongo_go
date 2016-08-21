package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

type userInfo struct {
  Username string `bson:"username"`           // Has to start with an uppercase else fails
  Pass     string `bson:"pass"`
}

func connect() (session *mgo.Session) {
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

func main() {
	session := connect()
  	defer session.Close()
	var results []userInfo
	collection := session.DB("sandyman").C("userslist")

	err := collection.Insert(&userInfo{Username: "Alex", Pass: "cool"})
	if err != nil {
		log.Println("Could not insert")
		log.Println(err)
	}

	errr := collection.Find(nil).All(&results)
        if errr != nil {
                log.Println("Not in db")
		log.Println(errr)
                return
        }
	fmt.Println("Results All: ", results)
        
}
