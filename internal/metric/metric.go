package metric

type Metrics struct{
	p99 float64
	p90 float64
	p95 float64
	latency float64
	request_no_response int
	total_requests int

}