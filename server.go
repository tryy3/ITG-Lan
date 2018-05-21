package main

import (
	"github.com/tryy3/itg-lan/api"
	"github.com/tryy3/itg-lan/models"
	"github.com/tryy3/itg-lan/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewMysqlDB("root", "", "itg_lan") // Update this
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
