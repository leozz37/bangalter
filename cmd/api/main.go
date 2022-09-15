package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/bgs", GetLastBgsValue)
	r.POST("/bgs", PostBgsValue)
	r.Run()
}

type Status struct {
	Now int64 `json:"now"`
}

type Bgs struct {
	Device             string `json:"device"`
	DeviceBatteryLevel int    `json:"battery"`
	Sgv                string `json:"sgv"`
	Datetime           int    `json:"datetime"`
}

type Payload struct {
	Status    []Status `json:"status"`
	BgsValues []Bgs    `json:"bgs"`
}

var LastBgsValue = Bgs{}

func GetLastBgsValue(c *gin.Context) {
	c.JSON(http.StatusOK, LastBgsValue)
}

func PostBgsValue(c *gin.Context) {
	var payload Payload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := parsePayload(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(payload)
}

func parsePayload(payload Payload) error {
	if LastBgsValue.Datetime != payload.BgsValues[0].Datetime {
		// TODO save to InfluxDB
		return nil
	}

	LastBgsValue = payload.BgsValues[0]

	return nil
}
