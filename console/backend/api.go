package backend

import (
	"net/http"

	"github.com/Alonza0314/free-ran-ue/console/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/gin-gonic/gin"
)

func (cs *console) handleConsoleLogin(c *gin.Context) {
	cs.LoginLog.Infoln("Attempting to login")

	var loginRequest model.ConsoleLoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		cs.LoginLog.Warnf("Failed to bind login request: %v", err)
		c.JSON(http.StatusBadRequest, model.ConsoleLoginResponse{
			Message: "Invalid request format",
		})
		return
	}

	if loginRequest.Username != cs.username || loginRequest.Password != cs.password {
		cs.LoginLog.Warnf("Invalid credentials")
		c.JSON(http.StatusUnauthorized, model.ConsoleLoginResponse{
			Message: "Invalid credentials",
		})
		return
	}

	token, err := util.CreateJWT(cs.jwt.secret, c.ClientIP(), cs.jwt.expiresIn, nil)
	if err != nil {
		cs.LoginLog.Errorf("Failed to create JWT: %v", err)
		c.JSON(http.StatusInternalServerError, model.ConsoleLoginResponse{
			Message: "Failed to create JWT",
		})
		return
	}

	c.JSON(http.StatusOK, model.ConsoleLoginResponse{
		Message: "Login successful",
		Token:   token,
	})

	cs.LoginLog.Infoln("Login successful")
}

func (cs *console) handleConsoleLogout(c *gin.Context) {
	cs.LogoutLog.Infoln("Attempting to logout")

	c.JSON(http.StatusOK, model.ConsoleLogoutResponse{
		Message: "Logout successful",
	})

	cs.LogoutLog.Infoln("Logout successful")
}
