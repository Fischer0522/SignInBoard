package data

type SignInParam struct {
	Name string `form:"name" json:"name"`
	Message string `form:"password" json:"message"`
}

type UpdateParam struct {
	Id int `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Message string `form:"message" json:"message"`
}
