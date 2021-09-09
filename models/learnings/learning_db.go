package learnings

import "time"

type Learning struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CourseId  uint      `json:"courseId"`
	UserId    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}
