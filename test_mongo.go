package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
//	"time"
)

type userInfo struct {
  Username string `bson:"username"`           // Has to start with an uppercase else fails
  Pass     string	`bson:"pass"`
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

  //result := userInfo {}
	//var result interface{}
	var results []userInfo
  collection := session.DB("sandyman").C("userslist")

	err := collection.Insert(&userInfo{Username: "Alex", Pass: "cool"})
	if err != nil {
		log.Println("Could not insert")
		log.Println(err)
	}


  //err := collection.Find(bson.M{"username": "raghu").One(&result)
	erre := collection.Find(nil).All(&results)
        if err != nil {
                //log.Fatal(err)
                log.Println("Not in db")
								log.Println(erre)
                return
        }
				 fmt.Println("Results All: ", results)
        //fmt.Println("Username:", result.pass)
}
