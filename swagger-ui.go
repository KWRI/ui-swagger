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
	_ "github.com/werpas/ui-swagger/static"
	_ "github.com/werpas/ui-swagger/static/css"
	_ "github.com/werpas/ui-swagger/static/fonts"
	_ "github.com/werpas/ui-swagger/static/images"
	_ "github.com/werpas/ui-swagger/static/js"
	_ "github.com/werpas/ui-swagger/static/lang"
	_ "github.com/werpas/ui-swagger/static/lib"
	_ "github.com/werpas/ui-swagger/static/json"
)

const (
	SWAGGER_FILE = "swagger.json"
)

func AttachSwaggerUI(router *mux.Router, base_path string) (err error) {

	// set swagger-ui routes
	staticPath, err1 := getWorkingDirectory()
	if err1 != nil {
		err = err1
	}

	// check if swagger doc exists
	if _, err2 := os.Stat("./api/" + SWAGGER_FILE); err2 == nil {

		// set swagger.json file route
		router.PathPrefix(base_path + "help/data").Handler(http.StripPrefix(base_path + "help/data", http.FileServer(http.Dir("./api"))))
	} else {
		// set default swagger doc
		router.PathPrefix(base_path + "help/data").Handler(http.StripPrefix(base_path + "help/data", http.FileServer(http.Dir(staticPath + "json"))))

		err = errors.New("swagger-ui.AttachSwaggerUI() -> ERROR: swagger.json file does not exists. " + err2.Error())
		log.Println(err.Error())
	}

	router.PathPrefix(base_path + "help/css").Handler(http.StripPrefix(base_path + "help/css", http.FileServer(http.Dir(staticPath + "css"))))
	router.PathPrefix(base_path + "help/fonts").Handler(http.StripPrefix(base_path + "help/fonts", http.FileServer(http.Dir(staticPath + "fonts"))))
	router.PathPrefix(base_path + "help/images").Handler(http.StripPrefix(base_path + "help/images", http.FileServer(http.Dir(staticPath + "images"))))
	router.PathPrefix(base_path + "help/js").Handler(http.StripPrefix(base_path + "help/js", http.FileServer(http.Dir(staticPath + "js"))))
	router.PathPrefix(base_path + "help/lang").Handler(http.StripPrefix(base_path + "help/lang", http.FileServer(http.Dir(staticPath + "lang"))))
	router.PathPrefix(base_path + "help/lib").Handler(http.StripPrefix(base_path + "help/lib", http.FileServer(http.Dir(staticPath + "lib"))))
	router.PathPrefix(base_path + "help").Handler(http.StripPrefix(base_path + "help", http.FileServer(http.Dir(staticPath))))

	return
}

func getWorkingDirectory() (staticPath string, err error) {

	// get static path from vendors first
	staticPath = "./vendor/github.com/werpas/ui-swagger/static/"
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
	staticPath = path.Dir(packagePath) + "/static/"

	return
}

