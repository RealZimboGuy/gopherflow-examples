package main

import (
	"log/slog"
	"net/http"
	"reflect"

	"github.com/RealZimboGuy/gopherflow/pkg/gopherflow"
	"github.com/RealZimboGuy/gopherflow_example/controllers"
	"github.com/RealZimboGuy/gopherflow_example/internal/workflows"
)

func main() {

	//you may do your own logger setup here or use this default one with slog
	gopherflow.SetupLogger()

	gopherflow.WorkflowRegistry = map[string]reflect.Type{
		"DemoWorkflow":  reflect.TypeOf(workflows.DemoWorkflow{}),
		"GetIpWorkflow": reflect.TypeOf(workflows.GetIpWorkflow{}),
	}
	mux := http.NewServeMux()

	demoController := controllers.NewDemoController()
	demoController.RegisterRoutes(mux)

	app := gopherflow.Setup(mux)

	if err := app.Run(); err != nil {
		slog.Error("Engine exited with error", "error", err)
	}

}
