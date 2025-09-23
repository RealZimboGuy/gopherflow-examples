package controllers

import (
	"net/http"
)

type DemoController struct {
}

func (c *DemoController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/demo", c.handleGetDemo)
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

func (c *DemoController) handleGetDemo(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World Demo"}`))
	return

}
