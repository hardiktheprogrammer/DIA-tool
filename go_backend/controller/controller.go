package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type PostData struct {
	OracleAddress string `json:"oracleAddress"`
	ChainID       string `json:"chainID"`
	RPCNode       string `json:"RPCnode"`
}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func Controller() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/67e0b3b448f84921b3049e5336279397")
	if err != nil {
		log.Fatal(err)
	}

	// Define the Oracle Address and Chain ID
	oracleAddress := common.HexToAddress("0x123456789abcdef123456789abcdef12345678")
	chainID := int64(1) // Mainnet

	// Retrieve the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the most recent blocks and retrieve transactions
	for i := blockNumber; i > blockNumber-10; i-- {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}

		for _, tx := range block.Transactions() {
			if tx.To() == &oracleAddress && tx.ChainId().Cmp(big.NewInt(chainID)) == 0 {
				// Process the transaction as needed
				fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
			}
		}
	}
}

func AnalizeData(c *gin.Context) {
	// Parse the JSON data from the request body into a struct
	var postData PostData
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// Do something with the data (e.g., save it to a database)
	// ...

	// Return a success response
	c.JSON(200, gin.H{"message": "Data submitted successfully"})
}

func Createconnection() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection

}

// func CreateUser(w http.ResponseWriter, r *http.Request) {

// 	// create an empty user of type models.User
// 	var user models.User

// 	// decode the json request to user
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		log.Fatalf("Unable to decode the request body.  %v", err)
// 	}

// 	// call insert user function and pass the user
// 	insertID := insertUser(user)

// 	// format a response object
// 	res := response{
// 		ID:      insertID,
// 		Message: "User created successfully",
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(res)
// }

// func insertUser(user models.User) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the insert sql query
// 	// returning userid will return the id of the inserted user
// 	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

// 	// the inserted id will store in this id
// 	var id int64

// 	// execute the sql statement
// 	// Scan function will save the insert id in the id
// 	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	fmt.Printf("Inserted a single record %v", id)

// 	// return the inserted id
// 	return id
// }
