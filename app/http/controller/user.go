package controller

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"self/app/models"
	"strings"
	"time"
)

func CreateAdminUser(c *gin.Context) {

	var d models.AdminUser
	c.BindJSON(&d)
	err := models.AdminUser{}.CreateA(&d)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(d))
	}
}

func GetAdminUserList(c *gin.Context) {

	claims := c.MustGet("claims").(*CustomClaims)
	if claims != nil {
	}

	AdminUserList, err := models.AdminUser{}.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(AdminUserList))
	}
}

func UpdateAAdminUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}
	d, err := models.AdminUser{}.GetA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
		return
	}
	c.BindJSON(&d)

	err = models.AdminUser{}.UpdateA(d)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(d))
	}
}

func DeleteAAdminUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, Fail("无效的id"))
		return
	}

	err := models.AdminUser{}.DeleteA(id)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, Success(id+":deleted"))
	}
}

func Login(c *gin.Context) {
	var loginReq models.LoginReq
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err!=nil {
		c.JSON(http.StatusOK,Fail("传入参数有误"))
		return
	}
	err = json.Unmarshal(bodyBytes, &loginReq)
	if bodyBytes != nil && err == nil {
		// 登陆逻辑校验(查库，验证用户是否存在以及登陆信息是否正确)
		isPass, user, err := models.LoginCheck(loginReq)
		if err!=nil {
			c.JSON(http.StatusOK,Fail("请查证用户密码后再输入 "+err.Error()))
			return
		}
		if isPass {
			generateToken(c, *user)
		} else {
			c.JSON(http.StatusOK, Fail("验证失败 "+err.Error()))
		}

	} else {
		c.JSON(http.StatusOK, Fail("用户数据解析失败 "+ err.Error()))
	}
}

func Info(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, Fail("请求未携带token，无权限访问"))
		c.Abort()
		return
	}
	token = strings.Split(token, " ")[1]
	// 初始化一个JWT对象实例，并根据结构体方法来解析token
	j := NewJWT()
	// 解析token中包含的相关信息(有效载荷)
	claims, err := j.ParserToken(token)

	if err != nil {
		if err.Error() == "TokenExpired" {
			c.JSON(http.StatusOK, Fail("token授权已过期，请重新申请授权"))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, Fail(err.Error()))
		c.Abort()
		return
	}

	user, err := models.AdminUser{}.GetAByName(claims.Name)
	if err!=nil || user == nil {
		c.JSON(http.StatusOK,Fail("username 为空 或用户查询失败："+err.Error()))
		return
	}
	c.JSON(http.StatusOK, Success(user))
}

func generateToken(c *gin.Context, user models.AdminUser) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := NewJWT()

	// 构造用户claims信息(负荷)
	claims := CustomClaims{
		user.Username,
		user.Role,
		jwt.StandardClaims{
			IssuedAt:  int64(time.Now().Unix()),        // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*24), // 签名过期时间
			Issuer:    "monitor",                       // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	log.Println(token)
	// 封装一个响应数据,返回用户名和token
	type LoginResult struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	}
	data := LoginResult{
		Name:  user.Username,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
	return

}
