package main

import (
	"context"
	"fmt"

	"github.com/deepakraj5/02-Go-pg/src/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api/v1", sampleHandler)

	app.Get("/api/v1/student", studentName)

	app.Listen(":3000")

}

func sampleHandler(c *fiber.Ctx) error {

	return c.Status(200).JSON(&fiber.Map{
		"message": "app running",
	})

}

type Student struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func studentName(c *fiber.Ctx) error {

	// var name string

	// db().QueryRow(context.Background(), "SELECT name FROM student").Scan(&name)

	// return c.Status(200).JSON(&fiber.Map{
	// 	"name": name,
	// })

	conn := db.Db

	var studentArray [4]Student

	if rows, err := conn().Query(context.Background(), "SELECT * FROM student"); err != nil {
		fmt.Println(err)
		// return
	} else {
		// defer db().Close()

		var std Student

		for rows.Next() {
			rows.Scan(&std.id, &std.name)
			studentArray[std.id] = std
		}
		if rows.Err() != nil {
			fmt.Println(err)
		}

	}

	fmt.Println(studentArray)

	return c.Status(200).JSON(studentArray)

}
