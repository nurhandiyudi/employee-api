package router

import (
	"github.com/labstack/echo/v4"
	"employee-api/logic"
	"employee-api/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/api/employees", logic.AddEmployee, middleware.AccessControlMiddleware)
	e.PUT("/api/employees/:id", logic.UpdateEmployee, middleware.AccessControlMiddleware)
	e.GET("/api/employees", logic.GetEmployees)
	e.GET("/api/employees/salary/average", logic.CalculateAverageSalary)
	e.PUT("/api/employees/:id/status", logic.UpdateEmployeeStatus, middleware.AccessControlMiddleware)
	e.GET("/api/employees/status", logic.GetEmployeesByStatus)
}
