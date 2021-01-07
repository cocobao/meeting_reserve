package server

import (
	"fmt"
	"net/http"
	"os"

	"git.qhfct.io/comm-go/log"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(log.GinLog)
	g.Static("/static", "./static")
	g.LoadHTMLGlob("./static/views/*.html")

	g.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	return g
}

func RunService(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router())
	if err != nil {
		fmt.Println(err, "setup secure api service fail")
		os.Exit(0)
	}
}
