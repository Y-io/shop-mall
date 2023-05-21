package service

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-mall/global"
	"shop-mall/middlewares"
	"shop-mall/model"
	"shop-mall/model/do"
	"shop-mall/model/entity"
	"shop-mall/utility"
	"strings"
	"time"
)

type UserService interface {
	GetUserList(c *gin.Context, in model.UserGetListInput) (*model.UserGetListOutput, error)
	CreateUser(c *gin.Context, in model.CreateUserInput)
	CheckMobileUnique(c *gin.Context, mobile int) error
	EncryptPassword(password string) string
	PasswordLogin(c *gin.Context, in model.PasswordLoginInput)
	GetUserByMobile(mobile string) (user *entity.User)
	CheckPassword(in model.CheckPasswordInput) bool
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

func (s *sUser) GetUserList(c *gin.Context, in model.UserGetListInput) (rsp *model.UserGetListOutput, err error) {
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
			UpdatedAt: user.UpdatedAt,
			Role:      user.Role,
		}

		if user.Birthday != nil {
			userInfo.Birthday = uint64(user.Birthday.Unix())
		}

		rsp.List = append(rsp.List, userInfo)
	}

	return rsp, nil
}

func (s *sUser) CreateUser(c *gin.Context, in model.CreateUserInput) {
	user := do.User{}
	out := model.CreateUserOutput{}

	// 判断用户是否已存在
	result := global.DB.Where(&do.User{Mobile: in.Mobile}).First(&user)

	if result.RowsAffected == 1 {
		c.JSON(http.StatusConflict, gin.H{
			"msg": "用户已存在",
		})
		return
	}

	user.Mobile = in.Mobile
	user.NickName = in.Mobile

	user.Password = s.EncryptPassword(in.Password)
	user.UpdatedAt = time.Now()

	result = global.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "创建用户失败",
		})
		return
	}

	result.Scan(&out)

	c.JSON(http.StatusCreated, out)

}

func (s *sUser) GetUserByMobile(mobile string) (user *entity.User) {

	// 判断用户是否已存在
	result := global.DB.Where(&do.User{Mobile: mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil
	}

	return user
}

func (s *sUser) PasswordLogin(c *gin.Context, in model.PasswordLoginInput) {

	userInfo := s.GetUserByMobile(in.Mobile)

	if userInfo == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	if checkPwd := s.CheckPassword(model.CheckPasswordInput{
		EncryptedPassword: userInfo.Password,
		Password:          in.Password,
	}); !checkPwd {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密码错误",
		})
		return
	}

	j := middlewares.NewJWT()

	claims := model.CustomClaims{
		ID:          uint(userInfo.ID),
		NickName:    userInfo.NickName,
		AuthorityId: uint(userInfo.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30 天过期时间
			Issuer:    "elvin",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成 token 失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        userInfo.ID,
		"nickName":  userInfo.NickName,
		"token":     token,
		"expiredAt": (time.Now().Unix() + 60*60*24*30) * 1000,
	})

}

// 检查手机号码是否唯一
func (s *sUser) CheckMobileUnique(c *gin.Context, mobile int) error {
	return nil
}

// 加密密码
func (s *sUser) EncryptPassword(passwordIn string) string {
	// 密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(passwordIn, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

// 检查密码
func (s *sUser) CheckPassword(in model.CheckPasswordInput) bool {

	pwdInfo := strings.Split(in.EncryptedPassword, "$")
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}

	return password.Verify(in.Password, pwdInfo[2], pwdInfo[3], options)
}
