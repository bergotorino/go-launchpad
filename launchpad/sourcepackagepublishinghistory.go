package launchpad

import (
	"log"
	"net/url"
	"time"
)

type SrcPkgPubHistory struct {
	PackageCreatorLink    string      `json:"package_creator_link"`
	PackageSignerLink     string      `json:"package_signer_link"`
	SourcePackageName     string      `json:"source_package_name"`
	RemovalComment        interface{} `json:"removal_comment"`
	SponsorLink           interface{} `json:"sponsor_link"`
	DisplayName           string      `json:"display_name"`
	CreatorLink           interface{} `json:"creator_link"`
	SourcePackageVersion  string      `json:"source_package_version"`
	DateSuperseded        interface{} `json:"date_superseded"`
	HTTPEtag              string      `json:"http_etag"`
	PackageuploadLink     interface{} `json:"packageupload_link"`
	SelfLink              string      `json:"self_link"`
	DistroSeriesLink      string      `json:"distro_series_link"`
	ComponentName         string      `json:"component_name"`
	Status                string      `json:"status"`
	DateRemoved           interface{} `json:"date_removed"`
	Pocket                string      `json:"pocket"`
	DatePublished         time.Time   `json:"date_published"`
	RemovedByLink         interface{} `json:"removed_by_link"`
	SectionName           string      `json:"section_name"`
	DateMadePending       interface{} `json:"date_made_pending"`
	ResourceTypeLink      string      `json:"resource_type_link"`
	ArchiveLink           string      `json:"archive_link"`
	PackageMaintainerLink string      `json:"package_maintainer_link"`
	DateCreated           time.Time   `json:"date_created"`
	ScheduledDeletionDate interface{} `json:"scheduled_deletion_date"`

	lp *Launchpad
}

func (spph SrcPkgPubHistory) SourceFileUrls() ([]string, error) {
	v := url.Values{}
	v.Add("ws.op", "sourceFileUrls")

	response, err := spph.lp.Get(spph.SelfLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	var data []string

	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	return data, nil
}
