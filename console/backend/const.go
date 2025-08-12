package backend

import (
	"net/http"

	"github.com/Alonza0314/free-ran-ue/util"
)

const (
	APPLICATION_JSON = "application/json"

	API_GNB_REGISTRATION        = util.GNB_API_PREFIX + "/registration"
	API_GNB_REGISTRATION_METHOD = http.MethodGet
)
