package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"cassandra-test/internal/handlers"
	"cassandra-test/pkg/cassandraCRUD"

	"github.com/julienschmidt/httprouter"
)

func main() {

	//set up cassandra cluster ip addresses, port, and keyspace (equiv. to MySQL DB)
	var ipAddress = fmt.Sprintf("%s,%s,%s", os.Getenv("NODE_ONE"), os.Getenv("NODE_TWO"), os.Getenv("NODE_THREE"))
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	var keyspace = os.Getenv("KEYSPACE")

	// open a new Cassandra Session
	session := cassandraCRUD.CreateCQLSession(ipAddress, port, keyspace)

	// make new router
	router := httprouter.New()
	router.GET("/:id", handlers.GetHandler(session))
	//router.POST("/", handlers.PostHandler(session))

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Running.....")
}
