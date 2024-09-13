package main

import (
	"FinancialAssistanceScheme/app"
	"FinancialAssistanceScheme/config"
	"FinancialAssistanceScheme/middleware/db"
	"FinancialAssistanceScheme/middleware/redis"
)

func main() {
	// Initialize configuration
	appConfig := config.InitConfig()

	// Initialize Database connection
	db.InitDatabase(appConfig)

	// Initialize Redis connection
	redis.InitRedis(appConfig)

	// run app server
	app.RunApp()
}
