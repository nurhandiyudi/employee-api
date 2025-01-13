package logic

import (
	"context"
	"net/http"
	//"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"employee-api/config"
	"employee-api/model"
	"employee-api/util"
)

var employeeCollection = config.GetCollection("employees")

func AddEmployee(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var newEmployee model.Employee
	if err := c.Bind(&newEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if employee name already exists in the department
	existingEmployee := model.Employee{}
	err := employeeCollection.FindOne(ctx, bson.M{"name": newEmployee.Name, "department": newEmployee.Department}).Decode(&existingEmployee)
	if err == nil {
		// If an employee with the same name in the same department exists
		return c.JSON(http.StatusConflict, map[string]string{"error": "Employee name must be unique within the department"})
	}

	result, err := employeeCollection.InsertOne(ctx, newEmployee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add employee"})
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateEmployee(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var updatedEmployee model.Employee
	if err := c.Bind(&updatedEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if employee name already exists in the department
	existingEmployee := model.Employee{}
	err := employeeCollection.FindOne(ctx, bson.M{"name": updatedEmployee.Name, "department": updatedEmployee.Department, "_id": bson.M{"$ne": objID}}).Decode(&existingEmployee)
	if err == nil {
		// If an employee with the same name in the same department exists
		return c.JSON(http.StatusConflict, map[string]string{"error": "Employee name must be unique within the department"})
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedEmployee}

	_, err = employeeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update employee"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Employee updated successfully"})
}

func GetEmployees(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := util.BuildFilter(c)
	cursor, err := employeeCollection.Find(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch employees"})
	}
	defer cursor.Close(ctx)

	employees := []model.Employee{}
	if err = cursor.All(ctx, &employees); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse employees"})
	}

	return c.JSON(http.StatusOK, employees)
}

func CalculateAverageSalary(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	department := c.QueryParam("department")
	if department == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Department is required"})
	}

	pipeline := []bson.M{
		{"$match": bson.M{"department": department, "status": "active"}},
		{"$group": bson.M{"_id": nil, "averageSalary": bson.M{"$avg": "$salary"}}},
	}

	cursor, err := employeeCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to calculate average salary"})
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil || len(result) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No data found"})
	}

	return c.JSON(http.StatusOK, result[0])
}
func UpdateEmployeeStatus(c echo.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
    }

    status := c.QueryParam("status")
    if status != "active" && status != "inactive" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid status, must be 'active' or 'inactive'"})
    }

    filter := bson.M{"_id": objID}
    update := bson.M{"$set": bson.M{"status": status}}

    _, err = employeeCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update employee status"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Employee status updated successfully"})
}
func GetEmployeesByStatus(c echo.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    status := c.QueryParam("status")
    if status != "active" && status != "inactive" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid status, must be 'active' or 'inactive'"})
    }

    filter := bson.M{"status": status}
    cursor, err := employeeCollection.Find(ctx, filter)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch employees"})
    }
    defer cursor.Close(ctx)

    employees := []model.Employee{}
    if err = cursor.All(ctx, &employees); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse employees"})
    }

    return c.JSON(http.StatusOK, employees)
}