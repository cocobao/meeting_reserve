package main

import (
	"os"

	"github.com/cocobao/comm-go/log"
	"github.com/cocobao/meeting_reserve/config"
	"github.com/cocobao/meeting_reserve/server"
	"github.com/cocobao/meeting_reserve/store"
)

func setupLog() {
	logpath := config.GetConfig().Log.LogPath
	if len(logpath) > 0 {
		os.MkdirAll(logpath, os.ModePerm)
	}

	log.NewLogger(logpath, config.GetConfig().Log.LogLevel)
}

func setupDB() {
	c := config.GetConfig().MongoDB
	store.SetupMgo(c.Host, c.UserName, c.Password)
}

func main() {
	setupLog()
	setupDB()
	server.RunService(config.GetConfig().Port)
}
