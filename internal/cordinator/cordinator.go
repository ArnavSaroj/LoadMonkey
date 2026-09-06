package Cordinator

import (
	"ArnavSaroj/LoadMonkey.git/internal/config"
	"ArnavSaroj/LoadMonkey.git/internal/metric"
	Worker "ArnavSaroj/LoadMonkey.git/internal/worker"
	"sync"
)

type Cordinator struct {
	workers []*Worker.Worker
	mutex sync.Mutex
	config config.Config
	metrics metric.Metrics
	//hearbeat will also come here will do that later


}