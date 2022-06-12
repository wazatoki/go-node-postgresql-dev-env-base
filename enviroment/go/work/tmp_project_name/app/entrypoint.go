package app

import (
	"tmp_project_name/app/infrastructures/echo"
	"tmp_project_name/app/infrastructures/postgresql"
	"tmp_project_name/app/infrastructures/viper"
)

// StartApp entoru point of application
func StartApp() {
	// configuration
	viper.SetupAppConfig()
	postgresql.Migrate()
	echo.StartEcho()
}
