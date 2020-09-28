#!/usr/bin/env bash

echo "Posting metrics for nodes..."

nodeIP="10.10.2.1"

echo "Posting metrics for $nodeIP"

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":24,
"percentage_memory_used":16
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":30,
"percentage_memory_used":15
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":10,
"percentage_memory_used":12
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":18,
"percentage_memory_used":45
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":20,
"percentage_memory_used":42
}'

# The max CPU and memory usage according to above data for node 10.10.2.1
# will be 30 and 45 respectively.

nodeIP="10.10.2.2"

echo "Posting metrics for $nodeIP"

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":26,
"percentage_memory_used":45
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":23,
"percentage_memory_used":34
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":21,
"percentage_memory_used":13
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":83,
"percentage_memory_used":97
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":89,
"percentage_memory_used":90
}'

# The max CPU and memory usage according to above data for node 10.10.2.2
# will be 89 and 97 respectively.

nodeIP="10.10.2.3"

echo "Posting metrics for $nodeIP"

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":25,
"percentage_memory_used":54
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":42,
"percentage_memory_used":61
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":23,
"percentage_memory_used":34
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":74,
"percentage_memory_used":78
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":21,
"percentage_memory_used":61
}'

# The max CPU and memory usage according to above data for node 10.10.2.3
# will be 74 and 78 respectively.


nodeIP="10.10.2.4"

echo "Posting metrics for $nodeIP"

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":20,
"percentage_memory_used":35
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":71,
"percentage_memory_used":64
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":25,
"percentage_memory_used":68
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":55,
"percentage_memory_used":22
}'

curl -X POST   http://localhost:8080/metrics   -H 'content-type: application/json'   -d '{
"node_ip": "'$nodeIP'",
"percentage_cpu_used":28,
"percentage_memory_used":11
}'

# The max CPU and memory usage according to above data for node 10.10.2.4
# will be 71 and 68 respectively.

echo "Getting the report..."
curl -X GET   http://localhost:8080/report   -H 'content-type: application/json'
