package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	cassandra "cassandra-test/pkg/cassandraCRUD"
	s "cassandra-test/internal/structs"

	"github.com/gocql/gocql"
	"github.com/julienschmidt/httprouter"
)

const tableName = "test_table3"

//GetHandler prints out a single document based on QuestID and UserID
func GetHandler(session *gocql.Session) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, _ := strconv.Atoi(p.ByName("id"))

		entry := cassandra.ReadRowJson(session, id, tableName)
		jsonData, _ := json.Marshal(entry)
		w.Write(jsonData)
	}
}

//PostHandler does the upsert using Cassandra as the database
func PostHandler(session *gocql.Session) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//check if body is empty
		if r.Body == nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		//Decode JSON body to struct
		var requestStruct s.SampleTable
		err := json.NewDecoder(r.Body).Decode(&requestStruct)
		if err != nil {
			fmt.Println("Error in decoding JSON")
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		jsonData, _ := json.Marshal(requestStruct)
		cassandra.InsertRowJson(session, string(jsonData), tableName)

	}
}
