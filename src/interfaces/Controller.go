package interfaces

import (
	"net/http"
)

type Controller interface {
	Test(req *http.Request, res http.ResponseWriter)
}
