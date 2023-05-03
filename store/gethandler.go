package store

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db *Database) Getfromnamehandler(ctx *gin.Context) {
    name := ctx.Query("name")
    res , err := db.Retrievebyname(name)
    if err != nil {
        ctx.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
    } else {
        ctx.JSON(http.StatusOK,gin.H{"result":res})
    } 
}
