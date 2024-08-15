package main

import (
	db "example.com/rest-api/database"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080" // localhost:8080

func main(){
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(PORT) 
}
