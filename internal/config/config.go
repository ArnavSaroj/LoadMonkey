package config


type http_method int

const (
	get    http_method = 1
	post   http_method = 2
	delete http_method = 3
	put    http_method = 4
)

type Config struct {
	url         string
	http_method http_method
	//duration in seconds
	duration int
	// currently only taking static , will write enum for it later currenlty only takes 1 for static test as default
	ramp_config int
	timeout     int
	//regarding body need to see how well it can go , i have not decided atp of writing code if we need to convert body to json or not 
	body        string
}