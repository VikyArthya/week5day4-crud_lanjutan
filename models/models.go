// models/models.go
package models

type Admin struct {
	AdminID  int    `json:"admin_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Student struct {
	StudentID int    `json:"student_id"`
	AdminID   int    `json:"admin_id"`
	ClassID   int    `json:"class_id"`
	Grade     string `json:"grade"`
}

type Mentor struct {
	MentorID  int    `json:"mentor_id"`
	AdminID   int    `json:"admin_id"`
	Expertise string `json:"expertise"`
}

type Material struct {
	MaterialID  int    `json:"material_id"`
	ClassID     int    `json:"class_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
}

type Schedule struct {
	ScheduleID int    `json:"schedule_id"`
	ClassID    int    `json:"class_id"`
	DateTime   string `json:"date_time"`
}

type Announcement struct {
	AnnouncementID int    `json:"announcement_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at"`
	VisibleTo      string `json:"visible_to"`
}
