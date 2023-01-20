package main

import (
	"bikefit/databases"
	"bikefit/routers"
)

func main() {
	databases.ConectaBanco()
	routers.IniciaRoteamento()
}
