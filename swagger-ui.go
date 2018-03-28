package swagger_ui

import (
	"github.com/gorilla/mux"
	"net/http"
	"runtime"
	"path"
	"log"
	"errors"
	"os"

	// include static files
	//_ "github.com/KWRI/ui-swagger/static"
	//_ "github.com/KWRI/ui-swagger/static/css"
	//_ "github.com/KWRI/ui-swagger/static/fonts"
	//_ "github.com/KWRI/ui-swagger/static/images"
	//_ "github.com/KWRI/ui-swagger/static/js"
	//_ "github.com/KWRI/ui-swagger/static/lang"
	//_ "github.com/KWRI/ui-swagger/static/lib"
	//_ "github.com/KWRI/ui-swagger/static/json"
	_ "github.com/KWRI/ui-swagger/node_modules"
	_ "github.com/KWRI/ui-swagger/node_modules/swagger-ui-dist"
	_ "github.com/KWRI/ui-swagger/node_modules/next-tick"
	_ "github.com/KWRI/ui-swagger/node_modules/es6-symbol"
	_ "github.com/KWRI/ui-swagger/node_modules/es6-iterator"
	_ "github.com/KWRI/ui-swagger/node_modules/es5-ext"
	_ "github.com/KWRI/ui-swagger/node_modules/d"
)

const (
	SWAGGER_FILE = "swagger.json"
)

func AttachSwaggerUI(router *mux.Router, base_path string, swaggerBase string) (err error) {

	// set swagger-ui routes
	staticPath, err1 := getWorkingDirectory()
	if err1 != nil {
		err = err1
	}

	// check if swagger doc exists
	if _, err2 := os.Stat(swaggerBase + SWAGGER_FILE); err2 == nil {

		// set swagger.json file route
		router.PathPrefix(base_path + "help/data").Handler(http.StripPrefix(base_path + "help/data", http.FileServer(http.Dir("./api"))))
	} else {
		// set default swagger doc
		router.PathPrefix(base_path + "help/data").Handler(http.StripPrefix(base_path + "help/data", http.FileServer(http.Dir(staticPath + "json"))))

		err = errors.New("swagger-ui.AttachSwaggerUI() -> ERROR: swagger.json file does not exists. " + err2.Error())
		log.Println(err.Error())
	}

	router.PathPrefix(base_path + "help/node-modules").Handler(http.StripPrefix(base_path + "help/node-modules", http.FileServer(http.Dir(staticPath + "node-modules"))))
	router.PathPrefix(base_path + "help/swagger-ui-dist").Handler(http.StripPrefix(base_path + "help/swagger-ui-dist", http.FileServer(http.Dir(staticPath + "swagger-ui-dist"))))
	router.PathPrefix(base_path + "help/next-tick").Handler(http.StripPrefix(base_path + "help/next-tick", http.FileServer(http.Dir(staticPath + "next-tick"))))
	router.PathPrefix(base_path + "help/es6-symbol").Handler(http.StripPrefix(base_path + "help/es6-symbol", http.FileServer(http.Dir(staticPath + "es6-symbol"))))
	router.PathPrefix(base_path + "help/es6-iterator").Handler(http.StripPrefix(base_path + "help/es6-iterator", http.FileServer(http.Dir(staticPath + "es6-iterator"))))
	router.PathPrefix(base_path + "help/es5-ext").Handler(http.StripPrefix(base_path + "help/es5-ext", http.FileServer(http.Dir(staticPath + "es5-ext"))))
	router.PathPrefix(base_path + "help/d").Handler(http.StripPrefix(base_path + "help/d", http.FileServer(http.Dir(staticPath + "d"))))
	router.PathPrefix(base_path + "help").Handler(http.StripPrefix(base_path + "help", http.FileServer(http.Dir(staticPath))))

	return
}

func getWorkingDirectory() (staticPath string, err error) {

	// get static path from vendors first
	staticPath = "./vendor/github.com/KWRI/ui-swagger/node_modules/"
	if _, err1 := os.Stat(staticPath + "json/swagger.json"); err1 == nil {
		return
	}

	// get static path from calling lib otherwise
	_, packagePath, _, ok := runtime.Caller(0)
	if !ok {
		err = errors.New("swagger-ui.AttachSwaggerUI() -> ERROR: Could not get swagger-ui package path")
		log.Println(err.Error())
	}

	// set swagger-ui routes
	staticPath = path.Dir(packagePath) + "/node_modules/"

	return
}

