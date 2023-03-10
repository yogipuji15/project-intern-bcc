package main

import (
	"log"
	"project-intern-bcc/src/business/repository"
	"project-intern-bcc/src/business/usecase"
	"project-intern-bcc/src/handler/rest"
	"project-intern-bcc/src/lib/auth"
	"project-intern-bcc/src/lib/database/sql"
	"project-intern-bcc/src/lib/midtrans"
	"project-intern-bcc/src/lib/storage"

	"github.com/joho/godotenv"
)

func init(){
	err:=godotenv.Load()
	if err!=nil{
		log.Panic("Error loading .env file")
	}
	
	sql.ConnectDatabase()
	sql.SyncDatabase()
}

func main()  {
	auth:= auth.Init()
	storage:=storage.Init()
	midtrans:=midtrans.Init()
	repository := repository.Init(sql.DB,storage,midtrans)
	usecase := usecase.Init(storage,auth,repository)
	rest := rest.Init(usecase)

	
	rest.Run()
}