package response

import "RestfulAPIElearningVideo/features/users"

type User struct {
	ID       uint    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func FromUserCore(req User) users.UserCore  {
	return users.UserCore{
		ID:          req.ID,
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
	}
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func ToUserLoginResponse(user users.UserCore) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: user.Token,
	}
}