package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/soulski/dmp-cli"
	"github.com/soulski/dmp-scenario/authen/server"
)

const (
	NAMESPACE     = "authen"
	CONTACT_POINT = "http://127.0.0.1:%d/dmp"
)

var urlActions = map[string]func(w http.ResponseWriter, r *http.Request, c *authen.Context){
	"PUT:/dmp": authen.DMP,
}

var fAssetPath string
var fTemplatePath string
var fPort int

func main() {
	flag.StringVar(&fAssetPath, "a", "./assets/", "Asset path")
	flag.StringVar(&fTemplatePath, "t", "./templates/", "Template path")
	flag.IntVar(&fPort, "p", 8081, "Bind port")
	flag.Parse()

	db, err := authen.NewDB("/usr/authen/authen.db")
	if err != nil {
		panic("Cannot open database : " + err.Error())
	}

	db.CreateTable()

	defer db.Close()

	ctx := authen.NewContext(db)
	router := mux.NewRouter()

	for pattern, action := range urlActions {
		token := strings.Split(pattern, ":")
		router.HandleFunc(token[1], func(w http.ResponseWriter, r *http.Request) {
			action(w, r, ctx)
		}).Methods(token[0])
	}

	router.PathPrefix("/assets/").Handler(
		http.StripPrefix("/assets/", http.FileServer(http.Dir(fAssetPath))),
	)

	n := negroni.New()

	n.UseHandler(router)

	go func() {
		_, err := dmpc.RegisterService(&dmpc.Service{
			Namespace:    NAMESPACE,
			ContactPoint: fmt.Sprintf(CONTACT_POINT, fPort),
		})

		if err != nil {
			panic(err)
		}
	}()

	n.Run(":" + strconv.Itoa(fPort))
}
