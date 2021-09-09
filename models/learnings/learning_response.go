package learnings

type LearningResponses struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Learning `json:"data"`
}

type LearningResponseSingle struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Learning `json:"data"`
}
