package router

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/oracle", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "oracle",
		})
	})

	r.GET("/submit", func(c *gin.Context) {
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

	})

	return r
}
