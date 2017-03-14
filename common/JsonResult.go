package common

type JsonList struct {
	Total       uint32         `json:"total"`
	PerPage     uint8          `json:"per_page"`
	CurrentPage uint32         `json:"current_page"`
	Data        []*interface{} `json:"data"`
}

type JsonData struct {
	List JsonList `json:"list"`
}

type JsonResult struct {
	Status uint16   `json:"status"`
	Data   JsonData `json:"data"`
}
