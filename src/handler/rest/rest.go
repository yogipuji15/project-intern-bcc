package rest

import (
	"project-intern-bcc/src/business/usecase"

	"github.com/gin-gonic/gin"
)

type Rest interface{
	Run()
}

type rest struct {
	uc  *usecase.Usecase
	gin *gin.Engine
}

func Init(usecase *usecase.Usecase) Rest {
	r := &rest{
		uc:usecase,
		gin: gin.Default(),
	}

	r.Register()
	return r
}

func (r *rest) Run() {
	r.gin.Run()
}

func (r *rest) Register() {
	v1:=r.gin.Group("api/v1")
	user:=v1.Group("/user")
	{
		user.POST("/signup", r.SignUp)
		user.POST("/login", r.Login)
		user.GET("/signup/verification",r.Verification)
		r.gin.MaxMultipartMemory = 8 << 20
		// user.POST("/upload",r.UploadFileSupabase)
		user.GET("/speaker-category",r.GetAllCategories)
		user.GET("/search-speakers",r.GetAllSpeakers)
		user.GET("/search-sponsors",r.GetAllSponsors)
		user.GET("/speaker-details/:id",r.GetSpeakerById)
		user.GET("/profile",r.RequireAuth,r.GetUserById)
		user.GET("/reviews/:id",r.GetReviewsByUserId)
	}
	
}