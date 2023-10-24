package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/routes/types"
	"github.com/go-my-admin/server/utils/bcrypt"
	"net/mail"
	"strings"
)

type HandleSignUpRequest struct {
	Name     string
	Username string
	Email    string
	Password string
}

type HandleSignUpResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// HandleSignup
// @BasePath /auth/
// @Summary Register a user
// @Description Creates a new user if anyone with the same username or email exists
// @Tags auth
// @Accept json
// @Produce json
// @Param request body HandleSignUpRequest true "Request body"
// @Success 200 {object} HandleSignUpResponse "The user was registered successfully"
// @Failure 400 {object} types.ErrorResponse "The request body is invalid or that some required fields are missing"
// @Failure 409 {object} types.ErrorResponse "User with the given username or email already exists"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /auth/signup [post]
func HandleSignup(c *gin.Context) {
	var req HandleSignUpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Invalid request",
		})
		return
	}

	if req.Name == "" || req.Username == "" || req.Email == "" || req.Password == "" {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Missing required fields",
		})
		return
	}

	email, err := mail.ParseAddress(req.Email)
	if err != nil {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Invalid email",
		})
		return
	}
	req.Email = email.Address

	if strings.Contains(req.Username, " ") {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Username cannot contain spaces",
		})
		return
	}

	db := database.InternalDb

	rows, err := db.RunQueryWithParams("SELECT * FROM users WHERE username = $1 OR email = $2", req.Username, req.Email)
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Error checking if user exists: " + err.Error(),
		})
		return
	}

	if rows.Next() {
		c.JSON(409, types.ErrorResponse{
			Error:   true,
			Message: "User with the given username or email already exists",
		})
		return
	}

	hashed, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Error hashing password: " + err.Error(),
		})
		return
	}
	_, err = db.RunQueryWithParams("INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4)", req.Name, req.Username, req.Email, hashed)
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Error creating user: " + err.Error(),
		})
		return
	}

	c.JSON(200, HandleSignUpResponse{
		Error:   false,
		Message: "User registered successfully",
	})

}
