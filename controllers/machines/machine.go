package machines

import (
	"github.com/gin-gonic/gin"
	"github.com/indigo-sadland/twok/config/holder"
	"github.com/indigo-sadland/twok/core/router"
	"github.com/indigo-sadland/twok/model/machine"
	"github.com/projectdiscovery/gologger"
)

type Machine struct {
	ID     uint32
	Name   string
	OS     string
	IP     string
	Status string
}

func Load() {
	r := router.GetRouter()
	r.GET("/machines", Get)
}

func Get(c *gin.Context) {

	db := holder.Context().DB

	res, err := machine.SelectAll(db)
	if err != nil {
		gologger.Error().Msg(err.Error())
	}

	var m []Machine
	for _, e := range res {
		m = append(m, Machine{
			ID: e.ID,
			Name: e.Name,
			OS: e.OS,
			IP: e.IP,
			Status: e.Status,
		})
	}

	c.HTML(200, "machines.html", m)
}
