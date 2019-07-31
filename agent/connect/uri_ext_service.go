package connect

import (
	"net/url"

	"github.com/hashicorp/consul/agent/structs"
)

type SpiffeIDExternalService struct {
	Host  string
	Path  string
}

// URI returns the *url.URL for this SPIFFE ID.
func (id *SpiffeIDExternalService) URI() *url.URL {
	var result url.URL
	result.Scheme = "spiffe"
	result.Host = id.Host
	result.Path = id.Path
	return &result
}

// CertURI impl.
func (id *SpiffeIDExternalService) Authorize(ixn *structs.Intention) (bool, bool) {
	switch ixn.SourceType {
	case structs.IntentionSourceExternalTrustDomain:
		return ixn.Action == structs.IntentionActionAllow, ixn.SourceName == "spiffe://" + id.Host
	case structs.IntentionSourceExternalURI:
		return ixn.Action == structs.IntentionActionAllow, ixn.SourceName == id.URI().String()
	}
	return false, false
}
