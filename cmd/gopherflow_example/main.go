package main

import (
	"log/slog"

	"github.com/RealZimboGuy/gopherflow/pkg/gopherflow"
	"github.com/RealZimboGuy/gopherflow/pkg/gopherflow/core"
	"github.com/RealZimboGuy/gopherflow_example/controllers"
	"github.com/RealZimboGuy/gopherflow_example/internal/workflows"
	_ "github.com/mkevac/debugcharts"
)

func main() {

	//you may do your own logger setup here or use this default one with slog
	gopherflow.SetupLogger()

	gopherflow.WorkflowRegistry = map[string]func() core.Workflow{
		"DemoWorkflow": func() core.Workflow {
			return &workflows.DemoWorkflow{}
		},
		"GetIpWorkflow": func() core.Workflow {
			// You can inject dependencies here
			return &workflows.GetIpWorkflow{
				// HTTPClient: httpClient,
				// MyService: myService,
			}
		},
	}

	demoController := controllers.NewDemoController()
	demoController.RegisterRoutes()

	app := gopherflow.Setup()

	if err := app.Run(); err != nil {
		slog.Error("Engine exited with error", "error", err)
	}

}
