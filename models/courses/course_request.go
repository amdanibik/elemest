package courses

type CourseRegister struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Level    string `json:"level"`
	Price    int    `json:"price"`
	IsFree   int    `json:"isfree"`
}

type CourseUpdater struct {
	Name     string `form:"name"`
	Category string `form:"category"`
	Level    string `form:"level"`
	Price    int    `form:"price"`
	IsFree   int    `form:"isfree"`
}
