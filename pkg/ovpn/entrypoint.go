package ovpn

import (
	"strings"

	"github.com/borchero/meerkat-operator/pkg/ovpn/static"
)

// EntrypointValues describes the set of values required to render the OVPN server entrypoint.
type EntrypointValues struct {
	Routes []string
}

// GetEntrypoint returns the file that should be used for starting the VPN server. It sets up IP
// tables according to the given configuration.
func GetEntrypoint(values EntrypointValues) (string, error) {
	entrypoint, err := renderTemplate("entrypoint", static.TemplateEntrypoint, values)
	if err != nil {
		return "", err
	}
	return strings.Trim(entrypoint, "\n\t\r "), nil
}
