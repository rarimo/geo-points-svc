package data

import (
	"github.com/rarimo/geo-points-svc/internal/data/evtypes/models"
)

type EventTypesQ interface {
	New() EventTypesQ
	Insert(...models.EventType) error
	Update(fields map[string]any) ([]models.EventType, error)
	Transaction(func() error) error

	Select() ([]models.EventType, error)
	Get(name string) (*models.EventType, error)
	FilterByNames(...string) EventTypesQ
}
