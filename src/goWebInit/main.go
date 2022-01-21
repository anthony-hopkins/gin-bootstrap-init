package main

import (
	"fmt"
	"log"
	"os"
)

// dirs contains all the directories we need to build for a generic web application. The root directories of nested
// directories are created automatically so there's no need to define then specifically.
var dirs = map[string]string{
	"assets:":     "./public/assets",
	"imgs":        "./public/imgs",
	"css":         "./public/css",
	"js":          "./public/js",
	"defines":     "./templates/defines",
	"views":       "./templates/views",
	"models":      "./models",
	"controllers": "./controllers/helpers",
	"src":         "./src",
	"appName":     fmt.Sprintf("./src/%s", parseAppName()),
}

// parseAppName checks to ensure an app name was passed as an argument and terminates otherwise.
func parseAppName() string {
	if len(os.Args) != 2 {
		log.Fatal("ERR: Please specify a SINGLE argument that will be used as your application name.")
	}
	appName := os.Args[1]
	return appName
}

// buildAppDirs handles actually building the web app directories.
func buildAppDirs() {
	for dirName, relativePath := range dirs {
		err := os.MkdirAll(relativePath, os.ModePerm)
		if err != nil {
			fmt.Println(relativePath)
			log.Fatal(err)
		}
		fmt.Printf("%s created\n", dirName)
	}
}

func main() {
	buildAppDirs()
}
