// Package boot carries out initial setup of the server's settings.
package boot

import (
	"github.com/indigo-sadland/twok/config"
	"github.com/indigo-sadland/twok/config/holder"
	"github.com/indigo-sadland/twok/controllers"
	"github.com/indigo-sadland/twok/core/router"
)

func Initialize(conf *config.Values) {
	// Setup gin router as default.
	r := router.SetRouter()

	// Connect to MySQL DB.
	dbCon, _ := conf.MySQL.Connect()

	// load html files
	r.Do.LoadHTMLGlob(conf.Assets.GetHTMLGlob())
	// load static
	r.Do.Static(conf.Assets.GetAssets())

	controllers.Load()

	// Store the DB connection in holder.
	holder.StoreDB(dbCon)

	r.Do.Run(conf.Server.GetAddress())

}
