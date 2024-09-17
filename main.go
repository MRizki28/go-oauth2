package main

import "github.com/MRizki28/go-oauth2/src/routes"

func main() {
	r := routes.Route()
	r.Run(":9000")
}
