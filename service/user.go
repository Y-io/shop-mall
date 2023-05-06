package service

import (
	"shop-mall/global"
	"shop-mall/model"
	"shop-mall/model/do"
	"shop-mall/utility"
)

type PageInfoReq struct {
	page  int
	pSize int
}

type UserRes struct {
	id       int32
	password string
	mobile   string
}

type UserListRes struct {
	List  []UserRes
	Page  int
	Size  int
	Total int
}

type UserService interface {
	GetUserList(req model.UserGetListInput) (rsp *model.UserGetListOutput, err error)
}

var localUser UserService

func User() UserService {
	if localUser == nil {
		panic("没有初始化用户服务")
	}
	return localUser
}

type (
	sUser struct{}
)

func init() {
	localUser = &sUser{}
}

func (s *sUser) GetUserList(in model.UserGetListInput) (rsp *model.UserGetListOutput, err error) {
	var users []do.User
	result := global.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	rsp = &model.UserGetListOutput{}

	rsp.Total = int32(result.RowsAffected)

	rsp.Page = in.Page
	rsp.PageSize = in.PageSize

	global.DB.Scopes(utility.Paginate(int(in.Page), int(in.PageSize))).Find(&users)

	for _, user := range users {
		userInfo := model.UserGetListItem{
			ID:        user.ID,
			NickName:  user.NickName,
			Gender:    user.Gender,
			Mobile:    user.Mobile,
			CreatedAt: user.CreatedAt,
			UpdateAt:  user.UpdateAt,
			Role:      user.Role,
		}

		if user.Birthday != nil {
			userInfo.Birthday = uint64(user.Birthday.Unix())
		}

		rsp.List = append(rsp.List, userInfo)
	}

	return rsp, nil
}
