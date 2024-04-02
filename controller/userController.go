package controller

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/common"
	"TGPersonInfo/model"
	"TGPersonInfo/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// isUserExist: 如果用户存在，则返回true
func isUserExist(db *gorm.DB, username string) bool {
	var user model.UserInfo
	err := db.Where("username=?", username).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func UserRegister(c *gin.Context) {
	db := model.GetDB()
	logger := logging.GetLog()
	var userinfo model.UserInfo

	err := c.Bind(&userinfo)
	if err != nil {
		logger.Error(err)
		return
	}
	if isUserExist(db, userinfo.UserName) {
		logger.Error("用户已注册")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			"msg":  "用户已注册",
		})
		return
	}

	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(userinfo.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("加密错误")
		return
	}
	newUser := model.UserInfo{
		UserID:       utils.RandomInt(10),
		UserName:     userinfo.UserName,
		UserPassword: string(hashPassWord),
		UserSign:     userinfo.UserSign,
		Following:    userinfo.Following,
		Followers:    userinfo.Followers,
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 510,
			"msg":  "注册失败",
		})
		logger.Error(err)
		return
	}
	token, err := common.ReleaseToken(userinfo)
	if err != nil {
		logger.Error("Token发放失败")
		return
	}
	logger.Info("注册成功")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": token,
	})
}

// UserLogin : 用户登录
func UserLogin(c *gin.Context) {
	// var userinfo model.UserInfo
	db := model.GetDB()
	logger := logging.GetLog()

	// err := c.Bind(&userinfo)
	// if err != nil {
	// 	logger.Error(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code": 400,
	// 		"msg":  "请求格式错误",
	// 	})
	// 	return
	// }

	name := c.Query("name")
	password := c.Query("password")

	// 根据用户名，进行查询用户是否存在
	var user model.UserInfo
	if err := db.Where("username = ?", name).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "密码错误",
		})
		return
	}
	tokenString, err := common.ReleaseToken(user)
	if err != nil {
		logger.Error("发放token错误", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "发放token错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": tokenString,
	})

}
