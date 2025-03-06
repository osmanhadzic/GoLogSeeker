/*
 * SPDX-FileCopyrightText: 2025 OCode
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/gin-gonic/gin"
	"go-log-seeker/src/data"
	"go-log-seeker/src/routes"
	"log"
)

func main() {
	database.InitDB()

	db := database.GetDB()

	if db != nil {
		log.Println("Database connection established successfully.")
	} else {
		log.Println("Failed to establish database connection.")
	}

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":5000") // listen and serve on 0.0.0.0:5000
}
