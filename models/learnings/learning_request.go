package learnings

type LearningRegister struct {
	CourseId uint `json:"courseId"`
	UserId   uint `json:"userId"`
}

type LearningUpdater struct {
	CourseId uint `form:"courseId"`
	UserId   uint `form:"userId"`
}
