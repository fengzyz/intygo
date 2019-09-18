package main

import(

	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"intygo/intygo"
)

func main() {
	cfg := new (intygo.Config)
	cfg.Parse("config/app.properties")
	intygo.SetCfg(cfg)

	intygo.Configuration(cfg.Logger["filepath"])

	gin.SetMode(cfg.App["mode"])
}