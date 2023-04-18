package middleware

import (
	"myapp/data"

	"github.com/gardnerscot/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
