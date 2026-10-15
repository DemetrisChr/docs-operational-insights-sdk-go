package main

const (
	connStr  = "couchbases://192.168.106.128"
	username = "Administrator"
	password = "password"
)

func main() {
	configuration()
	connecting()
	helloWorld()
	builtInLogger()
	logging()
	queries()
	errorHandling()
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
