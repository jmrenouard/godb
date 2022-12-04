package main

//import _ shell
func main() {
	Run(`echo 'jmr !'`)

	ExecTestDB("jmr.sqlite3")
}
