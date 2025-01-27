package controllers

import (
	"net/http"
	"practice/database"
	"practice/repository"
	"practice/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	person, err := repository.GetAllPerson(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person
	err := c.BindJSON(&person)
	if err != nil {
		panic(err)
	}

	err = repository.InsertPerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, person)
}

func UpdatePerson(c *gin.Context) {
	var person structs.Person
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&person)
	if err != nil {
		panic(err)
	}

	person.ID = id

	err = repository.UpdatePerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {
	var person structs.Person
	id, _ := strconv.Atoi(c.Param("id"))

	person.ID = id
	err := repository.DeletePerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, person)
}
