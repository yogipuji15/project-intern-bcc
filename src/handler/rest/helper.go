package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"project-intern-bcc/src/business/entity"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
)

func (r *rest) SuccessResponse(c *gin.Context, code int,message string, data interface{}){
	response:= entity.Response{
		Meta: entity.Meta{
			Message: message,
			Code: code,
			IsSuccess: true,
		},
		Data: data,
	}
	c.JSON(code,response)
}

func (r *rest) ErrorResponse(c *gin.Context, code int, err error, data interface{}){
	response:= entity.Response{
		Meta: entity.Meta{
			Message: err.Error(),
			Code: code,
			IsSuccess: false,
		},
		Data: data,
	}
	c.AbortWithStatusJSON(code, response)
}

func (r *rest) RequireAuth(c *gin.Context){
	tokenString,err:=c.Cookie("Authorization")
	if err != nil{
		r.ErrorResponse(c,http.StatusUnauthorized,errors.New("Token is not found"),nil)
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRETTOKEN")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			r.ErrorResponse(c,http.StatusUnauthorized,errors.New("Token has expired"),nil)
			c.Abort()
			return
		}
		
		user,statusCode,err:=r.uc.User.GetById(claims["id"])
		if err!=nil{
			r.ErrorResponse(c,statusCode,err,user)
			c.Abort()
			return
		}
		
		if user.Role=="premium-user"{
			if float64(time.Now().Unix())>float64(user.PremiumDue.Unix()){
				result,status,err:=r.uc.User.UpdateToFreeUser(user.ID)
				if err!=nil{
					r.ErrorResponse(c,status,err,result)
					c.Abort()
					return
				}
				
			}
		}

		c.Set("user",claims["id"])
		c.Next()
	} else {
		r.ErrorResponse(c,http.StatusUnauthorized,errors.New("Token is invalid"),nil)
		c.Abort()
		return
	}
}

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter required")
	}
	if !uppercasePresent {
		appendError("uppercase letter required")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character required")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

func (h *rest) BindBody(ctx *gin.Context, body interface{}) error { 
	return ctx.ShouldBindWith(body, binding.Default(ctx.Request.Method, ctx.ContentType())) 
}