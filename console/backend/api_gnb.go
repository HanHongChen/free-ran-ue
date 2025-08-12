package backend

import (
	"fmt"
	"net/http"

	"github.com/Alonza0314/free-ran-ue/console/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/gin-gonic/gin"
)

func (cs *console) handleConsoleGnbRegistration(c *gin.Context) {
	cs.GnbLog.Infoln("Attempting to register gNB")

	authenticateHeader := c.GetHeader("Authorization")
	if authenticateHeader == "" {
		cs.AuthLog.Warnln("No authentication header")
		c.JSON(http.StatusUnauthorized, model.ConsoleGnbRegistrationResponse{
			Message: "No authentication header",
		})
		return
	}

	if _, err := util.ValidateJWT(authenticateHeader, cs.jwt.secret); err != nil {
		cs.AuthLog.Warnf("Failed to validate JWT: %v", err)
		c.JSON(http.StatusUnauthorized, model.ConsoleGnbRegistrationResponse{
			Message: err.Error(),
		})
		return
	}

	var request model.ConsoleGnbRegistrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		cs.GnbLog.Warnf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, model.ConsoleGnbRegistrationResponse{
			Message: "Failed to bind JSON",
		})
		return
	}

	uri := fmt.Sprintf("http://%s:%d%s", request.Ip, request.Port, API_GNB_REGISTRATION)

	response, err := util.SendHttpRequest(uri, API_GNB_REGISTRATION_METHOD, nil, nil)
	if err != nil {
		cs.GnbLog.Warnln(err)
		c.JSON(http.StatusInternalServerError, model.ConsoleGnbRegistrationResponse{
			Message: err.Error(),
		})
		return
	}

	for key, values := range response.Headers {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Data(response.StatusCode, APPLICATION_JSON, response.Body)
}
