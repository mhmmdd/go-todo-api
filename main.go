package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
	"go-todo-api/src/database"
	"go-todo-api/src/todo"
	"go-todo-api/src/utils"
	"os"
)

func main() {
	configFile := "./config.prod.yml"

	if len(os.Args) > 1 {
		if os.Args[1] == "dev" {
			configFile = "./config.yml"
		}
	}

	viper.SetConfigFile(configFile)
	viper.SetDefault("host", "0.0.0.0:8000")

	err := viper.ReadInConfig()
	if err != nil {
		switch e := err.(type) {
		case viper.ConfigFileNotFoundError:
			pterm.Warning.Printfln("not found conf file, use default")
		case *os.PathError:
			pterm.Warning.Printfln("not find conf file in %s", e.Path)
		default:
			pterm.Error.Printfln("load config fail:%v", err)
			return
		}
	}

	db := database.Connect(viper.GetString(`database.postgres`))
	database.AutoMigrate(db)

	r := database.SetupRedis(viper.GetString(`database.redis`))
	utils.SetupCacheChannel(r)

	app := fiber.New()

	todo.SetupRoute(app, db, r)

	app.Listen(viper.GetString("host"))
}
