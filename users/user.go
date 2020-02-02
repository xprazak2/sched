package users



var Users = []UserModel{
	UserModel{ "George", "Clooney", "g@c.com" },
	UserModel{ "Matt", "Damon", "m@d.com" },
	UserModel{ "Russel", "Crowe", "r@c.com" },
}

type UserModel struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}
