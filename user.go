package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersList struct {
	Id     int `json:"-"`
	UserId int `json:"userId"`
	ListId int `json:"listId"`
}
