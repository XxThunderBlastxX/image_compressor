package main

import (
	"github.com/XxThunderBlastxX/utils"
	"github.com/gofiber/fiber/v2"
	"io"
)

func main() {
	app := fiber.New()

	app.Static("/uploads", "./uploads")
	app.Post("/", func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("picture")
		if err != nil {
			panic(err)
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()

		buffer, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		errDir := utils.CreateFolder("uploads")
		if errDir != nil {
			panic(errDir)
		}

		filename, err := utils.ImageProcessing(buffer, 30, "uploads")
		if err != nil {
			panic(err)
		}

		return c.JSON(fiber.Map{
			"picture": "http://localhost:3000/uploads/" + filename,
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
