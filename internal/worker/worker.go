package Worker

import (
	config "ArnavSaroj/LoadMonkey.git/internal/config"
	"ArnavSaroj/LoadMonkey.git/internal/metric"
	"sync"
)

//custom enum type to define worker_status and later for region too
//golang doesn't hv enum

type worker_status int
type region int 

const (
	idle worker_status =0
working worker_status=1
failed worker_status=2

)

const (
	//more regions to be added later only india and usa europr for now
	IND region =1
	USA region =2
	EUROPE region =3
)



type Worker struct {
	id    int 
	mutex sync.Mutex
	worker_status worker_status
region_id region
Metrics metric.Metrics
Config config.Config

}