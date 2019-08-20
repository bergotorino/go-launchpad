package launchpad

import (
	"log"
	"net/url"
	"time"
)

const (
	BuildFailure = iota
	BuildSuccess
	BuildNeedsBuilding
	BuildDependencyWait
	BuildChrootProblem
	BuildBuildSupersededSource
	BuildBuilding
	BuildUploadFailed
	BuildUploading
	BuildCancelling
	BuildCancelled
	BuildUnknown = 0xFF
)

type SnapBuild struct {
	CanBeRescored           bool        `json:"can_be_rescored"`
	BuilderLink             string      `json:"builder_link"`
	Datebuilt               time.Time   `json:"datebuilt"`
	DistroArchSeriesLink    string      `json:"distro_arch_series_link"`
	SnapLink                string      `json:"snap_link"`
	Duration                string      `json:"duration"`
	CanBeCancelled          bool        `json:"can_be_cancelled"`
	Title                   string      `json:"title"`
	Buildstate              string      `json:"buildstate"`
	RequesterLink           string      `json:"requester_link"`
	HTTPEtag                string      `json:"http_etag"`
	Score                   interface{} `json:"score"`
	SelfLink                string      `json:"self_link"`
	DateStarted             time.Time   `json:"date_started"`
	ResourceTypeLink        string      `json:"resource_type_link"`
	BuildLogURL             string      `json:"build_log_url"`
	StoreUploadURL          string      `json:"store_upload_url"`
	StoreUploadErrorMessage interface{} `json:"store_upload_error_message"`
	Pocket                  string      `json:"pocket"`
	Dependencies            interface{} `json:"dependencies"`
	DateFirstDispatched     time.Time   `json:"date_first_dispatched"`
	StoreUploadStatus       string      `json:"store_upload_status"`
	DistributionLink        string      `json:"distribution_link"`
	DistroSeriesLink        string      `json:"distro_series_link"`
	WebLink                 string      `json:"web_link"`
	Datecreated             time.Time   `json:"datecreated"`
	ArchiveLink             string      `json:"archive_link"`
	StoreUploadRevision     int         `json:"store_upload_revision"`
	ArchTag                 string      `json:"arch_tag"`
	UploadLogURL            interface{} `json:"upload_log_url"`

	lp *Launchpad
}

func (b SnapBuild) BuildState() int {
	switch b.Buildstate {
	case "Failed to build":
		return BuildFailure
	case "Successfully built":
		return BuildSuccess
	case "Needs building":
		return BuildNeedsBuilding
	case "Dependency wait":
		return BuildDependencyWait
	case "Chroot problem":
		return BuildChrootProblem
	case "Build for superseded Source":
		return BuildBuildSupersededSource
	case "Currently building":
		return BuildBuilding
	case "Failed to upload":
		return BuildUploadFailed
	case "Uploading build":
		return BuildUploading
	case "Cancelling build":
		return BuildCancelling
	case "Cancelled build":
		return BuildCancelled
	}
	return BuildUnknown
}

func (b SnapBuild) GetFileUrls() ([]string, error) {
	v := url.Values{}
	v.Add("ws.op", "getFileUrls")

	response, err := b.lp.Get(b.SelfLink, v)
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

type Build struct {
	ArchTag                      string      `json:"arch_tag"`
	ArchiveLink                  string      `json:"archive_link"`
	BuildLogURL                  string      `json:"build_log_url"`
	BuilderLink                  string      `json:"builder_link"`
	Buildstate                   string      `json:"buildstate"`
	CanBeCancelled               bool        `json:"can_be_cancelled"`
	CanBeRescored                bool        `json:"can_be_rescored"`
	CanBeRetried                 bool        `json:"can_be_retried"`
	ChangesfileURL               string      `json:"changesfile_url"`
	CurrentSourcePublicationLink interface{} `json:"current_source_publication_link"`
	DateFirstDispatched          time.Time   `json:"date_first_dispatched"`
	DateStarted                  time.Time   `json:"date_started"`
	Datebuilt                    time.Time   `json:"datebuilt"`
	Datecreated                  time.Time   `json:"datecreated"`
	Dependencies                 interface{} `json:"dependencies"`
	DistributionLink             string      `json:"distribution_link"`
	Duration                     string      `json:"duration"`
	ExternalDependencies         interface{} `json:"external_dependencies"`
	HTTPEtag                     string      `json:"http_etag"`
	Pocket                       string      `json:"pocket"`
	ResourceTypeLink             string      `json:"resource_type_link"`
	Score                        interface{} `json:"score"`
	SelfLink                     string      `json:"self_link"`
	SourcePackageName            string      `json:"source_package_name"`
	Title                        string      `json:"title"`
	UploadLogURL                 interface{} `json:"upload_log_url"`
	WebLink                      string      `json:"web_link"`

	lp *Launchpad
}

func (b Build) GetLatestSourcePublication() ([]string, error) {
	v := url.Values{}
	v.Add("ws.op", "getLatestSourcePublication")

	response, err := b.lp.Get(b.SelfLink, v)
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
