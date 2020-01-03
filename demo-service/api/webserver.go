package api

import (
	"log"
	"net/http"
)

func StartWebServer(port string)  {
	router := NewRoute()
	http.Handle("/",router)
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Println("An error occurred starting HTTP listener at port" + port)
		log.Println("Error:" + err.Error())
	}
}
