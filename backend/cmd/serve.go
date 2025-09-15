package cmd

import (
	"ecommerce/config"
	"ecommerce/database/product_query"
	"ecommerce/rest"
)

func Serve() {
	cnf := config.GetConfig()

	product_query.GetConnection(cnf)
	defer product_query.CloseConnection()

	rest.Start(cnf)

}
