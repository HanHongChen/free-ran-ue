package backend

import (
	"net/http"

	"github.com/Alonza0314/free-ran-ue/util"
)

const (
	APPLICATION_JSON = "application/json"

	API_GNB_INFO        = util.GNB_API_PREFIX + "/info"
	API_GNB_INFO_METHOD = http.MethodGet
)
