package db

import "github.com/jacky-htg/brilian/lib"

func IsConn() bool {
	println("from postgres")
	println(lib.Pagi())
	return true
}

func Close() {
	println("Close DB")
}
