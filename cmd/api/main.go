package main

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gustavoteixeira8/url-shortener/src/db"
	"github.com/gustavoteixeira8/url-shortener/src/routes"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
	"github.com/sirupsen/logrus"
)

func SetupServer() error {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		EnableIPValidation:      true,
		Concurrency:             runtime.NumCPU(),
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(recover.New(recover.ConfigDefault))
	app.Use(cors.New())
	app.Use(helmet.New())

	routes.SetupRoutes(app)

	serverPort := utils.GetEnv("SERVER_PORT")

	if serverPort == "" {
		serverPort = "3000"
	}

	logrus.Infof("Listening on port %v using %d CPUs", serverPort, runtime.NumCPU())

	err := app.Listen(fmt.Sprintf(":%s", serverPort))

	if err != nil {
		return err
	}

	return nil
}

func main() {
	utils.LoadDotenv()

	db.SetupDatabase()
	db.RunMigration()

	logrus.Fatal(SetupServer())
}
