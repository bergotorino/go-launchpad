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

type Bug struct {
	UsersUnaffectedCollectionLink        string     `json:"users_unaffected_collection_link"`
	LatestPatchUploaded                  *time.Time `json:"latest_patch_uploaded"`
	UsersAffectedCountWithDupes          int        `json:"users_affected_count_with_dupes"`
	SecurityRelated                      bool       `json:"security_related"`
	Private                              bool       `json:"private"`
	BugWatchesCollectionLink             string     `json:"bug_watches_collection_link"`
	DateMadePrivate                      *time.Time `json:"date_made_private"`
	LinkedBranchesCollectionLink         string     `json:"linked_branches_collection_link"`
	SubscriptionsCollectionLink          string     `json:"subscriptions_collection_link"`
	NumberOfDuplicates                   int        `json:"number_of_duplicates"`
	ID                                   int        `json:"id"`
	UsersUnaffectedCount                 int        `json:"users_unaffected_count"`
	Title                                string     `json:"title"`
	OtherUsersAffectedCountWithDupes     int        `json:"other_users_affected_count_with_dupes"`
	Name                                 string     `json:"name"`
	HTTPEtag                             string     `json:"http_etag"`
	MessagesCollectionLink               string     `json:"messages_collection_link"`
	SelfLink                             string     `json:"self_link"`
	InformationType                      string     `json:"information_type"`
	WhoMadePrivateLink                   string     `json:"who_made_private_link"`
	AttachmentsCollectionLink            string     `json:"attachments_collection_link"`
	ResourceTypeLink                     string     `json:"resource_type_link"`
	ActivityCollectionLink               string     `json:"activity_collection_link"`
	DateLastUpdated                      time.Time  `json:"date_last_updated"`
	Description                          string     `json:"description"`
	DuplicatesCollectionLink             string     `json:"duplicates_collection_link"`
	Tags                                 []string   `json:"tags"`
	MessageCount                         int        `json:"message_count"`
	Heat                                 int        `json:"heat"`
	BugTasksCollectionLink               string     `json:"bug_tasks_collection_link"`
	DuplicateOfLink                      string     `json:"duplicate_of_link"`
	LinkedMergeProposalsCollectionLink   string     `json:"linked_merge_proposals_collection_link"`
	UsersAffectedWithDupesCollectionLink string     `json:"users_affected_with_dupes_collection_link"`
	CvesCollectionLink                   string     `json:"cves_collection_link"`
	WebLink                              string     `json:"web_link"`
	UsersAffectedCount                   int        `json:"users_affected_count"`
	OwnerLink                            string     `json:"owner_link"`
	DateCreated                          *time.Time `json:"date_created"`
	DateLastMessage                      *time.Time `json:"date_last_message"`
	UsersAffectedCollectionLink          string     `json:"users_affected_collection_link"`
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

	Core Bug
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
	for i, _ := range data.Entries {
		v := url.Values{}
		response, err := s.lp.Get(data.Entries[i].BugLink, v)
		if err != nil {
			log.Println("Failed to fetch bug info", err)
			return nil, err
		}
		err = DecodeResponse(response, &data.Entries[i].Core)
		if err != nil {
			log.Println("Decoding error: ", err)
			return nil, err
		}
	}

	return data.Entries, nil
}
