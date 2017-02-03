package controllers

import "comic-go/models"

var (
	Catalog = models.Catalog{}.Shared()
)
