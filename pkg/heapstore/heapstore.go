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

package heapstore

import (
	"container/heap"
	"errors"
	"github.com/sonasingh46/metrics-server/model"
	"strings"
)

// MemoryPercentage stores memory percentage used. This is a
// max heap data structure.
type MemoryPercentage []int

func (h MemoryPercentage) Len() int           { return len(h) }
func (h MemoryPercentage) Less(i, j int) bool { return h[i] > h[j] }
func (h MemoryPercentage) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MemoryPercentage) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MemoryPercentage) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// CPUPercentage stores CPU percentage used. This is a
// max heap data structure.
type CPUPercentage []int

func (h CPUPercentage) Len() int           { return len(h) }
func (h CPUPercentage) Less(i, j int) bool { return h[i] > h[j] }
func (h CPUPercentage) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CPUPercentage) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *CPUPercentage) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


// HeapStore store is an in memory storage for metrics data.
// It is a max heap based
type HeapStore struct {
	NodeList map[string]bool
	CPUMetrics map[string]CPUPercentage
	MemoryMetrics map[string]MemoryPercentage
}

// NewHeapStore returns a new heap store to store metrics data
func NewHeapStore() *HeapStore {
	return &HeapStore{
		NodeList: map[string]bool{},
		CPUMetrics: map[string]CPUPercentage{},
		MemoryMetrics: map[string]MemoryPercentage{},
	}
}

// Create pushes the metrics in the heap store.
func (hs *HeapStore) Create(m model.NodeMetrics) error {
	if strings.TrimSpace(m.NodeIP)==""{
		return errors.New("Node ip is empty!")
	}
	hs.NodeList[m.NodeIP]=true
	hs.PushCPUMetric(m)
	hs.PushMemoryMetric(m)
	return nil
}

// Create pushes the metrics in the heap store.
func (hs *HeapStore) Report() (model.MetricsReport) {
	mr:=model.MetricsReport{
		Stats:make([]model.NodeTopMetrics,len(hs.NodeList)),
	}
	index:=0
	for nodeIP,_:=range hs.NodeList{
		mr.Stats[index].IP=nodeIP
		mr.Stats[index].MaxCPU=hs.GetTopCPU(nodeIP)
		mr.Stats[index].MaxMemory=hs.GetTopMemory(nodeIP)
		index++
	}
	return mr
}

// PushCPUMetric pushes the CPU metrics in the heap store.
func (hs *HeapStore) PushCPUMetric(m model.NodeMetrics) {
	// If there is no CPU metrics for the particular node
	// then initialize the heap store for that node and
	// push.
	if hs.CPUMetrics[m.NodeIP].Len()==0{
		cpu:=&CPUPercentage{m.PercentageCPUUsed}
		heap.Init(cpu)
		hs.CPUMetrics[m.NodeIP]=*cpu

	}else {
		// If some metrics already exists for the node
		// just push this incoming metric.
		cpu:=hs.CPUMetrics[m.NodeIP]
		heap.Push(&cpu,m.PercentageCPUUsed)
		hs.CPUMetrics[m.NodeIP]=cpu
	}
}

// PushMemoryMetric pushes the memory metrics in the heap store.
func (hs *HeapStore) PushMemoryMetric(m model.NodeMetrics) {
	// If there is no memory metrics for the particular node
	// then initialize the heap store for that node and
	// push.
	if hs.MemoryMetrics[m.NodeIP].Len()==0{
		mem:=&MemoryPercentage{m.PercentageMemoryUsed}
		heap.Init(mem)
		hs.MemoryMetrics[m.NodeIP]=*mem
	}else {
		// If some metrics already exists for the node
		// just push this incoming metric.
		mem:=hs.MemoryMetrics[m.NodeIP]
		heap.Push(&mem,m.PercentageMemoryUsed)
		hs.MemoryMetrics[m.NodeIP]=mem
	}
}

// GetTopMemory returns the top memmory usage metric for a particular node.
func (hs *HeapStore)GetTopMemory(nodeIP string)int  {
	if hs.MemoryMetrics[nodeIP].Len()==0{
		return 0
	}
	return hs.MemoryMetrics[nodeIP][0]
}

// GetTopCPU returns the top CPU usage metric for a particular node.
func (hs *HeapStore)GetTopCPU(nodeIP string)int  {
	if hs.CPUMetrics[nodeIP].Len()==0{
		return 0
	}
	return hs.CPUMetrics[nodeIP][0]
}
