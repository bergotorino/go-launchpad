package launchpad

import (
	"log"
	"net/url"
	"time"
)

type Snap struct {
	StoreChannels                 []string    `json:"store_channels"`
	DateLastModified              time.Time   `json:"date_last_modified"`
	StoreUpload                   bool        `json:"store_upload"`
	Private                       bool        `json:"private"`
	WebLink                       string      `json:"web_link"`
	AutoBuild                     bool        `json:"auto_build"`
	CanUploadToStore              bool        `json:"can_upload_to_store"`
	StoreName                     string      `json:"store_name"`
	Name                          string      `json:"name"`
	DateCreated                   time.Time   `json:"date_created"`
	GitRefLink                    string      `json:"git_ref_link"`
	BranchLink                    interface{} `json:"branch_link"`
	GitRepositoryLink             string      `json:"git_repository_link"`
	RequireVirtualized            bool        `json:"require_virtualized"`
	PendingBuildsCollectionLink   string      `json:"pending_builds_collection_link"`
	AutoBuildArchiveLink          interface{} `json:"auto_build_archive_link"`
	HTTPEtag                      string      `json:"http_etag"`
	BuildsCollectionLink          string      `json:"builds_collection_link"`
	SelfLink                      string      `json:"self_link"`
	ResourceTypeLink              string      `json:"resource_type_link"`
	Description                   interface{} `json:"description"`
	GitPath                       string      `json:"git_path"`
	WebhooksCollectionLink        string      `json:"webhooks_collection_link"`
	StoreSeriesLink               string      `json:"store_series_link"`
	GitRepositoryURL              interface{} `json:"git_repository_url"`
	DistroSeriesLink              string      `json:"distro_series_link"`
	RegistrantLink                string      `json:"registrant_link"`
	OwnerLink                     string      `json:"owner_link"`
	CompletedBuildsCollectionLink string      `json:"completed_builds_collection_link"`
	AutoBuildPocket               interface{} `json:"auto_build_pocket"`
	ProcessorsCollectionLink      string      `json:"processors_collection_link"`

	lp *Launchpad
}

func (s Snap) String() string {
	return s.WebLink
}

func (s Snap) Processors() ([]Processor, error) {
	v := url.Values{}
	response, err := s.lp.Get(s.ProcessorsCollectionLink, v)
	if err != nil {
		log.Println("API error: ", err)
		return nil, err
	}

	data := struct {
		Entries          []Processor `json:"entries"`
		ResourceTypeLink string      `json:"resource_type_link"`
		Start            int         `json:"start"`
		TotalSize        int         `json:"total_size"`
	}{}

	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	return data.Entries, nil
}
