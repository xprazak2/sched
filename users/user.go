package users

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}

type UserModel struct {
	ID int `gorm:"primary_key"`
	Name string
	Surname string
	Email string `gorm:"column:email;unique_index"`
}

func (self *UserModel) ToResponse() UserResponse {
	return UserResponse{
		ID: self.ID,
		Name: self.Name,
		Surname: self.Surname,
		Email: self.Email,
	}
}


func ToResponse(users[] UserModel) []UserResponse {
	var response[] UserResponse
	for _, user := range(users) {
		response = append(response, user.ToResponse())
	}
	return response
}
