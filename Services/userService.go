package Services

import (
	"shuyuxuan_blog_go/Models"
	"shuyuxuan_blog_go/Types"
	"time"
)
var userModel Models.UserModel

type sUser interface {
	upsertUser() error
}

type UserServer struct {

}

func (s UserServer) UpsertUser(req Types.UserReq) error {
	newUser := Types.User{}

	newUser.Name = req.Name
	newUser.Mail = req.Mail
	newUser.Phone = req.Phone
	newUser.Description = req.Description
	// 修改用户
	if req.Id != 0 {
		newUser.ID = req.Id
		var err error
		err = userModel.UpdateUser(newUser)
		return err
	} else {
		// 新增用户
		newUser.Status = 1
		t := time.Now()
		newUser.CreatedAt = t.Format("2006-01-02 15:04:05")
		var err error
		err = userModel.AddUser(newUser)
		return err
	}
}
