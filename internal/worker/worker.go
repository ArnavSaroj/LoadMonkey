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
	Id    int 
	Mutex sync.Mutex
	Worker_status worker_status
Region_id region
Metrics metric.Metrics
Config config.Config

}

func  NewWorker(id int,region_id region,config config.Config) (*Worker){

return &Worker{
Id:id,
Mutex: sync.Mutex{},
Worker_status: 0,
Metrics:*metric.New_metric() ,
Config: config,
}

}





