package machines

import (
	"encoding/json"
	"fmt"

	"github.com/indigo-sadland/twok/config/holder"
	"github.com/indigo-sadland/twok/core/router"
	"github.com/indigo-sadland/twok/model/machinesTable"
	"github.com/projectdiscovery/gologger"

	"github.com/gin-gonic/gin"
	"github.com/terra-farm/go-virtualbox"
)

// Machine holds structure of data from DB.
type Machine struct {
	ID     uint32
	Name   string
	OS     string
	IP     string
	Status string
}

// TableData holds structure of data from POST request.
type TableData struct {
	Name   string `json:"name"`
	Status string `json:"currentStatus"`
}

func Load() {
	r := router.GetRouter()
	r.GET("/machines", Get)
	r.POST("/machines", ManageState)
}

func Get(c *gin.Context) {
	db := holder.Context().DB
	res, err := machinesTable.SelectAll(db)
	if err != nil {
		gologger.Error().Msg(err.Error())
		c.Status(500)
		c.Writer.WriteString(err.Error())
	}

	var m []Machine
	for _, e := range res {
		m = append(m, Machine{
			ID:     e.ID,
			Name:   e.Name,
			OS:     e.OS,
			IP:     e.IP,
			Status: e.Status,
		})
	}

	c.HTML(200, "machines.html", m)
}

//TODO: ADD NOTIFICATION WINDOW IN FRONTEND TO SHOW GENERAL ERRORS.
//TODO: CHANGE LOGGER! WRITE LOGS TO A FILE.

// ManageState performs a machine's state changing.
func ManageState(c *gin.Context) {
	var td TableData

	// Read body of POST request.
	r := c.Request.Body
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&td)
	if err != nil {
		gologger.Error().Msg(err.Error())
	}
	if string(td.Status) == "STOPPED" {
		if statusCheck() {
			runMachine(td.Name)
		}

	} else if string(td.Status) == "RUNNING" {
		stopMachine(td.Name)

	} else {
		err := "Unknown machine state"
		gologger.Error().Msg(err)
		_, err2 := c.Writer.WriteString(err)
		if err2 != nil {
			return
		}
	}
}

func statusCheck() bool {
	db := holder.Context().DB

	_, noRows, err := machinesTable.ByStatus(db, machinesTable.StatusRunning)
	if noRows {
		return true
	} else {
		msg := fmt.Sprintf("Can't run the machine because there is already running one. \n")
		gologger.Warning().Msg(msg)
	}

	if err.Error() != "sql: no rows in result set" {
		gologger.Error().Msg(err.Error())
	}

	return false
}

func runMachine(name string) {
	instance := virtualbox.Machine{
		Name:  name,
		State: virtualbox.Poweroff,
	}
	err := instance.Start()
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}

	db := holder.Context().DB
	_, err = machinesTable.UpdateStatus(db, name, machinesTable.StatusRunning)
	if err != nil {
		gologger.Error().Msg(err.Error())
	}
}

func stopMachine(name string) {
	instance := virtualbox.Machine{
		Name:  name,
		State: virtualbox.Running,
	}
	err := instance.Stop()
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}

	db := holder.Context().DB
	_, err = machinesTable.UpdateStatus(db, name, machinesTable.StatusStopped)
	if err != nil {
		gologger.Error().Msg(err.Error())
	}
}
