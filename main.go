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

	//readUser("users.json")
	//readLike("likes.json")


	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/likes/matches", controllers.GetMatchesForAll)
	api.GET("/likes/matches/id/:userid", controllers.GetMatchesForOne)
	api.GET("/users/id/:userid/k/:k", controllers.GetUsersAtDistance)
	api.GET("/users/", controllers.GetUsersNameQuery)

	router.Run(":4800")
}

func readUser(fileName string){
	var user []models.User

	file, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal([]byte(file), &user)
	for i := 0; i<len(user);i++{
		controllers.UserAdd(user[i])
	}
}
func readLike(fileName string){
	var likes []models.Like	
	file, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal([]byte(file), &likes)
	for i := 0; i<len(likes);i++{
		controllers.LikeAdd(likes[i])
	}
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}