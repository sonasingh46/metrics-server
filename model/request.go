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

package model

// NodeMetrics is the REST API request type contract that clients should push to this
// metrics server in JSON format.
// e.g
/*
{
	"node_ip": "10.12.2.2"
	"percentage_cpu_used": 55
	"percentage_memory_used": 25
}

{
	"node_ip":"10.10.2.1",
	"percentage_cpu_used":20,
	"percentage_memory_used":40
}
*/
type NodeMetrics struct {
	NodeIP               string `json:"node_ip"`
	PercentageCPUUsed    int    `json:"percentage_cpu_used"`
	PercentageMemoryUsed int    `json:"percentage_memory_used"`
}

// MetricsStore is the interfaces for persisting the metrics
// The persistent layers should implement the methods.
// Currently, there is a heapstore layer that implements this interface.
// Heapstore is a in-memory store for metrics data and is volatile.
type MetricsStore interface {
	Create(NodeMetrics) error
	Report() MetricsReport
}
