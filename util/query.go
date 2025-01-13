package util

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func BuildFilter(c echo.Context) bson.M {
	filter := bson.M{}

	if department := c.QueryParam("department"); department != "" {
		filter["department"] = department
	}
	if status := c.QueryParam("status"); status != "" {
		filter["status"] = status
	}

	return filter
}
