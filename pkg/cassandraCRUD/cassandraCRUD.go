package cassandraCRUD

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

const tableName = "test_table3"

//CreateCQLSession opens a Session Object that connects to the cassandra cluster
func CreateCQLSession(ipAddress string, port int, keyspace string) *gocql.Session {
	// Initiate a session
	clusterconfig := gocql.NewCluster(ipAddress)
	clusterconfig.Port = port
	clusterconfig.Keyspace = keyspace
	clusterconfig.Consistency = gocql.Quorum
	clusterconfig.Timeout = time.Duration(time.Hour)
	/*
		clusterconfig.Authenticator = gocql.PasswordAuthenticator{
			Username: "",
			Password: "",
		}
	*/
	session, err := clusterconfig.CreateSession()
	if err != nil {
		log.Fatal(err.Error())
	}
	return session
}

//ReadRowJson reads one row and returns it as a json
func ReadRowJson(session *gocql.Session, id int, tableName string) []int {

	default_time := time.Date(2020, 02, 02, 02, 02, 02, 0, time.UTC)
	t := default_time.Add(time.Hour * time.Duration(id)).Truncate(time.Millisecond)

	fmt.Println(t.String(), id)

	var queryID int
	var queryIDList []int
	desc := "Hello World"
	status := 1

	iter := session.Query("SELECT id FROM "+tableName+" WHERE start_time < ? AND end_time > ? AND description = ? AND status = ? ALLOW FILTERING ;", t, t, desc, status).Iter()

	for iter.Scan(&queryID) {
		//fmt.Println(queryID)
		queryIDList = append(queryIDList, queryID)
	}
	fmt.Println(queryIDList)
	return queryIDList
}

//InsertRowJson inserts a row to the keyspace by way of json
func InsertRowJson(session *gocql.Session, jsonStruct string, tableName string) {
	err := session.Query("INSERT INTO " + tableName + " JSON ? " + jsonStruct).Exec()
	if err != nil {
		fmt.Println("Insert Error")
		fmt.Println(err)
	}
}
