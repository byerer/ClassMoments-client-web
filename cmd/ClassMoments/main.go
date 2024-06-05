package main

import "ClassMoments-client-web/cmd"

func main() {
	app, err := cmd.InitApplication()
	if err != nil {
		panic(err)
	}
	app.Run()
}
