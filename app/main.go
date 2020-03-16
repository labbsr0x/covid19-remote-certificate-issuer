package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// baseURL is this service base URL. Used to build output and return URLs
var baseURL string

func main() {

	logrus.Infof("Starting Certificate Issuer")
	logrus.SetLevel(logrus.DebugLevel)
	flag.StringVar(&baseURL, "baseURL", "http://localhost:8000", "Base URL of this service, used to print responses")
	flag.Parse()

	logrus.Infof("baseURL=%s", baseURL)

	logrus.Infof("Banco do Brasil COVID-19 Certificate Issuer up and running!")

	r := gin.Default()
	r.Use(ginprom.PromMiddleware(nil))
	v0 := r.Group("/v0")

	v0.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	v0.GET("/certificate", func(c *gin.Context) {
		cert := issueCertificate()
		c.JSON(200, gin.H{
			"certificate": cert,
		})
	})

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
