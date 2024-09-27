package config

/*
	The purpose of this class is to consolidate all the configuration needed to bootstrap the
	golang application.
*/

//TODO Put in more detail

import (
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"

	"sdk_workbench_authentication/src/clients"
	"sdk_workbench_authentication/src/config/env"
	"sdk_workbench_authentication/src/config/logger"
	"sdk_workbench_authentication/src/constants"
)

type Application struct {
	Props *env.Properties // application properties
}

func Bootstrap() Application {
	app := &Application{}
	// read properties
	app.Props = env.NewProperties()
	// set logging
	logger.InitLogger(app.Props.AppEnv, app.Props.LoggingLevel)
	// initialize gin
	SetGinMode(app.Props.GinMode)
	// print banner if dev mode
	if app.Props.AppEnv == "DEV" {
		printBanner(app.Props)
	}
	logger.Debug().Msg("[🛠️] Setting up ArangoDB ...")
	if app.Props.ArangoDbCreateMode != string(constants.ARANGO_DB_CREATION_MODE_NONE) {
		client, _ := clients.ArangoDBConnect(env.GetProperties().ArangoAddr, env.GetProperties().ArangoUser, env.GetProperties().ArangoPwd)
		clients.CreateOrUpdateDB(client, app.Props.ArangoDbName, app.Props.ArangoCollectionName)
		clients.CreateOrUpdateDB(client, app.Props.ArangoDbName, app.Props.ArangoHistoryCollectionName)
		clients.CreateOrUpdateDB(client, app.Props.ArangoDbName, "entitlement")
	}

	logger.Debug().Msg("[🛠️] Caching and Pooling processes...")
	// load all the process in-memory
	//processDeploymentService, err := services.NewProcessDeploymentService()
	/*if err != nil {
		logger.Error().Msg(err.Error())
	}
	processModels, _, err := processDeploymentService.GetActiveProcesses()
	if err != nil {
		logger.Error().Msg("Error getting active processes")
	}
	processDefinitionStore := bpmnEngineManager.CacheProcessesDefinitions(processModels)
	bpmnEngineManager.CreateNewProcessEngine()

	// create instance pools for active processes
	for _, processDeploymentModel := range processDefinitionStore.ProcessStore {
		logger.Debug().Msg("Starting pool creation..")
		if processDeploymentModel.InstancePool > 0 {
			go bpmnEngineManager.CreateProcessInstancePool(processDeploymentModel.Key, processDeploymentModel.InstancePool)
		}
	}
	*/
	return *app
}

func printBanner(prop *env.Properties) {
	templ := `{{ .Title "` + prop.AppName + `" "" 7 }}
	{{ .AnsiColor.BrightCyan }}` + prop.AppDesc + `{{ .AnsiColor.Default }}

	GoVersion: {{ .GoVersion }}
	GOOS: {{ .GOOS }}
	GOARCH: {{ .GOARCH }}
	NumCPU: {{ .NumCPU }}
	GOPATH: {{ .GOPATH }}
	GOROOT: {{ .GOROOT }}
	Compiler: {{ .Compiler }}
	ENV: {{ .Env "GOPATH" }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}

	{{ .AnsiColor.BrightGreen }}` + prop.AppCopyright + `{{ .AnsiColor.BrightMagenta }}
	===================================================================================
	
`

	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
}
