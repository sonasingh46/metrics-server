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

// MetricsReport is the REST API response type for '/report' endpoint
type MetricsReport struct {
	Stats []NodeTopMetrics `json:"stats"`
}

type NodeTopMetrics struct {
	IP string `json:"ip"`
	MaxCPU int `json:"max_cpu"`
	MaxMemory int `json:"max_memory"`
}