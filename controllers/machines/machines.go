package machines

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/indigo-sadland/twok/config/holder"
	"github.com/indigo-sadland/twok/core/router"
	"github.com/indigo-sadland/twok/model/machines"
)

var tmpl *template.Template

type Machine struct {
	ID     uint32
	Name   string
	OS     string
	Status string
	Class  string
}

func Load() {
	r := router.GetRouter()
	r.GET("/machines", Get)
}

func Get(c *gin.Context) {
	//r := router.GetRouter()
	db := holder.Context().DB

	res, err := machines.SelectAll(db)
	if err != nil {
		fmt.Println(err)
	}

	var machines []Machine
	for _, m := range res {
		machines = append(machines, Machine{
			ID: m.ID,
			Name: m.Name,
			OS: m.OS,
			Status: m.Status,
			Class: "status danger",
		})
	}

	//tmpl, err = tmpl.ParseGlob(holder.Context().Config.Assets.HTMLGlob)

	//r.SetHTMLTemplate(tmpl)
	c.HTML(200, "machines.html", machines)
}
