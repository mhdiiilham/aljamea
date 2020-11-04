package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var jamaah []*user

func main() {
	l, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now()
	r := gin.Default()
	c := cron.New(cron.WithLocation(l))

	jamaah = newJamaah()

	// Dzuhur
	c.AddFunc("CRON_TZ=Asia/Jakarta 30 12 * * 1-5", getImams)

	// Ashar
	c.AddFunc("CRON_TZ=Asia/Jakarta 0 16 * * 1-5", getImams)

	// Maghrib
	c.AddFunc("CRON_TZ=Asia/Jakarta 0 4 * * 1-5", getImams)

	c.Start()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"currentTime": now,
			"list":        jamaah,
		})
	})
	r.Run(":8080")
}

func getImams() {
	randomIndex := randomIndex(len(jamaah))
	imams := []string{}
	for i := range randomIndex {
		imams = append(imams, jamaah[i].Name)
	}
	for _, user := range jamaah {
		if user.Email != "" {
			sendMail(user.Email, imams)
		}
	}
}
