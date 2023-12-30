package boot

import (
	"twok/config"
	"twok/controllers"
	"twok/core/router"
)

func Initialize(conf *config.Values) {
	r := router.SetRouter()

	// load html files
	r.Do.LoadHTMLGlob(conf.Assets.GetHTMLGlob())
	// load static
	r.Do.Static(conf.Assets.GetAssets())

	controllers.Load()

	r.Do.Run(conf.Server.GetAddress())

}
