package dtos

type HealthProbeResponse struct {
	State   bool        `json:"state"`
	Message interface{} `json:"message,omitempty"`
}

type LiveProbeResponse struct {
	AppName   string `json:"app_name"`
	Env       string `json:"environment"`
	AppStatus string `json:"app_status"`
	UpTime    string `json:"up_time"`
}
