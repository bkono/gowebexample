package main

import (
	"github.com/micro/go-log"

	"github.com/bkono/gowebexample/web/handler"
	"github.com/micro/go-web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.image"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// This route will be hit if using `micro api --handler http --namespace go.micro.web`
	service.HandleFunc("/image/upload", handler.ImageUpload)

	// This route will be hit if using `micro web` instead of api
	service.HandleFunc("/upload", handler.ImageUpload)

	// This route is everything not caught above
	service.HandleFunc("/", handler.IndexHandler)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
