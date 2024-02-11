
# APM Elastic Stack

Distributed Tracing And Logging With Elastic Stack


![Image1](image1.png)



## Usage

### Install Elastic Stack
```bash
cd elk-stack
docker compose up -d
```

### Start Backend Apps
```bash
cd backend
go mod tidy
go run app/main.go
```

### Running Load Test
```bash
k6 run load-test.js
```


## Tech Stack

**Load Testing:** [Grafana K6](https://k6.io/)

**Backend:** [Golang](https://go.dev/), [Gin](https://gin-gonic.com/)

**Observability:** [Elastic Stack](https://www.elastic.co/)




## Reference

Github : [Docker ELK](https://github.com/deviantony/docker-elk)

