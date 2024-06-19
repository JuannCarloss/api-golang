package main

import (
	"example/payments/routers"
)

func main() {

	r := routers.SetupRouters()

	r.Run()
}
