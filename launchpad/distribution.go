package launchpad

import (
	"log"
	"net/url"
)

type Distribution struct {
	Active                            bool        `json:"active"`
	ActiveMilestonesCollectionLink    string      `json:"active_milestones_collection_link"`
	AllMilestonesCollectionLink       string      `json:"all_milestones_collection_link"`
	AllSpecificationsCollectionLink   string      `json:"all_specifications_collection_link"`
	ArchiveMirrorsCollectionLink      string      `json:"archive_mirrors_collection_link"`
	ArchivesCollectionLink            string      `json:"archives_collection_link"`
	BranchSharingPolicy               string      `json:"branch_sharing_policy"`
	BugReportedAcknowledgement        string      `json:"bug_reported_acknowledgement"`
	BugReportingGuidelines            string      `json:"bug_reporting_guidelines"`
	BugSharingPolicy                  string      `json:"bug_sharing_policy"`
	BugSupervisorLink                 string      `json:"bug_supervisor_link"`
	CdimageMirrorsCollectionLink      string      `json:"cdimage_mirrors_collection_link"`
	CurrentSeriesLink                 string      `json:"current_series_link"`
	DateCreated                       string      `json:"date_created"`
	DerivativesCollectionLink         string      `json:"derivatives_collection_link"`
	Description                       string      `json:"description"`
	DevelopmentSeriesAlias            string      `json:"development_series_alias"`
	DisplayName                       string      `json:"display_name"`
	DomainName                        string      `json:"domain_name"`
	DriverLink                        string      `json:"driver_link"`
	HomepageContent                   interface{} `json:"homepage_content"`
	HTTPEtag                          string      `json:"http_etag"`
	IconLink                          string      `json:"icon_link"`
	LogoLink                          string      `json:"logo_link"`
	MainArchiveLink                   string      `json:"main_archive_link"`
	MembersLink                       string      `json:"members_link"`
	MirrorAdminLink                   string      `json:"mirror_admin_link"`
	MugshotLink                       string      `json:"mugshot_link"`
	Name                              string      `json:"name"`
	OfficialBugTags                   []string    `json:"official_bug_tags"`
	OfficialPackages                  bool        `json:"official_packages"`
	OwnerLink                         string      `json:"owner_link"`
	RedirectReleaseUploads            bool        `json:"redirect_release_uploads"`
	RegistrantLink                    string      `json:"registrant_link"`
	ResourceTypeLink                  string      `json:"resource_type_link"`
	SelfLink                          string      `json:"self_link"`
	SeriesCollectionLink              string      `json:"series_collection_link"`
	SpecificationSharingPolicy        string      `json:"specification_sharing_policy"`
	Summary                           string      `json:"summary"`
	SupportsMirrors                   bool        `json:"supports_mirrors"`
	SupportsPpas                      bool        `json:"supports_ppas"`
	Title                             string      `json:"title"`
	TranslationgroupLink              string      `json:"translationgroup_link"`
	Translationpermission             string      `json:"translationpermission"`
	TranslationsUsage                 string      `json:"translations_usage"`
	ValidSpecificationsCollectionLink string      `json:"valid_specifications_collection_link"`
	Vcs                               interface{} `json:"vcs"`
	WebLink                           string      `json:"web_link"`

	lp *Launchpad
}

func (d *Distribution) GetSourcePackage(name string) (*SourcePackage, error) {
	v := url.Values{}
	v.Add("name", name)
	v.Add("ws.op", "getSourcePackage")

	response, err := d.lp.Get(d.SelfLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	var data SourcePackage
	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding went bad")
		return nil, err
	}

	data.lp = d.lp

	return &data, nil
}

func (d *Distribution) SearchTasks() ([]BugTask, error) {
	v := url.Values{}
	v.Add("ws.op", "searchTasks")
	v.Add("order_by", "-date_last_updated")
	v.Add("start", "0")

	response, err := d.lp.Get(d.SelfLink, v)
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
		response, err := d.lp.Get(data.Entries[i].BugLink, v)
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
