package iconfinderapi

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Iconfinder struct {
	apikey string
}

const apiuri = "https://api.iconfinder.com/v4/"

// NewIconFinder returns an iconfinder object
func NewIconFinder(usrAPIKey string) *Iconfinder {
	nwicofndr := &Iconfinder{
		apikey: usrAPIKey,
	}
	return nwicofndr
}

// ChangeAPIKey allows the user to change the api key after initilization
func (icofdr Iconfinder) ChangeAPIKey(usrAPIKey string) {
	icofdr.apikey = usrAPIKey
}

// UserInfo Contains all of the information returned by UserIDDetails
type UserInfo struct {
	UserID          int32  'json:"user_id"'
	SocialTwitter   string 'json:"social_twitter"'
	WebsiteURL      string 'json:"website_url"'
	Company         string 'json:"company"'
	Location        string 'json:"location"'
	IsDesigner      bool   'json:"is_designer"'
	IconsetsCount   uint32 'json:"iconsets_count"'
	Name            string 'json:"name"'
	SocialInstagram string 'json:"social_instagram"'
	Username        string 'json:"username"'
	SocialBehance   string 'json:"social_behance"'
	SocialDribbble  string 'json:"social_dribbble"'
}

// UserIDDetails returns information about the specifyed user
func (icofdr *Iconfinder) UserIDDetails(id int32) *UserInfo {
	resp, err := http.Get(apiuri + "user_id")
	resp.Header.Add("Bearer", apikey)
	if err != nil {
        log.Fatalln(err)
	}
	
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var nwusrinfo UserInfo
	json.Unmarshal(bodyBytes, &todoStruct)

	return nwusrinfo
}
