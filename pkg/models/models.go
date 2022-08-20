package models

//Port structure info
type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"time_zone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type PortDetails struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Port Port   `json:"port"`
}
