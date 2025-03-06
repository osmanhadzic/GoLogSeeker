/*
 * SPDX-FileCopyrightText: 2025 OCode
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	createDBQuery := "CREATE DATABASE logseekerdb;"
	if err := DB.Exec(createDBQuery).Error; err != nil {
		log.Println("Failed to create database (might already exist):", err)
	} else {
		log.Println("Database 'logseekerdb' created successfully!")
	}

	dsn = "host=localhost user=postgres password=postgres dbname=logseekerdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to 'logseekerdb': %v", err)
	}

	log.Println("Connected to 'logseekerdb'!")
}

func GetDB() *gorm.DB {
	return DB
}
