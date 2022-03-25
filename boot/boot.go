// Package boot carries out initial setup of the server's settings.
package boot

import (
	"github.com/indigo-sadland/twok/config"
	"github.com/indigo-sadland/twok/config/holder"
	"github.com/indigo-sadland/twok/controllers"
	"github.com/indigo-sadland/twok/core/router"

	"github.com/projectdiscovery/gologger"

)

func Initialize(conf *config.Values) {
	// Setup gin router as default.
	r := router.SetRouter()

	// Connect to MySQL DB.
	dbCon, err := conf.MySQL.Connect()
	if err != nil {
		gologger.Error().Msgf(err.Error())
	}

	// load html files
	r.Do.LoadHTMLGlob(conf.Assets.GetHTMLGlob())
	// load static
	r.Do.Static(conf.Assets.GetAssets())

	controllers.LoadRoutes()

	// Store the config connection in holder.
	holder.StoreConfig(*conf)

	// Store the DB connection in holder.
	holder.StoreDB(dbCon)

	r.Do.Run(conf.Server.GetAddress())
}
