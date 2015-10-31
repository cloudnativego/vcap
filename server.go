package main

import (
	"net/http"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/cf-tools"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

//NewServer configures and returns a Server.
func NewServer(appEnv *cfenv.App) *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter, appEnv)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, appEnv *cfenv.App) {
	mx.HandleFunc("/vcap", vcapHandler(formatter, appEnv)).Methods("GET")
}

func vcapHandler(formatter *render.Render, appEnv *cfenv.App) http.HandlerFunc {
	val, err := cftools.GetVCAPServiceProperty("dispenser-task-service", "uri", appEnv)

	return func(w http.ResponseWriter, req *http.Request) {
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, struct{ Error string }{err.Error()})
			return
		}
		formatter.JSON(w, http.StatusOK, struct{ Foo string }{val})
	}
}
