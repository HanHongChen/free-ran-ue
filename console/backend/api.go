package backend

import (
	"net/http"

	"github.com/Alonza0314/free-ran-ue/console/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/gin-gonic/gin"
)

func (r *Console) handleConsoleLogin(c *gin.Context) {
	r.LoginLog.Infoln("Attempting to login")

	var loginRequest model.ConsoleLoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		r.LoginLog.Warnf("Failed to bind login request: %v", err)
		c.JSON(http.StatusBadRequest, model.ConsoleLoginResponse{
			Message: "Invalid request format",
		})
		return
	}

	if loginRequest.Username != r.username || loginRequest.Password != r.password {
		r.LoginLog.Warnf("Invalid credentials")
		c.JSON(http.StatusUnauthorized, model.ConsoleLoginResponse{
			Message: "Invalid credentials",
		})
		return
	}

	token, err := util.CreateJWT(r.jwt.secret, c.ClientIP(), r.jwt.expiresIn, nil)
	if err != nil {
		r.LoginLog.Errorf("Failed to create JWT: %v", err)
		c.JSON(http.StatusInternalServerError, model.ConsoleLoginResponse{
			Message: "Failed to create JWT",
		})
		return
	}

	c.JSON(http.StatusOK, model.ConsoleLoginResponse{
		Message: "Login successful",
		Token:   token,
	})

	r.LoginLog.Infoln("Login successful")
}

func (r *Console) handleConsoleLogout(c *gin.Context) {
	r.LogoutLog.Infoln("Attempting to logout")

	c.JSON(http.StatusOK, model.ConsoleLogoutResponse{
		Message: "Logout successful",
	})

	r.LogoutLog.Infoln("Logout successful")
}
