package launchpad

import (
	"log"
	"net/url"
	"time"
)

type GitRepositories struct {
	lp *Launchpad
}

func (gr *GitRepositories) GetByPath(path string) (*GitRepository, error) {
	v := url.Values{}
	v.Add("path", path)
	v.Add("ws.op", "getByPath")

	response, err := gr.lp.Get("https://api.launchpad.net/devel/+git", v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	var data GitRepository
	err = gr.lp.DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding went bad")
		return nil, err
	}

	data.lp = gr.lp

	return &data, nil
}

type GitRepository struct {
	UniqueName                      string      `json:"unique_name"`
	DateLastModified                time.Time   `json:"date_last_modified"`
	Private                         bool        `json:"private"`
	RegistrantLink                  string      `json:"registrant_link"`
	SubscriptionsCollectionLink     string      `json:"subscriptions_collection_link"`
	OwnerDefault                    bool        `json:"owner_default"`
	TargetDefault                   bool        `json:"target_default"`
	RepositoryType                  string      `json:"repository_type"`
	GitIdentity                     string      `json:"git_identity"`
	DisplayName                     string      `json:"display_name"`
	HTTPEtag                        string      `json:"http_etag"`
	BranchesCollectionLink          string      `json:"branches_collection_link"`
	ReviewerLink                    interface{} `json:"reviewer_link"`
	RefsCollectionLink              string      `json:"refs_collection_link"`
	SelfLink                        string      `json:"self_link"`
	InformationType                 string      `json:"information_type"`
	ResourceTypeLink                string      `json:"resource_type_link"`
	GitSSHURL                       string      `json:"git_ssh_url"`
	Description                     interface{} `json:"description"`
	LandingTargetsCollectionLink    string      `json:"landing_targets_collection_link"`
	SubscribersCollectionLink       string      `json:"subscribers_collection_link"`
	DefaultBranch                   string      `json:"default_branch"`
	RecipesCollectionLink           string      `json:"recipes_collection_link"`
	WebhooksCollectionLink          string      `json:"webhooks_collection_link"`
	TargetLink                      string      `json:"target_link"`
	LandingCandidatesCollectionLink string      `json:"landing_candidates_collection_link"`
	Name                            string      `json:"name"`
	CodeImportLink                  interface{} `json:"code_import_link"`
	WebLink                         string      `json:"web_link"`
	DependentLandingsCollectionLink string      `json:"dependent_landings_collection_link"`
	OwnerLink                       string      `json:"owner_link"`
	GitHTTPSURL                     string      `json:"git_https_url"`
	DateCreated                     time.Time   `json:"date_created"`

	lp *Launchpad
}

type MergeProposal struct {
	Address                       string      `json:"address"`
	AllCommentsCollectionLink     string      `json:"all_comments_collection_link"`
	BugsCollectionLink            string      `json:"bugs_collection_link"`
	CommitMessage                 string      `json:"commit_message"`
	DateCreated                   string      `json:"date_created"`
	DateMerged                    string      `json:"date_merged"`
	DateReviewRequested           string      `json:"date_review_requested"`
	DateReviewed                  string      `json:"date_reviewed"`
	Description                   string      `json:"description"`
	HTTPEtag                      string      `json:"http_etag"`
	MergeReporterLink             string      `json:"merge_reporter_link"`
	MergedRevisionID              string      `json:"merged_revision_id"`
	MergedRevno                   interface{} `json:"merged_revno"`
	PrerequisiteBranchLink        interface{} `json:"prerequisite_branch_link"`
	PrerequisiteGitPath           interface{} `json:"prerequisite_git_path"`
	PrerequisiteGitRepositoryLink interface{} `json:"prerequisite_git_repository_link"`
	PreviewDiffLink               string      `json:"preview_diff_link"`
	PreviewDiffsCollectionLink    string      `json:"preview_diffs_collection_link"`
	Private                       bool        `json:"private"`
	QueueStatus                   string      `json:"queue_status"`
	RegistrantLink                string      `json:"registrant_link"`
	ResourceTypeLink              string      `json:"resource_type_link"`
	ReviewedRevid                 string      `json:"reviewed_revid"`
	ReviewerLink                  string      `json:"reviewer_link"`
	SelfLink                      string      `json:"self_link"`
	SourceBranchLink              interface{} `json:"source_branch_link"`
	SourceGitPath                 string      `json:"source_git_path"`
	SourceGitRepositoryLink       string      `json:"source_git_repository_link"`
	SupersededByLink              interface{} `json:"superseded_by_link"`
	SupersedesLink                interface{} `json:"supersedes_link"`
	TargetBranchLink              interface{} `json:"target_branch_link"`
	TargetGitPath                 string      `json:"target_git_path"`
	TargetGitRepositoryLink       string      `json:"target_git_repository_link"`
	VotesCollectionLink           string      `json:"votes_collection_link"`
	WebLink                       string      `json:"web_link"`
}

func (gr *GitRepository) LandingTargets() ([]MergeProposal, error) {
	v := url.Values{}

	response, err := gr.lp.Get(gr.LandingTargetsCollectionLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	data := struct {
		Entries          []MergeProposal `json:"entries"`
		ResourceTypeLink string          `json:"resource_type_link"`
		Start            int             `json:"start"`
		TotalSize        int             `json:"total_size"`
	}{}

	err = gr.lp.DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	return data.Entries, nil
}

func (gr *GitRepository) LandingCandidates() ([]MergeProposal, error) {
	v := url.Values{}

	response, err := gr.lp.Get(gr.LandingCandidatesCollectionLink, v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	data := struct {
		Entries          []MergeProposal `json:"entries"`
		ResourceTypeLink string          `json:"resource_type_link"`
		Start            int             `json:"start"`
		TotalSize        int             `json:"total_size"`
	}{}

	err = gr.lp.DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding error: ", err)
		return nil, err
	}

	return data.Entries, nil
}
