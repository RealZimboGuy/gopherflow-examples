package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/RealZimboGuy/gopherflow/pkg/gopherflow"
	"github.com/RealZimboGuy/gopherflow/pkg/gopherflow/core"
	"github.com/RealZimboGuy/gopherflow_example/controllers"
	"github.com/RealZimboGuy/gopherflow_example/internal/workflows"
	_ "github.com/mkevac/debugcharts"
)

func main() {

	//use your own database type here, this is just to make the example easy to start and run
	os.Setenv("GFLOW_DATABASE_TYPE", "SQLLITE")

	ctx := context.Background()
	//you may do your own logger setup here or use this default one with slog
	gopherflow.SetupLogger(slog.LevelInfo)

	workflowRegistry := map[string]func() core.Workflow{
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

	app := gopherflow.Setup(workflowRegistry)

	if err := app.Run(ctx); err != nil {
		slog.Error("Engine exited with error", "error", err)
	}

}
