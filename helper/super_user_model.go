package helper

import (
	"github.com/Ardnh/go-ems/model/domain"
	"github.com/Ardnh/go-ems/model/web"
)

func ToUserResponseByUsername(user domain.SuperUser) web.SuperUserResponseByUserName {
	return web.SuperUserResponseByUserName{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Password:  user.Password,
	}
}
