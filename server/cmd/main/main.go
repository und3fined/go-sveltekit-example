/**
 * File: main.go
 * Project: cmd
 * File Created: 06 Jan 2022 11:02:21
 * Author: und3fined (me@und3fined.com)
 * -----
 * Last Modified: 08 Jan 2022 19:25:38
 * Modified By: und3fined (me@und3fined.com)
 * -----
 * Copyright (c) 2022 und3fined.com
 */
package main

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"

	"gosvelte/services/sveltekit"
	"gosvelte/utils/path"
)

func main() {
	root := path.Root()
	staticPath := filepath.Join(root, "./dest/client/static")
	prerenderPath := filepath.Join(root, "./dest/client/prerendered")
	appPath := filepath.Join(root, "./dest/client/app")

	app := fiber.New()
	kit := sveltekit.NewKit(sveltekit.UseMain("./dest/main.js"))

	app.Static("/", staticPath)
	app.Static("/prerendered", prerenderPath)
	app.Static("/app", appPath)

	app.Use(func(c *fiber.Ctx) error {
		resp, err := kit.Render(c.OriginalURL(), c.Route().Method, c.GetReqHeaders(), c.Body())
		if err != nil {
			log.Printf("%+v", err)
		}

		for key, val := range resp.Headers {
			c.Set(key, val)
		}

		c.Response().Header.SetStatusCode(int(resp.Status))
		return c.SendString(resp.Body)
	})

	log.Fatal(app.Listen(":3000"))
}
