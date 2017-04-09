package launchpad

type Processor struct {
	Name                   string `json:"name"`
	Title                  string `json:"title"`
	Restricted             bool   `json:"restricted"`
	SupportsVirtualized    bool   `json:"supports_virtualized"`
	BuildByDefault         bool   `json:"build_by_default"`
	SupportsNonvirtualized bool   `json:"supports_nonvirtualized"`
	HTTPEtag               string `json:"http_etag"`
	SelfLink               string `json:"self_link"`
	ResourceTypeLink       string `json:"resource_type_link"`
	Description            string `json:"description"`
}
