package launchpad

import (
	"log"
	"net/url"
	"time"
)

type BinPkgPubHistory struct {
	DistroArchSeriesLink   string      `json:"distro_arch_series_link"`
	PhasedUpdatePercentage interface{} `json:"phased_update_percentage"`
	BuildLink              string      `json:"build_link"`
	RemovalComment         interface{} `json:"removal_comment"`
	DisplayName            string      `json:"display_name"`
	DateMadePending        interface{} `json:"date_made_pending"`
	DateSuperseded         interface{} `json:"date_superseded"`
	PriorityName           string      `json:"priority_name"`
	HTTPEtag               string      `json:"http_etag"`
	SelfLink               string      `json:"self_link"`
	BinaryPackageVersion   string      `json:"binary_package_version"`
	ResourceTypeLink       string      `json:"resource_type_link"`
	ComponentName          string      `json:"component_name"`
	Status                 string      `json:"status"`
	DateRemoved            interface{} `json:"date_removed"`
	Pocket                 string      `json:"pocket"`
	DatePublished          time.Time   `json:"date_published"`
	RemovedByLink          interface{} `json:"removed_by_link"`
	SectionName            string      `json:"section_name"`
	ArchitectureSpecific   bool        `json:"architecture_specific"`
	BinaryPackageName      string      `json:"binary_package_name"`
	IsDebug                bool        `json:"is_debug"`
	ArchiveLink            string      `json:"archive_link"`
	DateCreated            time.Time   `json:"date_created"`
	ScheduledDeletionDate  interface{} `json:"scheduled_deletion_date"`

	lp *Launchpad
}

func (bpph BinPkgPubHistory) BinaryFileUrls() ([]string, error) {
	v := url.Values{}
	v.Add("ws.op", "binaryFileUrls")

	response, err := bpph.lp.Get(bpph.SelfLink, v)
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
