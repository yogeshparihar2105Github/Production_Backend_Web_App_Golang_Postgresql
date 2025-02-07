package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux)) //using fatal bcs if the server is not running there is no point of running the program further so fatal will exit the program

}
