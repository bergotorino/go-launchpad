package launchpad

import (
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

type Build struct {
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
}

func (b Build) BuildState() int {
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
