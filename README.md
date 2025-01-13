
# Employee Management API

This project is a Golang-based API for managing employee data, implemented using the Echo framework and MongoDB. The API supports CRUD operations, insights into employee salary data, and access control for sensitive operations.

---

## **Features**

- **Employee Management**: Add, update, and retrieve employee details.
- **Data Insights**: Calculate the average salary of employees in a specific department.
- **Access Control**: Protect sensitive routes with authentication middleware.
- **MongoDB Integration**: Store and query employee data efficiently.
- **Configuration File**: Easily manage database settings via `config.json`.

---

## **Directory Structure**

```
project/
├── main.go           # Entry point for the application
├── config/           # Configuration for MongoDB
│   ├── db.go         # MongoDB connection logic
│   └── config.json   # Database configuration file
├── models/           # Data structures and related logic
│   └── employees.go  # Employee model
├── routes/           # Router configurations
│   └── boot.go       # API route definitions
├── logic/            # Request handlers and business logic
│   └── employees.go  # Employee operations logic
├── middleware/       # Middleware for access control
│   └── middleware.go # Authentication middleware
├── utils/            # Utility functions
│   └── query.go      # Query construction helper
```

---

## **Setup Instructions**

### **Prerequisites**

- Go (1.18 or later)
- MongoDB (local or cloud instance)

### **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/employee-management-api.git
   cd employee-management-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the database:
   Update `config/config.json` with your MongoDB connection details:
   ```json
   {
     "mongoURI": "mongodb://localhost:27017",
     "database": "company",
     "user": "dbUser",
     "password": "dbPassword"
   }
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

---

## **API Endpoints**

### **Employee Management**

#### Add Employee
- **POST** `/api/employees`
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "department": "Engineering",
    "status": "active",
    "salary": 60000
  }
  ```
- **Response**:
  ```json
  {
    "insertedId": "<id>"
  }
  ```

#### Update Employee
- **PUT** `/api/employees/:id`
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "department": "Engineering",
    "status": "inactive",
    "salary": 65000
  }
  ```

#### Get Employees
- **GET** `/api/employees`
- **Query Parameters**:
  - `department` (optional)
  - `status` (optional)

#### Calculate Average Salary
- **GET** `/api/employees/salary/average`
- **Query Parameters**:
  - `department` (required)

---

## **Access Control**

Sensitive routes require an `Authorization` header with a valid token. Middleware checks the token and validates access.

Example:
```bash
Authorization: Bearer valid_token
```

---

## **Security Considerations**

- **Token-Based Authentication**: Protect sensitive operations with authentication middleware.
- **Input Validation**: Ensure all inputs are validated to prevent SQL injection and other attacks.
- **Secure Configurations**: Store secrets (e.g., database credentials) in environment variables or encrypted files.

---

## **Future Enhancements**

- Add unit and integration tests.
- Implement role-based access control (RBAC).
- Support for advanced filtering and sorting.
- Dockerize the application for easy deployment.

---

## **License**

This project is licensed under the MIT License. See the LICENSE file for details.

---

## **Contributions**

Contributions are welcome! Feel free to open issues or submit pull requests to improve this project.

---

## **API Endpoints**

### **1. Add Employee**
- **Method**: `POST`
- **Endpoint**: `/api/employees`
- **Description**: Adds a new employee to the database.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "department": "Engineering",
    "status": "active",
    "salary": 60000
  }
  ```
- **Response**:
  - **Success**: Returns the ID of the inserted employee.
  - **Error**: Conflict if the name already exists in the same department.

### **2. Update Employee**
- **Method**: `PUT`
- **Endpoint**: `/api/employees/:id`
- **Description**: Updates an existing employee's details.
- **Request Body**: Same as `Add Employee`.

### **3. Get Employees**
- **Method**: `GET`
- **Endpoint**: `/api/employees`
- **Description**: Retrieves a list of employees, supports filtering by department or status.

### **4. Calculate Average Salary**
- **Method**: `GET`
- **Endpoint**: `/api/employees/salary/average`
- **Description**: Calculates the average salary of employees in a specific department.
- **Query Parameters**:
  - `department` (required): The department to calculate the average salary for.

### **5. Update Employee Status**
- **Method**: `PUT`
- **Endpoint**: `/api/employees/:id/status`
- **Description**: Updates the status of an employee (active/inactive).
- **Query Parameter**:
  - `status` (required): The new status, either `active` or `inactive`.

### **6. Get Employees by Status**
- **Method**: `GET`
- **Endpoint**: `/api/employees/status`
- **Description**: Retrieves employees filtered by their status (active/inactive).
- **Query Parameter**:
  - `status` (required): The status to filter employees by, either `active` or `inactive`.
---

## **Access Control**

Sensitive routes require an `Authorization` header with a valid token. Middleware checks the token and validates access.

Example:
```bash
Authorization: Bearer valid_token
```

---

## **Security Considerations**

- **Token-Based Authentication**: Protect sensitive operations with authentication middleware.
- **Input Validation**: Ensure all inputs are validated to prevent SQL injection and other attacks.
- **Secure Configurations**: Store secrets (e.g., database credentials) in environment variables or encrypted files.

---
