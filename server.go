package main

import (
        "net/http"
        "github.com/gin-gonic/gin")

type USER struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
}

var users = []USER {
    {ID: "1", Name: "Marco", Age: 21},
    {ID: "2", Name: "Anthony", Age: 20},
    {ID: "3", Name: "Aidan", Age: 20},
    {ID: "4", Name: "John", Age: 20},
    {ID: "5", Name: "Keator", Age: 75},
}

func get_users(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func post_add_users(c *gin.Context){
    var new_user USER
    if err := c.BindJSON(&new_user); err != nil {
        return
    }

    users = append(users, new_user)
    c.IndentedJSON(http.StatusCreated, new_user)

}

func main() {
    router := gin.Default()
    router.GET("/users", get_users)
    router.POST("/users", post_add_users)

    router.Run("localhost:8085")
    
}