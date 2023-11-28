package main

import (
	"fmt" 
	"github.com/sarracesi/user-service/user-service/service"
)

func main() {
	service.Start()
	fmt.Println("To close connection CTRL+C :-)")
}
