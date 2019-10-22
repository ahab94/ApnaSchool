package main

import (
	"strconv"

	loads "github.com/go-openapi/loads"
	"github.com/spf13/viper"

	"github.com/ahab94/ApnaSchool/config"
	"github.com/ahab94/ApnaSchool/gen/restapi"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	api := operations.NewApnaSchoolAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
