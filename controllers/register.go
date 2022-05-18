package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"myapp/models"
)

//////Add User//////
func UserAdd(user models.User) {

	addUser := models.User{ID: user.ID, Name: user.Name, Location: user.Location, Gender: user.Gender, Email: user.Email}

	if err := models.MPosGORM.Create(&addUser).Error; err != nil {
		fmt.Println("error add User")
		return
	}
	fmt.Printf("%s User added/n", user.Name)
}

//////Add Likes//////
func LikeAdd(like models.Like) {

	addLike := models.Like{Id: like.Id, Liker: like.Liker, Likee: like.Likee}

	if err := models.MPosGORM.Create(&addLike).Error; err != nil {
		fmt.Println("error add Like")
		return
	}
	fmt.Printf("%d Like added/n", like.Id)
}

//////Delete User//////
func UserDelete(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	var user models.User
	c.BindJSON(&user)

	if err := models.MPosGORM.Where("user_id = ?", user.ID).Delete(&models.User{}).Error; err != nil {
		fmt.Println("error delete User")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.Name,
		"user_id":user.ID,
	})
}

//////Get All Matches//////
func GetMatchesForAll(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	var likes []models.Like

	if err := models.MPosGORM.Raw("SELECT A.* from likes A, likes B where A.liker = B.likee and A.likee = B.liker").Scan(&likes).Error; err != nil {
		fmt.Println("error Find Matches")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if(likes != nil){
		c.JSON(http.StatusOK, likes)
	} else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}
}

//////Get Match for One user//////
func GetMatchesForOne(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	user_id := c.Param("userid")

	var users []models.User

	if err := models.MPosGORM.Raw("SELECT U.* from users U, likes A, likes B where A.liker = B.likee and A.likee = B.liker and A.likee = ? and U.id = B.likee", user_id).Scan(&users).Error; err != nil {
		fmt.Println("error Find Matches")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if(users != nil){
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}
}

//////Get Users within distance k from User X//////
func GetUsersAtDistance(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	user_id := c.Param("userid")
	k := c.Param("k")

	var users []models.User

	if err := models.MPosGORM.Raw("SELECT B.* from users A, users B where A.id = ? and ((B.location <= (A.location+?)) and (B.location >= (A.location-?))) and B.id != ? ", user_id, k, k, user_id).Scan(&users).Error; err != nil {
		fmt.Println("error Find Matches")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if(users != nil){
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}
}

//////Query By Name letters//////
func GetUsersNameQuery(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	name_str := "%"+c.Query("name")+"%"

	var users []models.User

	if err := models.MPosGORM.Raw("SELECT A.* from users A where A.name like ?", name_str).Scan(&users).Error; err != nil {
		fmt.Println("error Find Matches")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if(users != nil){
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}
}