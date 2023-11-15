package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"postmanapi/models"
	"postmanapi/storage"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

// @Summary Admin login
// @Description Logs in the admin user
// @Accept json
// @Produce json
// @Param username query string true "Admin username" Format(username)
// @Param password query string true "Admin password" Format(password)
// @Success 200 {string} string "Successfully Login to Admin panel"
// @Failure 401 {object} ErrorResponse
// @Router /admin/login [post]
func AdminLoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	_, err := storage.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		if err.Error() == "User not found" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{"Username or password invalid"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Login to Admin panel"})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetProductsHandler returns a list of products
// @Summary Get products
// @Description Get a list of products in the admin panel
// @Produce json
// @Success 200 {array} models.Products
// @Router /admin_panel/products [get]
func GetProductsHandler(c *gin.Context) {
	var products []models.Products
	storage.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

// AddProductHandler adds a new product
// @Summary Add product
// @Description Add a new product in the admin panel
// @Accept json
// @Produce json
// @Param product body models.Products true "Product object"
// @Success 201 {object} Response
// @Router /admin_panel/products/add_product [post]
func AddProductsHandler(c *gin.Context) {
	var product models.Products
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert product into the database
	storage.DB.Create(&product)

	c.JSON(http.StatusCreated, Response{"Product added successfully"})
}

// @Summary Delete a Product by ID
// @Description Delete a User by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /products/{id} [delete]
func DeleteProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := storage.DB.Delete(&models.Products{}, id)
	if result.Error != nil {
		fmt.Println("Error deleting product:", result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// @Summary Update a Product by ID
// @Description Update a Product by its ID
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param updatedUser body models.Products true "Updated Product object"
// @Success 200 {object} models.User
// @Router /products/{id} [put]
func UpdateProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	var updatedProduct models.Products
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var product models.Products
	result := storage.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	storage.DB.Model(&result).Updates(updatedProduct)

	c.JSON(http.StatusOK, result)
}

// @Summary Get all User
// @Description Get a list of all Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsersHandler(c *gin.Context) {
	users := storage.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// @Summary Get a User by ID
// @Description Get a User by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	user := storage.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Create a new User
// @Description Create a new User with title and content
// @Accept json
// @Produce json
// @Param User body models.User true "User object"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	storage.AddUser(user)
	c.JSON(http.StatusCreated, user)
}

// @Summary Update a User by ID
// @Description Update a User by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param updatedUser body models.User true "Updated User object"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	storage.UpdateUser(id, updatedUser)
	c.JSON(http.StatusOK, updatedUser)
}

// @Summary Delete a User by ID
// @Description Delete a User by its ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	storage.DeleteUser(id)
	c.Status(http.StatusNoContent)
}
