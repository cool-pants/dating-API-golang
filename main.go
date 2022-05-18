package main

import (
	"encoding/json"
	"io/ioutil"
	"myapp/controllers"
	"myapp/models"

	"github.com/gin-gonic/gin"
)

func main(){
	models.InitGormPostgres()
	defer models.MPosGORM.Close()
	var err error
	readUser("users.json", err)
	readLike("likes.json", err)


	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/likes/matches", controllers.GetMatchesForAll)
	api.GET("/likes/matches/id/:userid", controllers.GetMatchesForOne)
	api.POST("/likes/add", controllers.AddLike)
	
	api.GET("/users/", controllers.GetUsersNameQuery)
	api.GET("/users/id/:userid/k/:k", controllers.GetUsersAtDistance)
	api.POST("/users/add", controllers.AddUser)
	api.PUT("/users/update", controllers.UpdateUser)
	api.DELETE("/users/delete", controllers.UserDelete)

	router.Run(":4800")
}

func readUser(fileName string, err error){
	var user []models.User

	file, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal([]byte(file), &user)
	for i := 0; i<len(user);i++{
		if err = controllers.UserAdd(user[i]);err!=nil{
			return
		}
	}
}
func readLike(fileName string, err error){
	var likes []models.Like	
	file, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal([]byte(file), &likes)
	for i := 0; i<len(likes);i++{
		if err = controllers.LikeAdd(likes[i]);err!=nil{
			return
		}
	}
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}