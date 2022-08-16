package api

import (
	"fmt"
	"log"
	"net/http"
	"user/models"

	"github.com/gin-gonic/gin"
)

func AddPerson(c *gin.Context) {
	person := models.Person{}
	c.BindJSON(&person)
	id := models.AddRecord(&person)
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func GetPerson(c *gin.Context) {
	person := models.Person{}
	c.BindJSON(&person)
	log.Printf("%v", &person)
	rs, _ := models.QueryRecord(&person)
	// c.JSON(http.StatusOK, gin.H{
	// 	"id":   rs.Id,
	// 	"name": rs.Name,
	// 	"desc": rs.Desc,
	// })
	c.JSON(http.StatusOK, rs)
}

func GetPersonAllPerson(c *gin.Context) {
	rs, _ := models.QueryAllRecord()
	c.JSON(http.StatusOK, gin.H{
		"data": rs,
	})
}

func UpdatePerson(c *gin.Context) {
	person := models.Person{}
	c.BindJSON(&person)
	rows := models.UpdateRecord(&person)
	msg := fmt.Sprintf("update successful effectived %d rows", rows)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func DeletePerson(c *gin.Context) {
	person := models.Person{}
	c.BindJSON(&person)
	rows := models.DeleteRecord(person.Id)
	msg := fmt.Sprintf("delete successful effectived %d rows", rows)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
