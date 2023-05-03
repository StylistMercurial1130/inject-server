package main

import (
	"inject-server/credentials"
)

func init() {
    credentials.Load()
}
func main() {
    server := Initserver(":8080")
    server.Runserver()   
}
