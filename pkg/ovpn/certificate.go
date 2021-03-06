package ovpn

import (
	"strings"

	"github.com/borchero/meerkat-operator/pkg/ovpn/static"
)

// CertificateValues describes the set of values required to render OVPN certificates.
type CertificateValues struct {
	Secrets  CertificateSecrets
	Host     string
	Port     uint16
	Protocol string
	Security ConfigSecurity
}

// CertificateSecrets contains all relevant secrets for generating an OVPN client file.
type CertificateSecrets struct {
	TLSClientKey string
	TLSClientCrt string
	TLSCaCrt     string
	TLSAuth      string
}

// GetCertificate generates the OVPN certificate file that is given to the clients.
func GetCertificate(values CertificateValues) (string, error) {
	certificate, err := renderTemplate("certificate", static.TemplateClient, values)
	if err != nil {
		return "", err
	}
	return strings.Trim(certificate, "\n\t\r "), nil
}
