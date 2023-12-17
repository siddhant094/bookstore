package controllers

import (
	// "encoding/json"
	"go-mysql/models"

	"gofr.dev/pkg/gofr"
)

var NewBook models.Book 

func GetBooks(ctx *gofr.Context) (interface{}, error) {
	return "Hello World! /book", nil
}
