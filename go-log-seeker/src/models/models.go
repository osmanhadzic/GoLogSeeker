/*
 * SPDX-FileCopyrightText: 2025 OCode
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package models

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
