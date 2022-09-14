package main

import (
	"fmt"

	"github.com/NaufalHSyahputra/alterra-agmc/config"
	"github.com/NaufalHSyahputra/alterra-agmc/routes"
)

var (
	cfg config.Config
)

func init() {
	cfg = config.InitConfig()
}

func main() {
	config.InitDB(cfg)
	e := routes.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.AppPort)))
}
