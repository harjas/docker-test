package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
// 	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

// 	fmt.Printf("%s, %s\n", accessKey, secretKey)

// 	r := mux.NewRouter()
// 	r.HandleFunc("/", Hello)
// 	http.Handle("/", r)
// 	fmt.Println("Starting up no 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // Hello is so good
// func Hello(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintln(w, "Hello World!")
// }

import (
	"fmt"
	"net/http"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func checkTableExists(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.ListTablesInput{}
	tables, _ := svc.ListTables(input)

	for _, table := range tables.TableNames {
		if *table == "music" {
			log.Printf("%s exists!\n", *table)
			fmt.Fprintf(w, "Hello astaxie!")
		}
	}
}

func main() {
	http.HandleFunc("/", checkTableExists)
	fmt.Println("Starting up no 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
