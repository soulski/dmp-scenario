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
	"github.com/soulski/dmp-scenario/search/server"
)

const (
	NAMESPACE     = "search"
	CONTACT_POINT = "http://127.0.0.1:%d/dmp"
)

var urlActions = map[string]func(w http.ResponseWriter, r *http.Request, c *search.Context){
	"PUT:/dmp": search.DMP,
}

var fAssetPath string
var fTemplatePath string
var fPort int

func main() {
	flag.IntVar(&fPort, "p", 8081, "Bind port")
	flag.Parse()

	db, err := search.NewDB("/usr/search/search.db")
	if err != nil {
		panic("Cannot open database : " + err.Error())
	}

	db.CreateTable()

	defer db.Close()

	ctx := search.NewContext(db)
	router := mux.NewRouter()

	for pattern, action := range urlActions {
		token := strings.Split(pattern, ":")
		router.HandleFunc(token[1], func(w http.ResponseWriter, r *http.Request) {
			action(w, r, ctx)
		}).Methods(token[0])
	}

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

		err = dmpc.SubscribeTopic("create-user")
		if err != nil {
			panic(err)
		}

		err = dmpc.SubscribeTopic("create-item")
		if err != nil {
			panic(err)
		}
	}()

	n.Run(":" + strconv.Itoa(fPort))
}
