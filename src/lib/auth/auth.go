package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"os"
	"project-intern-bcc/src/business/entity"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthInterface interface {
	GenerateToken(user entity.Users) (UserAuthInfo, error)
	HashPassword(password string)([]byte,error)
	EmailVerification(email string,code string) (error)
	Hash512(input string) string
}

type auth struct {

}

func Init() AuthInterface {
	return &auth{}
}

func (a *auth) GenerateToken(user entity.Users) (UserAuthInfo, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id" : user.ID,
		"role" : user.RoleID,
		"is_guest" : false,
		"exp":time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Role:=""
	// if user.RoleID==0{
	// 	Role="Free user"
	// }else if user.RoleID==1{
	// 	Role="Premium user"
	// }
	userResponse:=entity.UserResponse{
		ID		 : user.ID,
		Email	 : user.Email,
		Username : user.Username,
		Fullname : user.Fullname,
		Phone	 : user.Phone ,
		Address  : user.Address,
		Role	 : user.Role.Role,
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETTOKEN")))
	userAuth := UserAuthInfo{
		User: userResponse,
		Token: tokenString,
	}
	return userAuth, err
}


func (a *auth) HashPassword(password string)([]byte,error){
	hash,err:=bcrypt.GenerateFromPassword([]byte(password), 10)
	return hash,err
}

func (a *auth) EmailVerification(email string,code string) (error){
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_FROM"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))

	message := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: CreatOr Email Verification\r\n\r\nDear User,\r\nThanks for starting the new CreatOr account creation process. We want to make sure it's really you. Please click the verification link below to activate your account! If you don’t want to create an account, you can ignore this message.\r\nVerification Link: https://yogi-puji.aenzt.tech/v1/user/signup/verification?token=%s\r\n", os.Getenv("EMAIL_FROM"), email, code)

	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("EMAIL_FROM"), []string{email}, []byte(message))
	return err
}

func (a *auth) Hash512(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	pass := hex.EncodeToString(hash.Sum(nil))
	return pass
}