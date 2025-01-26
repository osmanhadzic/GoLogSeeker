/*
 * SPDX-FileCopyrightText: 2025 OCode
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package routes

import (
	"go-log-seeker/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/items", controllers.GetItems)
	r.GET("/items/:id", controllers.GetItem)
	r.POST("/items", controllers.CreateItem)
	r.PUT("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)
}
