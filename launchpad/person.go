package launchpad

import (
	"log"
	"net/url"
)

type Person struct {
	LanguagesCollectionLink                 string `json:"languages_collection_link"`
	MembersCollectionLink                   string `json:"members_collection_link"`
	SubTeamsCollectionLink                  string `json:"sub_teams_collection_link"`
	DeactivatedMembersCollectionLink        string `json:"deactivated_members_collection_link"`
	PpasCollectionLink                      string `json:"ppas_collection_link"`
	Private                                 bool   `json:"private"`
	TimeZone                                string `json:"time_zone"`
	ProposedMembersCollectionLink           string `json:"proposed_members_collection_link"`
	MembershipsDetailsCollectionLink        string `json:"membership_details_collection_link"`
	AllSpecificationsCollectionLink         string `json:"all_specifications_collection_link"`
	AccountStatus                           string `json:"account_status"`
	IsProbationary                          bool   `json:"is_probationary"`
	TeamOwnerLink                           string `json:"team_owner_link"`
	MugshotLink                             string `json:"mugshot_link"`
	DisplayName                             string `json:"display_name"`
	GpgKeysCollectionLink                   string `json:"gpg_keys_collection_link"`
	LogoLink                                string `json:"logo_link"`
	OpenMembershipInvitationsCollectionLink string `json:"open_membership_invitations_collection_link"`
	IrcNicknamesCollectionLink              string `json:"irc_nicknames_collection_link"`
	HttpEtag                                string `json:"http_etag"`
	IsValid                                 bool   `json:"is_valid"`
	SelfLink                                string `json:"self_link"`
	ConfirmedEmailAddressesCollectionLink   string `json:"confirmed_email_address_collection_link"`
	Karma                                   uint   `json:"karma"`
	MailingListAutoSubscribePolicy          string `json:"mailing_list_auto_subscribe_policy"`
	Description                             string `json:"description"`
	MembersDetailsCollectionLink            string `json:"members_details_collection_link"`
	HideEmailAddresses                      bool   `json:"hide_email_addresses"`
	AdminsCollectionLink                    string `json:"admins_collection_link"`
	ValidSpecificationsCollectionLink       string `json:"valid_specifications_collection_link"`
	Visibility                              string `json:"visibility"`
	RecipesCollectionLink                   string `json:"recipes_collection_link"`
	DateCreated                             string `json:"date_created"`
	PreferredEmailAddressLink               string `json:"preffered_email_address_link"`
	IsUbuntuCocSigner                       bool   `json:"is_ubuntu_coc_signer"`
	InvitedMembersCollectionLink            string `json:"invited_members_collection_link"`
	ExpiredMembersCollectionLink            string `json:"expired_members_collection_link"`
	SshkeysCollectionLink                   string `json:"ssh_keys_collection_link"`
	Name                                    string `json:"name"`
	ResourceTypeLink                        string `json:"resource_type_link`
	SuperTeamsCollectionLink                string `json:"super_teams_collection_link"`
	ParticipantsCollectionLink              string `json:"participants_collection_link"`
	WebLink                                 string `json:"web_link"`
	HardwareSubmissionsCollectionLink       string `json:"hardware_submissions_collection_link"`
	ArchiveLink                             string `json:"archive_link"`
	IsTeam                                  bool   `json:"is_team"`
	AccountStatusHistory                    string `json:"account_status_history"`
	WikiNamesCollectionLink                 string `json:"wiki_names_collection_link"`
	HomepageContent                         string `json:"homepage_content"`
	JabberIdsCollectionLink                 string `json:"jabber_ids_collection_link"`

	lp *Launchpad
}

func (p *Person) SearchTasks() ([]BugTask, error) {
	v := url.Values{}
	v.Add("ws.op", "searchTasks")
	v.Add("order_by", "-date_last_updated")
	v.Add("start", "0")

	response, err := p.lp.Get(p.SelfLink, v)
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
