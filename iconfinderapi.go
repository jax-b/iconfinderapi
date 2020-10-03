package iconfinderapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Iconfinder is the object that contains all of the api interfacing commands
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

// User Contains all of the information returned by UserIDDetails
type User struct {
	UserID          int32  `json:"user_id"`
	SocialTwitter   string `json:"social_twitter"`
	WebsiteURL      string `json:"website_url"`
	Company         string `json:"company"`
	Location        string `json:"location"`
	IsDesigner      bool   `json:"is_designer"`
	IconsetsCount   uint32 `json:"iconsets_count"`
	Name            string `json:"name"`
	SocialInstagram string `json:"social_instagram"`
	Username        string `json:"username"`
	SocialBehance   string `json:"social_behance"`
	SocialDribbble  string `json:"social_dribbble"`
}

// GetUserIDDetails returns information about the specifyed user
func (icofdr *Iconfinder) GetUserIDDetails(id int32) *User {
	req, err := http.NewRequest("GET", apiuri+"users/"+strconv.Itoa(int(id)), nil)
	req.Header.Add("Authorization", "Bearer "+icofdr.apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	var nwusrinfo User
	json.Unmarshal(bodyBytes, &nwusrinfo)

	return &nwusrinfo
}

// Style Object for containing a style and its details
type Style struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// GetStyleDetails returns a Style strut with information about the requested style
func (icofdr *Iconfinder) GetStyleDetails(styleID string) *Style {
	req, err := http.NewRequest("GET", apiuri+"styles/"+styleID, nil)
	req.Header.Add("Authorization", "Bearer "+icofdr.apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	var nwstlinfo Style
	json.Unmarshal(bodyBytes, &nwstlinfo)

	return &nwstlinfo
}

//Styles holds multiple style objects
type Styles struct {
	Total  int     `json:"total_count"`
	Styles []Style `json:"styles"`
}

// ListAllStyles returns a map of all the styles after can be initalized as "" if you dont want to use it
func (icofdr *Iconfinder) ListAllStyles(count int32, after string) *Styles {

	var reqstr string = apiuri + "styles?count=" + strconv.Itoa(int(count))

	if len(after) > 0 {
		reqstr = reqstr + "&after=" + after
	}

	req, err := http.NewRequest("GET", reqstr, nil)
	req.Header.Add("Authorization", "Bearer "+icofdr.apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	var nwstlsinfo Styles
	json.Unmarshal(bodyBytes, &nwstlsinfo)

	return &nwstlsinfo
}

// ListAllStylesFast returns all posible styles without any variables
func (icofdr *Iconfinder) ListAllStylesFast() *Styles {
	return icofdr.ListAllStyles(100, "")
}
