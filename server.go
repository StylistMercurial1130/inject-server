package main

import (
	"inject-server/store"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
    port string
    db *store.Database  
    gin_engine *gin.Engine
}

func Initserver(port string) Server {
    var server Server
    var err error
    server.port = port
    if server.db , err = store.NewDB(); err != nil {
        log.Fatal(err)
    }
    server.gin_engine = gin.Default(); 
    server.gin_engine.GET("/query",server.db.Getfromnamehandler)
    return server
}

func (server Server) Runserver() {
    log.Fatal(server.gin_engine.Run())
}






