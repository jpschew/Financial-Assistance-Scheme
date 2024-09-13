package app

import (
	"FinancialAssistanceScheme/controller"
	"FinancialAssistanceScheme/middleware/jwt"
	"FinancialAssistanceScheme/middleware/log"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"os"
)

func RunApp() {
	app := iris.New()               // Create new iris application
	app.Validator = validator.New() // initialize the validator
	app.Use(Cors)
	app.Use(recover.New()) // use recover middleware to prevent server from crashing on
	//app.Use(log.AccessLogMiddleware) // middleware for access log
	//app.Use(HTTPAuthMiddleware) // middleware to perform various tasks like authentication, logging and error handling before a request reaches actual route

	// initialized access log
	ac := log.InitAccessLog(os.Stdout)
	defer ac.Close()
	app.UseRouter(ac.Handler)

	// api test endpoint
	app.Get("/api/test", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "This is a test endpoint"})
	})

	// admin api
	app.Post("api/create_admin", controller.CreateAdmin)
	app.Post("api/login", controller.Login)
	app.Post("api/logout", controller.Logout)

	// applicants api
	app.Get("api/applicants", jwt.JWTAuthMiddleware, controller.GetAllApplicants)
	app.Post("api/applicants", jwt.JWTAuthMiddleware, controller.CreateApplicant)

	// schemes api
	app.Get("api/schemes", jwt.JWTAuthMiddleware, controller.GetAllSchemes)
	app.Post("api/schemes", jwt.JWTAuthMiddleware, controller.CreateScheme)
	app.Get("api/schemes/eligible", jwt.JWTAuthMiddleware, controller.GetEligibleScheme)

	// applications api
	app.Get("api/applications", jwt.JWTAuthMiddleware, controller.GetAllApplications)
	app.Post("api/applications", jwt.JWTAuthMiddleware, controller.CreateApplication)
	app.Post("api/update_application", jwt.JWTAuthMiddleware, controller.UpdateApplication)

	app.Listen(":8080")
}

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}
