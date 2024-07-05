package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func NewGetEventType(r *http.Request) (name string, err error) {
	name = chi.URLParam(r, "id")
	return name, validation.Errors{"id": validation.Validate(name, validation.Required)}.
		Filter()
}
