package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
	"project-intern-bcc/src/lib/auth"
	"project-intern-bcc/src/lib/storage"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Test(c *gin.Context) (entity.Users,error)
	SignUp(userInput entity.UserSignup) (interface{},int,error)
	Login(c *gin.Context, userInput entity.UserLogin) (interface{},int,error)
	UserVerification(token string)(interface{},int,error)
	GetById(id any) (entity.UserResponse,int,error)
	ConvertToUserResponse(user entity.Users) (entity.UserResponse)
	UpdateUserPremium(userId uint, premiumDue int) (interface{},int,error)
	UpdateToFreeUser(userId uint) (interface{},int,error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	auth auth.AuthInterface
	storage storage.StorageInterface
}

func NewUserUsecase(ur repository.UserRepository, auth auth.AuthInterface, storage storage.StorageInterface) UserUsecase {
	return &userUsecase{
		userRepository: ur,
		auth: auth,
		storage: storage,
	}
}

func (h *userUsecase) Test(c *gin.Context)(entity.Users,error){
	return h.userRepository.Test()
}

// func (h *userUsecase) SignUp(userInput entity.UserSignup) (interface{},int,error){
// 	if userInput.Password != userInput.ConfirmPass{
// 		return "The password confirmation doesn't match",http.StatusBadRequest,fmt.Errorf("Invalid password confirmation")
// 	}
	
// 	user := entity.Users{
// 		Email	 : userInput.Email,
// 		Username : userInput.Username,
// 	}

// 	userResponse:=entity.UserResponse{
// 		Email	 : user.Email,
// 		Username : user.Username,
// 	}

// 	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
// 	if err!= nil{
// 		return "Failed to hash password",http.StatusBadRequest,err
// 	}

// 	user.Password=string(hash)
// 	err = h.userRepository.Create(user)
// 	if err!= nil{
// 		return "Failed to create user",http.StatusInternalServerError,err
// 	}

// 	return userResponse,http.StatusOK,err
// }

func (h *userUsecase) SignUp(userInput entity.UserSignup) (interface{},int,error){
	if userInput.Password != userInput.ConfirmPass{
		return "The password confirmation doesn't match",http.StatusBadRequest,errors.New("Invalid password confirmation")
	}

	userOld,err:=h.userRepository.FindByEmail(userInput.Email)
	if err==nil{
		if userOld.Email==userInput.Email && userOld.IsActive==false{
			h.userRepository.Delete(userOld)
		}
	}
	
	user := entity.Users{
		Username : userInput.Username,
		RoleID    : 1,
		IsActive : false,
		Fullname: userInput.Fullname,
		Address: userInput.Address,
		Phone: userInput.Phone,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}

	token,err:=h.auth.GenerateToken(user)
	if err!=nil{
		return "Failed to create token",http.StatusBadRequest,err
	}
	
	hash, err := h.auth.HashPassword(userInput.Password)
	if err!= nil{
		return "Failed to hash password",http.StatusBadRequest,err
	}
	fmt.Println("===============")
	
	user.Password = string(hash)
	user.VerificationCode = token.Token
	user.Email = userInput.Email
	
	err = h.userRepository.Create(user)
	if err!= nil{
		return "Failed to create user",http.StatusInternalServerError,err
	}
	user.Role.Role="free-user"
	
	userResponse:=h.ConvertToUserResponse(user)
	
	err=h.auth.EmailVerification(user.Email,token.Token)
	if err!=nil{
		return "Failed to send email",http.StatusInternalServerError,err
	}

	return userResponse,http.StatusOK,err
}

func (h *userUsecase) UserVerification(token string)(interface{},int,error){
	user,err:=h.userRepository.FindUserByToken(token)
	if err!=nil{
		return "Failed to Querying User Data",http.StatusNotFound,err
	}

	user.IsActive=true
	err=h.userRepository.Update(user)
	if err!=nil{
		return "Failed to Updating User Data",http.StatusInternalServerError,err
	}

	userResponse := h.ConvertToUserResponse(user)

	return userResponse,http.StatusOK,err
}

func (h *userUsecase) Login(c *gin.Context, userInput entity.UserLogin)(interface{},int,error){
	user,err:=h.userRepository.FindByEmailUsername(userInput.EmailUsername)
	if err!= nil{
		return "Invalid email/username or password",http.StatusNotFound,err
	}

	if user.IsActive==false{
		return "Check your email to activate your account",http.StatusUnauthorized,errors.New("Your account is not active")
	}

	if user.ID==0{
		return "Invalid email/username or password",http.StatusBadRequest,err
	}

	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(userInput.Password))
	if err!= nil{
		return "Invalid email/username or password",http.StatusBadRequest,err
	}

	token,err:=h.auth.GenerateToken(user)
	if err!=nil{
		return "Failed to create token",http.StatusBadRequest,err
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",token.Token,3600*24*30,"","",false,true)

	return token,http.StatusOK,nil
}



func (h *userUsecase) GetById(id any) (entity.UserResponse,int,error){
	user,err:=h.userRepository.FindByIdWithRole(id)
	
	userResponse:= h.ConvertToUserResponse(user)
	if err!=nil{
		return userResponse,http.StatusNotFound,err
	}

	return userResponse,http.StatusOK,nil
}

func (h *userUsecase) ConvertToUserResponse(user entity.Users) (entity.UserResponse){
	return entity.UserResponse{
		ID 		 : user.ID,
		Email 	 : user.Email,
		Username : user.Username,
		Fullname : user.Fullname,
		Address  : user.Address,
		Phone	 : user.Phone,
		Role	 : user.Role.Role,
		PremiumDue: user.PremiumDue,
	}
}

func (h *userUsecase) UpdateUserPremium(userId uint, premiumDue int) (interface{},int,error){
	user,err:=h.userRepository.FindById(userId)
	if err!=nil{
		return "Failed to querying user's data",http.StatusNotFound,err
	}

	user.PremiumDue=time.Now().Add(time.Hour * 24 * 30 * time.Duration(premiumDue))
	user.RoleID=2
	
	err=h.userRepository.Update(user)
	if err!=nil{
		return "Failed to updating user's data",http.StatusInternalServerError,err
	}

	return "Upgrade Premium Successfully",http.StatusOK,nil
}

func (h *userUsecase) UpdateToFreeUser(userId uint) (interface{},int,error){
	user,err:=h.userRepository.FindById(userId)
	if err!=nil{
		return "Failed to querying user's data",http.StatusNotFound,err
	}

	user.RoleID=1
	
	err=h.userRepository.Update(user)
	if err!=nil{
		return "Failed to updating user's data",http.StatusInternalServerError,err
	}

	return "Update User's Role Successfully",http.StatusOK,nil
}