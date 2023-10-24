package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/routes/types"
	"github.com/go-my-admin/server/utils/auth"
	"github.com/go-my-admin/server/utils/bcrypt"
	"strings"
)

type HandleLoginRequest struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type HandleLoginResponse struct {
	Error bool   `json:"error"`
	Token string `json:"token"`
}

// HandleLogin
// @BasePath /auth/
// @Summary Login
// @Description Logs in a user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body HandleLoginRequest true "Request body"
// @Success 200 {object} HandleLoginResponse "The user was logged in successfully"
// @Failure 400 {object} types.ErrorResponse "The request body is invalid or that some required fields are missing"
// @Failure 401 {object} types.ErrorResponse "Incorrect password"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /auth/login [post]
func HandleLogin(c *gin.Context) {
	var req HandleLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Invalid request",
		})
		return
	}

	// Get user from database
	db := database.InternalDb

	type userRow struct {
		Id             int    `json:"id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		HashedPassword string `json:"password"`
	}

	isEmail := strings.Contains(req.UsernameOrEmail, "@")
	var rows *sql.Rows
	if isEmail {
		rows, err = db.RunQueryWithParams("SELECT id, username, email, password FROM users WHERE email = $1", req.UsernameOrEmail)
	} else {
		rows, err = db.RunQueryWithParams("SELECT id, username, email, password FROM users WHERE username = $1", req.UsernameOrEmail)
	}

	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Internal server error" + err.Error(),
		})
		return
	}

	for rows.Next() {
		var user userRow
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword)
		if err != nil {
			c.JSON(500, types.ErrorResponse{
				Error:   true,
				Message: "Internal server error" + err.Error(),
			})
			return
		}

		// Check if password is correct
		if !bcrypt.ValidatePassword(req.Password, user.HashedPassword) {
			c.JSON(401, types.ErrorResponse{
				Error:   true,
				Message: "Incorrect password",
			})
			return
		}

		// Create JWT
		token, err := auth.CreateJWT(user.Username)
		if err != nil {
			c.JSON(500, types.ErrorResponse{
				Error:   true,
				Message: "Internal server error" + err.Error(),
			})
			return
		}

		c.JSON(200, HandleLoginResponse{
			Error: false,
			Token: token,
		})
		return
	}

}
