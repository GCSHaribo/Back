package model

//User struct
type User struct {
	ID         string `json:"id" bson:"id"`
	Username   string `json:"username" bson:"username"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	IsTeacher  bool   `json:"is_teacher" bson:"is_teacher"`
	TotalScore int    `json:"total_score" bson:"total_score"`
	Token      string `json:"token"`
}

//Score struct for User
type Score struct {
	Email  string `json:"email"`
	Kind   int    `json:"kind"`
	Detail string `json:"detail"`
	Score  int    `json:"score"`
}
