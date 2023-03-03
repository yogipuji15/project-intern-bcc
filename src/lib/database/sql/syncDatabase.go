package sql

import(
	"project-intern-bcc/src/business/entity"
)

func SyncDatabase(){
	DB.AutoMigrate(
		&entity.Users{},
		&entity.Categories{},
		&entity.Companies{},
		&entity.Orders{},
		&entity.Payments{},
		&entity.Proposals{},
		&entity.Reviews{},
		&entity.Roles{},
		&entity.Schedules{},
		&entity.Speakers{},
	)
}