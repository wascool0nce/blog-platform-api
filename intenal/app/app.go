package app

import (
	"net/http"

	"github.com/wascool0nce/blog-platform-api/config"
	"github.com/wascool0nce/blog-platform-api/intenal/handler"
	v1 "github.com/wascool0nce/blog-platform-api/intenal/handler/v1"
	sqlite "github.com/wascool0nce/blog-platform-api/intenal/infra"
	"github.com/wascool0nce/blog-platform-api/intenal/repository"
	"github.com/wascool0nce/blog-platform-api/intenal/service"
)

func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	db, err := sqlite.ConnectDB(*cfg)
	if err != nil {
		return err
	}

	defer db.Close()

	postRepo := repository.NewPostRepository(db)

	postService := service.NewPostService(postRepo)

	postHandler := v1.NewPostHandler(postService)
	router := handler.NewRouter(postHandler)
	return http.ListenAndServe(cfg.HTTP.Port, router)
}
