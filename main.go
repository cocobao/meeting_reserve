package main

import (
	"fmt"
	"os"

	"git.qhfct.io/comm-go/log"
	"git.qhfct.io/comm-go/sig"
	"git.qhfct.io/comm-go/utility"
	"git.qhfct.io/meeting_reserve/config"
	"git.qhfct.io/meeting_reserve/server"
	"git.qhfct.io/meeting_reserve/store"
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
	go sig.GracefullyStopSever(func() {
		fmt.Println("~stop~~", utility.NowN())
	})
	setupLog()
	setupDB()
	server.RunService(config.GetConfig().Port)
}
