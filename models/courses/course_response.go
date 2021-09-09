package courses

type CourseResponses struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    []Course `json:"data"`
}

type CourseResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Course `json:"data"`
}

type CourseResponseCount struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Jumlah  int64  `json:"jumlah"`
}
