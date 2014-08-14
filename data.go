package interlock

type (
	InterlockData struct {
		// these are custom vals for upstreams
		Port         int      `json:"port,omitempty"`
		AliasDomains []string `json:"alias_domains,omitempty"`
		Warm         bool     `json:"warm,omitempty"`
		// these are custom vals for hosts
		Check string `json:"check,omitempty"`
	}
)
