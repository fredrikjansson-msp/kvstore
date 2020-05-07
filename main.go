package main

import (
	"KVStore/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run()
}
