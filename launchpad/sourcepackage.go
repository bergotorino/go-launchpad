package launchpad

import (
	"log"
	"net/url"
	"time"
)

type SourcePackage struct {
	BugReportedAcknowledgement string   `json:"bug_reported_acknowledgement"`
	DisplayName                string   `json:"display_name"`
	Name                       string   `json:"name"`
	Title                      string   `json:"title"`
	DistributionLink           string   `json:"distribution_link"`
	WebLink                    string   `json:"web_link"`
	OfficialBugTags            []string `json:"official_bug_tags"`
	HTTPEtag                   string   `json:"http_etag"`
	SelfLink                   string   `json:"self_link"`
	BugReportingGuidelines     string   `json:"bug_reporting_guidelines"`
	ResourceTypeLink           string   `json:"resource_type_link"`
	UpstreamProductLink        string   `json:"upstream_product_link"`

	lp *Launchpad
}

type BugTask struct {
	DateClosed                 *time.Time `json:"date_closed"`
	DateAssigned               *time.Time `json:"date_assigned"`
	Title                      string     `json:"title"`
	BugLink                    string     `json:"bug_link"`
	BugWatchLink               string     `json:"bug_watch_link"`
	MilestoneLink              string     `json:"milestone_link"`
	HTTPEtag                   string     `json:"http_etag"`
	DateLeftClosed             *time.Time `json:"date_left_closed"`
	DateFixCommitted           *time.Time `json:"date_fix_committed"`
	DateFixReleased            *time.Time `json:"date_fix_released"`
	DateInProgress             *time.Time `json:"date_in_progress"`
	ResourceTypeLink           string     `json:"resource_type_link"`
	Status                     string     `json:"status"`
	BugTargetName              string     `json:"bug_target_name"`
	Importance                 string     `json:"importance"`
	AssigneeLink               string     `json:"assignee_link"`
	DateTriaged                *time.Time `json:"date_triaged"`
	SelfLink                   string     `json:"self_link"`
	TargetLink                 string     `json:"target_link"`
	BugTargetDisplayName       string     `json:"bug_target_display_name"`
	RelatedTasksCollectionLink string     `json:"related_tasks_collection_link"`
	DateConfirmed              *time.Time `json:"date_confirmed"`
	DateLeftNew                *time.Time `json:"date_left_new"`
	WebLink                    string     `json:"web_link"`
	OwnerLink                  string     `json:"owner_link"`
	DateCreated                *time.Time `json:"date_created"`
	DateIncomplete             *time.Time `json:"date_incomplete"`
	IsComplete                 bool       `json:"is_complete"`
}

func (s *SourcePackage) SearchBugs() ([]BugTask, error) {
	v := url.Values{}
	v.Add("ws.op", "searchTasks")
	v.Add("order_by", "-date_last_updated")
	v.Add("start", "0")

	response, err := s.lp.Get(s.SelfLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	data := struct {
		Entries            []BugTask `json:"entries"`
		Start              int       `json:"start"`
		TotalSizeLink      string    `json:"total_size_link"`
		NextCollectionLink string    `json:"next_collection_link"`
	}{}

	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	return data.Entries, nil
}
