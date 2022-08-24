package main

import (
	"docker_go/api/config"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	_ = config.NewMysqlConfig()
	router := httprouter.New()
	router.GET("/", Index)
	http.ListenAndServe(":8081", router)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
