/*
Copyright 2020, Author: Ashutosh Kumar (GithubID: @sonasingh46).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"github.com/gorilla/mux"
	"github.com/sonasingh46/metrics-server/model"
	"github.com/sonasingh46/metrics-server/pkg/heapstore"
	"log"
	"net/http"
)

type Controller struct {
	store model.MetricsStore
}

func NewController() *Controller {
	return &Controller{
		store: heapstore.NewHeapStore(),
	}
}

var ControllerInstance *Controller

func StartMetricsServer() {
	r := mux.NewRouter()
	// Endpoint for checking service liveness
	r.HandleFunc("/healthz", HealthCheckHandler)

	// Endpoint for posting metrics
	r.HandleFunc("/metrics", saveMetrics).Methods("POST")

	// Endpoint for metrics report
	// This endpoint returns top memory and CPU usage till time
	// for the particular node ip.
	r.HandleFunc("/report", reportMetrics).Methods("GET")
	ControllerInstance = NewController()
	log.Print("Metrics server started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
