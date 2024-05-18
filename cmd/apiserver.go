package main

import (
	"davinci/cmd/app"
	"davinci/config"
	_ "davinci/config"
	"fmt"
)

func main() {
	cmd := app.NewApiServerCommand()
	fmt.Println("config.SysYamlconfig.Server.Name = ", config.SysYamlconfig.Server.Name)
	cmd.Execute()
}

// go run cmd/apiserver.go apiserver --port=8888
