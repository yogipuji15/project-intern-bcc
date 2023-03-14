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
	r.gin.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})

	r.gin.Run()
}

func (r *rest) Register() {
	v1:=r.gin.Group("api/v1")
	
	user:=v1.Group("/user")
	{
		user.POST("/signup", r.SignUp)
		user.POST("/login", r.Login)
		user.GET("/signup/verification",r.Verification)
		// user.POST("/upload",r.UploadFileSupabase)
		user.GET("/order-history",r.RequireAuth,r.GetOrderHistory)
		user.POST("/upgrade-premium",r.RequireAuth,r.OrderPremiumAccount)
		user.POST("/order-history/pay-order",r.RequireAuth,r.CreateTransactionByOrderCode)
		user.GET("/speaker-category",r.GetAllCategories)
		user.GET("/search-speakers",r.GetAllSpeakers)
		user.GET("/speaker-details/:id",r.GetSpeakerById)
		user.GET("/speaker-details/reviews/:id",r.GetReviewsByUserId)
		user.GET("/speaker-details/schedules/:id",r.RequireAuth,r.GetAllSchedule)
		user.POST("/speaker-details/create-order",r.RequireAuth,r.CreateOrder)
		user.POST("/midtrans/update-order-status",r.CheckOrderTransaction)
		user.GET("/search-sponsors",r.GetAllSponsors)
		user.GET("/sponsor-details/:id",r.GetCompanyById)
		user.POST("/sponsor-details/apply-proposal",r.RequireAuth,r.UploadProposal)
		user.POST("/create-review",r.RequireAuth,r.PostReview)
		user.GET("/profile",r.RequireAuth,r.GetUserById)
		user.GET("/test",r.RequireAuth)
	}
}