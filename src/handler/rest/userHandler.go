package rest

import (
	"errors"
	"net/http"

	// "os"
	"project-intern-bcc/src/business/entity"
	"github.com/gin-gonic/gin"
	// storage_go "github.com/supabase-community/storage-go"
)

func (h *rest) Test(c *gin.Context){
	// data,err:= h.uc.User.Test(c)
	// fmt.Println(err)

	h.SuccessResponse(c, http.StatusOK,"Test Berhasil", "data1")
}


func (h *rest) SignUp(c *gin.Context){
	var body entity.UserSignup
	if err := c.BindJSON(&body); err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}
	
	if err:=verifyPassword(body.Password);err!=nil{
		h.ErrorResponse(c,http.StatusBadRequest, err,err.Error())
		return
	}

	result,statusCode,err:=h.uc.User.SignUp(body)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return 
	}
	h.SuccessResponse(c, statusCode,"Check your email to activate your account", result)
}

func (h *rest) Verification(c *gin.Context){
	token:=c.Query("token")
	
	result, statusCode, err:= h.uc.User.UserVerification(token)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return 
	}

	h.SuccessResponse(c, statusCode,"User registered successfully, your account is active", result)
}

func (h *rest) Login(c *gin.Context){
	var body entity.UserLogin
	if err := c.BindJSON(&body); err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}

	result,statusCode,err:=h.uc.User.Login(c,body)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return
	}

	h.SuccessResponse(c, statusCode, "User login successfully",result)
}

func (h *rest) GetUserById(c *gin.Context){
	id,exist:=c.Get("user")
	if exist==false{
		h.ErrorResponse(c,http.StatusUnauthorized,errors.New("User ID is not found in token"),"User ID doesn't exist")
		return
	}

	result,statusCode,err:=h.uc.User.GetById(id)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,"User is not found")
		return
	}

	h.SuccessResponse(c,statusCode,"Querying user's data successfully",result)
}


// func (h *rest) UploadFileSupabase(c *gin.Context){
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		h.ErrorResponse(c,400,err,"Failed to read file")
// 		return
// 	}
	
// 	link, statusCode, err := h.uc.User.UploadFile(c,file)
// 	if err != nil {
// 		h.ErrorResponse(c,statusCode,err,"Failed to upload file (Maximum file size: 3 MB)")
// 		return
// 	}
// 	h.SuccessResponse(c, statusCode, "Upload File successfully",link)
// }

// func (h *rest) UploadImage(c *gin.Context){
// 	image, err := c.FormFile("image")
// 	if err != nil {
// 		h.ErrorResponse(c,http.StatusBadRequest,err,"Failed to read image")
// 		return
// 	}

// 	imageIo, err := image.Open()
// 	if err != nil {
// 		h.ErrorResponse(c,http.StatusBadRequest,err,"Failed to open image")
// 		return
// 	}
	 
// 	client := storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SERVICE_TOKEN"), nil)
// 	client.UploadFile("image2", image.Filename, imageIo)
	

// 	c.JSON(http.StatusOK, gin.H{
// 		"success":    true,
// 		"statusCode": http.StatusOK,
// 		"message":    "successfully upload file and description",
// 		"linkImage":  os.Getenv("BASE_URL") + image.Filename,
// 	})
// }

