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
	"testing"

	"github.com/sonasingh46/metrics-server/model"
)

func TestHeapStore_PushCPUMetric(t *testing.T) {
	var hs *HeapStore
	hs = NewHeapStore()
	type args struct {
		m model.NodeMetrics
	}
	tests := []struct {
		name    string
		args    args
		wantTop int
	}{
		{
			name: "Test Case #1: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:            "10.2.1.1",
					PercentageCPUUsed: 23,
				},
			},
			wantTop: 23,
		},

		{
			name: "Test Case #2: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:            "10.2.1.1",
					PercentageCPUUsed: 26,
				},
			},
			wantTop: 26,
		},

		{
			name: "Test Case #3: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:            "10.2.1.1",
					PercentageCPUUsed: 10,
				},
			},
			wantTop: 26,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:            "10.2.1.1",
					PercentageCPUUsed: 40,
				},
			},
			wantTop: 40,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.2",
			args: args{
				m: model.NodeMetrics{
					NodeIP:            "10.2.1.2",
					PercentageCPUUsed: 5,
				},
			},
			wantTop: 5,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.2",
			args: args{
				m: model.NodeMetrics{
					NodeIP:            "10.2.1.2",
					PercentageCPUUsed: 17,
				},
			},
			wantTop: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs.PushCPUMetric(tt.args.m)
			gotTop := hs.GetTopCPU(tt.args.m.NodeIP)
			if gotTop != tt.wantTop {
				t.Errorf("Test case '%s' failed as "+
					"want top:%d but got top:%d", tt.name, tt.wantTop, gotTop)
			}
		})
	}
}

func TestHeapStore_PushMemoryMetric(t *testing.T) {
	var hs *HeapStore
	hs = NewHeapStore()
	type args struct {
		m model.NodeMetrics
	}
	tests := []struct {
		name    string
		args    args
		wantTop int
	}{
		{
			name: "Test Case #1: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageMemoryUsed: 23,
				},
			},
			wantTop: 23,
		},

		{
			name: "Test Case #2: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageMemoryUsed: 26,
				},
			},
			wantTop: 26,
		},

		{
			name: "Test Case #3: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageMemoryUsed: 10,
				},
			},
			wantTop: 26,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageMemoryUsed: 40,
				},
			},
			wantTop: 40,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.2",
			args: args{
				m: model.NodeMetrics{
					NodeIP:               "10.2.1.2",
					PercentageMemoryUsed: 5,
				},
			},
			wantTop: 5,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.2",
			args: args{
				m: model.NodeMetrics{
					NodeIP:               "10.2.1.2",
					PercentageMemoryUsed: 17,
				},
			},
			wantTop: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs.PushMemoryMetric(tt.args.m)
			gotTop := hs.GetTopMemory(tt.args.m.NodeIP)
			if gotTop != tt.wantTop {
				t.Errorf("Test case '%s' failed as "+
					"want top:%d but got top:%d", tt.name, tt.wantTop, gotTop)
			}
		})
	}
}

func TestHeapStore_GetTopMemory(t *testing.T) {
	var hs *HeapStore
	hs = NewHeapStore()
	type args struct {
		m model.NodeMetrics
	}
	tests := []struct {
		name            string
		args            args
		ModifyHeapStore func()
		want            int
	}{
		{
			name: "Test Case #1: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				// No-op
			},
			want: 0,
		},

		{
			name: "Test Case #2: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				metrics := model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageCPUUsed:    23,
					PercentageMemoryUsed: 45,
				}
				hs.PushMemoryMetric(metrics)
			},
			want: 45,
		},

		{
			name: "Test Case #3: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				metrics := model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageCPUUsed:    23,
					PercentageMemoryUsed: 40,
				}
				hs.PushMemoryMetric(metrics)
			},
			want: 45,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				metrics := model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageCPUUsed:    23,
					PercentageMemoryUsed: 50,
				}
				hs.PushMemoryMetric(metrics)
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ModifyHeapStore()
			if got := hs.GetTopMemory(tt.args.m.NodeIP); got != tt.want {
				t.Errorf("HeapStore.GetTopMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapStore_GetTopCPU(t *testing.T) {
	var hs *HeapStore
	hs = NewHeapStore()
	type args struct {
		m model.NodeMetrics
	}
	tests := []struct {
		name            string
		args            args
		ModifyHeapStore func()
		want            int
	}{
		{
			name: "Test Case #1: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				// No-op
			},
			want: 0,
		},

		{
			name: "Test Case #2: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				metrics := model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageCPUUsed:    23,
					PercentageMemoryUsed: 45,
				}
				hs.PushCPUMetric(metrics)
			},
			want: 23,
		},

		{
			name: "Test Case #3: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				metrics := model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageCPUUsed:    46,
					PercentageMemoryUsed: 40,
				}
				hs.PushCPUMetric(metrics)
			},
			want: 46,
		},

		{
			name: "Test Case #4: Node IP: 10.2.1.1",
			args: args{
				m: model.NodeMetrics{
					NodeIP: "10.2.1.1",
				},
			},
			ModifyHeapStore: func() {
				metrics := model.NodeMetrics{
					NodeIP:               "10.2.1.1",
					PercentageCPUUsed:    15,
					PercentageMemoryUsed: 50,
				}
				hs.PushCPUMetric(metrics)
			},
			want: 46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ModifyHeapStore()
			if got := hs.GetTopCPU(tt.args.m.NodeIP); got != tt.want {
				t.Errorf("HeapStore.GetTopMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
