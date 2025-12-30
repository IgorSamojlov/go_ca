package funds

import (
	"net/http"
)

type HttpResponse struct {
	Error      string `json:"error,omitempty"`
	Collection any    `json:"collection,omitempty"`
	Record     any    `json:"record,omitempty"`
}

func (o *HttpResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
