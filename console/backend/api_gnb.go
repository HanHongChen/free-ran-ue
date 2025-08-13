package backend

import (
	"fmt"
	"net/http"

	"github.com/Alonza0314/free-ran-ue/console/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/gin-gonic/gin"
)

func (cs *console) handleConsoleGnbInfo(c *gin.Context) {
	cs.GnbLog.Infoln("Attempting to register gNB")

	if err := authticate(c, cs.jwt.secret); err != nil {
		cs.AuthLog.Warnln(err)
		c.JSON(http.StatusUnauthorized, model.ConsoleGnbInfoResponse{
			Message: err.Error(),
		})
		return
	}

	var request model.ConsoleGnbInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		cs.GnbLog.Warnf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, model.ConsoleGnbInfoResponse{
			Message: "Failed to bind JSON",
		})
		return
	}

	uri := fmt.Sprintf("http://%s:%d%s", request.Ip, request.Port, API_GNB_INFO)

	response, err := util.SendHttpRequest(uri, API_GNB_INFO_METHOD, nil, nil)
	if err != nil {
		cs.GnbLog.Warnln(err)
		c.JSON(http.StatusInternalServerError, model.ConsoleGnbInfoResponse{
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
