package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/tomyhero/battleship-game/html5_client/app/game"
	"github.com/tomyhero/battleship-game/utils"
	"github.com/zenazn/goji"
	"net/http"
)

var flagValue struct {
	ConfigPath string
	Port       int
}

func init() {
	flag.StringVar(&flagValue.ConfigPath, "config", "./etc/config/html5-example.toml", "set config path")
	flag.IntVar(&flagValue.Port, "port", 23456, "port")
	flag.Parse()
	flag.Set("bind", fmt.Sprintf(":%d", flagValue.Port))
}

func main() {
	config, err := utils.NewConfigFromFile(flagValue.ConfigPath)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Fail to load config")
		return
	}

	log.WithFields(log.Fields{"config": config}).Info("Loaded Config")

	webApp := setupWebApp(config)

	setupGoji(webApp)
	goji.Serve()

}

func setupWebApp(config *utils.Config) utils.WebApp {
	webApp := utils.WebApp{Config: config}
	return webApp
}

func setupGoji(webApp utils.WebApp) {
	goji.Handle("/game/*", game.NewMux(webApp))
	goji.Get("/static/*", http.FileServer(http.Dir(webApp.Config.HTML5ClientServer.AssetsPath)))
}
