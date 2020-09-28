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
	"encoding/json"
	"fmt"
	"github.com/sonasingh46/metrics-server/model"
	"github.com/sonasingh46/metrics-server/pkg/decoder"
	"io"
	"log"
	"net/http"
)

func saveMetrics(w http.ResponseWriter, r *http.Request) {
	log.Print("Metrics save request received...")
	metrics := &model.NodeMetrics{}
	err := decoder.DecodeBody(r, metrics)
	if err != nil {
		log.Print("failed to ingest metrics:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to ingest metrics:"+err.Error())
		return
	}
	err = ControllerInstance.store.Create(*metrics)
	if err != nil {
		log.Print("failed to ingest metrics", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to ingest metrics:"+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Metrics ingested successfully")
}

func reportMetrics(w http.ResponseWriter, r *http.Request) {
	log.Print("Metrics report request received...")
	cType,err:=decoder.GetContentType(r)

	if err!=nil{
		log.Print("failed to report metrics:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "failed to report metrics:"+err.Error())
		return
	}

	if cType != "application/json" {
		err:= fmt.Errorf("Invalid content type:%s", cType)
		log.Print("failed to report metrics:", err.Error())
		w.WriteHeader(http.StatusUnsupportedMediaType)
		io.WriteString(w, "failed to report metrics:"+err.Error())
		return
	}
	mr:=ControllerInstance.store.Report()
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mr)
	if err != nil {
		log.Print("failed to report metrics", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to report metrics:"+err.Error())
		return
	}
	io.WriteString(w, "Metrics report generated successfully")
}
