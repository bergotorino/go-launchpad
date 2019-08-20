package launchpad

import (
	"log"
	"net/url"
)

type Archive struct {
	ExternalDependencies                      interface{} `json:"external_dependencies"`
	Reference                                 string      `json:"reference"`
	BuildDebugSymbols                         bool        `json:"build_debug_symbols"`
	Private                                   bool        `json:"private"`
	ProcessorsCollectionLink                  string      `json:"processors_collection_link"`
	RequireVirtualized                        bool        `json:"require_virtualized"`
	SuppressSubscriptionNotifications         bool        `json:"suppress_subscription_notifications"`
	EnabledRestrictedProcessorsCollectionLink string      `json:"enabled_restricted_processors_collection_link"`
	PublishDebugSymbols                       bool        `json:"publish_debug_symbols"`
	HTTPEtag                                  string      `json:"http_etag"`
	SelfLink                                  string      `json:"self_link"`
	ResourceTypeLink                          string      `json:"resource_type_link"`
	Status                                    string      `json:"status"`
	DependenciesCollectionLink                string      `json:"dependencies_collection_link"`
	AuthorizedSize                            interface{} `json:"authorized_size"`
	Displayname                               string      `json:"displayname"`
	Description                               string      `json:"description"`
	PermitObsoleteSeriesUploads               bool        `json:"permit_obsolete_series_uploads"`
	RelativeBuildScore                        int         `json:"relative_build_score"`
	Name                                      string      `json:"name"`
	DistributionLink                          string      `json:"distribution_link"`
	WebLink                                   string      `json:"web_link"`
	OwnerLink                                 string      `json:"owner_link"`
	SigningKeyFingerprint                     string      `json:"signing_key_fingerprint"`

	lp *Launchpad
}

func (a Archive) GetPublishedSources(distroSeries string, sourceName string) ([]SrcPkgPubHistory, error) {
	v := url.Values{}
	v.Add("ws.op", "getPublishedSources")
	v.Add("distro_series", distroSeries)
	v.Add("source_name", sourceName)

	response, err := a.lp.Get(a.SelfLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	data := struct {
		Entries            []SrcPkgPubHistory `json:"entries"`
		Start              int                `json:"start"`
		TotalSizeLink      string             `json:"total_size_link"`
		NextCollectionLink string             `json:"next_collection_link"`
	}{}

	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	for i, _ := range data.Entries {
		data.Entries[i].lp = a.lp
	}

	return data.Entries, nil
}

func (a Archive) GetPublishedBinaries(distroArchSeries string, binaryName string, version string) ([]BinPkgPubHistory, error) {
	v := url.Values{}
	v.Add("ws.op", "getPublishedBinaries")
	v.Add("distro_arch_series", distroArchSeries)
	v.Add("binary_name", binaryName)
	v.Add("status", "Published")
	v.Add("exact_match", "true")
	v.Add("order_by_date", "true")
	if version != "" {
		v.Add("version", version)
	}

	response, err := a.lp.Get(a.SelfLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	data := struct {
		Entries            []BinPkgPubHistory `json:"entries"`
		Start              int                `json:"start"`
		TotalSizeLink      string             `json:"total_size_link"`
		NextCollectionLink string             `json:"next_collection_link"`
	}{}

	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	for i, _ := range data.Entries {
		data.Entries[i].lp = a.lp
	}

	return data.Entries, nil
}
