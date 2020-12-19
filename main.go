package main

import "golang-postgresql/routes"

func main() {

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
