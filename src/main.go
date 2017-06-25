package main

import (
	"config"
	"eureka"
	"net/http"
	"strconv"
	"time"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

/**
* /service-mgr/list
* /service-mgr/read/{id}
* /service-mgr/save/{id}
* /service-mgr/del/{id}
* /service-mgr/run/{id}
* /service-mgr/stop/{id}
**/
func main() {
	log.I("Start MicroService Manager @localhost:", config.Port)

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/service-mgr/list", list)
	router.GET("/service-mgr/read/:id", read)
	router.POST("/service-mgr/save/:id", save)
	router.DELETE("/service-mgr/del/:id", delete)
	router.GET("/service-mgr/run/:id", run)
	router.GET("/service-mgr/stop/:id", stop)
	router.GET("/find-user/:name", findUser)
	router.GET("/quit", quit)
	router.GET("/query/:service", query)
	router.NotFound = http.FileServer(http.Dir("page"))

	// regist in Eureka
	e := eureka.Register(config.EurekaURL, config.AppName)
	if e != nil {
		log.E(e)
	}

	log.H("Server is started")

	go func() {
		for {
			time.Sleep(20000 * time.Millisecond)
			eureka.Renew(config.EurekaURL, config.AppName)
		}

	}()

	log.E(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}
