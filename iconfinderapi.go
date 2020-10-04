package iconfinderapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Iconfinder is the object that contains all of the api interfacing commands
type Iconfinder struct {
	apikey string
}

type badrequest struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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
func (icofdr *Iconfinder) ChangeAPIKey(usrAPIKey string) {
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
func (icofdr *Iconfinder) GetUserIDDetails(id int32) (*User, error) {
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwusrinfo User
	json.Unmarshal(bodyBytes, &nwusrinfo)

	return &nwusrinfo, nil
}

// Style Object for containing a style and its details
type Style struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// GetStyleDetails returns a Style strut with information about the requested style
func (icofdr *Iconfinder) GetStyleDetails(styleID string) (*Style, error) {
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwstlinfo Style
	json.Unmarshal(bodyBytes, &nwstlinfo)

	return &nwstlinfo, nil
}

//Styles holds multiple style objects
type Styles struct {
	Total  int     `json:"total_count"`
	Styles []Style `json:"styles"`
}

// ListAllStyles returns a map of all the styles after can be initalized as "" if you dont want to use it.
// Count should be 0-100
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListAllStyles(Count int32, After string) (*Styles, error) {
	var reqstr string = apiuri + "styles"
	if Count == -1 || len(After) != 0 {
		reqstr += "?"
	}
	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if len(After) != 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + After
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwstlsinfo Styles
	json.Unmarshal(bodyBytes, &nwstlsinfo)

	return &nwstlsinfo, nil
}

// ListAllStylesFast returns all posible styles without any variables
func (icofdr *Iconfinder) ListAllStylesFast() *Styles {
	stl, _ := icofdr.ListAllStyles(100, "")
	return stl
}

// License details about the license
type License struct {
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	LicenseID int    `json:"license_id"`
	URL       string `json:"url"`
}

// GetLicenseDetails retuns details about the licence
func (icofdr *Iconfinder) GetLicenseDetails(licenseID int32) (*License, error) {
	req, err := http.NewRequest("GET", apiuri+"licenses/"+strconv.Itoa(int(licenseID)), nil)
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwlicinfo License
	json.Unmarshal(bodyBytes, &nwlicinfo)

	return &nwlicinfo, nil
}

// Category Object for containing a category and its details
type Category struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

// GetCategoryDetails Get details about a specific category identified by its identifier.
func (icofdr *Iconfinder) GetCategoryDetails(CategoryIdentifier string) (*Category, error) {
	req, err := http.NewRequest("GET", apiuri+"categories/"+CategoryIdentifier, nil)
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwcatinfo Category
	json.Unmarshal(bodyBytes, &nwcatinfo)

	return &nwcatinfo, nil
}

// Catagories holds multiple style objects
type Catagories struct {
	Total  int     `json:"total_count"`
	Styles []Style `json:"categories"`
}

// ListAllCategories lists all catagories.
// Count should be 0-100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListAllCategories(Count int32, After string) (*Catagories, error) {
	var reqstr string = apiuri + "categories"
	if Count != -1 || len(After) != 0 {
		reqstr += "?"
	}

	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if len(After) != 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + After
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwcatsinfo Catagories
	json.Unmarshal(bodyBytes, &nwcatsinfo)

	return &nwcatsinfo, nil
}

// ListAllCatagoriesFast returns all posible styles without any variables
func (icofdr *Iconfinder) ListAllCatagoriesFast() *Catagories {
	cats, _ := icofdr.ListAllCategories(-1, "")
	return cats
}

// Author contains details about an author
type Author struct {
	Name          string `json:"name"`
	IconsetsCount int    `json:"iconsets_count"`
	AuthorID      uint32 `json:"author_id"`
	WebsiteURL    string `json:"website_url"`
}

// GetAuthorDetails Get details about a specific author identified by a unique ID.
func (icofdr *Iconfinder) GetAuthorDetails(AuthorID int32) (*Author, error) {
	req, err := http.NewRequest("GET", apiuri+"authors/"+strconv.Itoa(int(AuthorID)), nil)
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwatorinfo Author
	json.Unmarshal(bodyBytes, &nwatorinfo)

	return &nwatorinfo, nil
}

// Price contains details about a price on an iconset
type Price struct {
	License  License `json:"license"`
	Currency string  `json:"currency"`
	Price    float32 `json:"price"`
}

// IconSet contains details about an icon
type IconSet struct {
	AllIconsGlyph bool       `json:"are_all_icons_glyph"`
	Styles        []Style    `json:"styles"`
	IsPremium     bool       `json:"is_Premium"`
	Identifier    string     `json:"identifier"`
	Prices        []Price    `json:"prices"`
	Author        User       `json:"author"`
	IconsetID     int        `json:"iconset_id"`
	Type          string     `json:"type"`
	PublishedAt   string     `json:"published_at"`
	WebsiteURL    string     `json:"website_url"`
	License       License    `json:"license"`
	Name          string     `json:"name"`
	ReadMe        string     `json:"readme"`
	IconsCount    int        `json:"icons_count"`
	Catagories    []Category `json:"categories"`
}

// GetIconSetDetails Get details about a specific icon set
func (icofdr *Iconfinder) GetIconSetDetails(IconsetID int32) (*IconSet, error) {
	req, err := http.NewRequest("GET", apiuri+"iconsets/"+strconv.Itoa(int(IconsetID)), nil)
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
	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicoset IconSet
	json.Unmarshal(bodyBytes, &nwicoset)

	return &nwicoset, nil
}

// IconSets contains multiple iconsets
type IconSets struct {
	TotalCount int       `json:"total_count"`
	IconSets   []IconSet `json:"iconsets"`
}

// ListIconSetsOfStyle returns IconSets by a specific Style
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListIconSetsOfStyle(StyleIdentifier string, Count int32, After int32, Premium int8, Vector int8, Licence string) (*IconSets, error) {
	var reqstr string = apiuri + "styles/" + StyleIdentifier + "/iconsets"

	// Check and apply filters to strings
	if Count != -1 || After != -1 || Premium != -1 || Vector != -1 || len(Licence) != 0 {
		reqstr += "?"
	}
	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if After != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + strconv.Itoa(int(After))
		prefix = true
	}

	if Premium != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "premium=" + strconv.Itoa(int(Premium))
		prefix = true
	}

	if Vector != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "vector=" + strconv.Itoa(int(Vector))
		prefix = true
	}

	if len(Licence) > 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "Licence=" + Licence
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicosets IconSets
	json.Unmarshal(bodyBytes, &nwicosets)

	return &nwicosets, nil
}

// ListIconSetsOfAuthor returns IconSets by a specific Author
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListIconSetsOfAuthor(AuthorID int32, Count int32, After int32, Premium int8, Vector int8, Licence string) (*IconSets, error) {
	var reqstr string = apiuri + "authors/" + strconv.Itoa(int(AuthorID)) + "/iconsets"

	// Check and apply filters to strings
	if Count != -1 || After != -1 || Premium != -1 || Vector != -1 || len(Licence) != 0 {
		reqstr += "?"
	}
	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if After != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + strconv.Itoa(int(After))
		prefix = true
	}

	if Premium != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "premium=" + strconv.Itoa(int(Premium))
		prefix = true
	}

	if Vector != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "vector=" + strconv.Itoa(int(Vector))
		prefix = true
	}

	if len(Licence) > 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "Licence=" + Licence
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicosets IconSets
	json.Unmarshal(bodyBytes, &nwicosets)

	return &nwicosets, nil
}

// ListIconSetsOfUser returns IconSets by a specific User
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListIconSetsOfUser(UserID string, Count int32, After int32, Premium int8, Vector int8, Licence string) (*IconSets, error) {
	var reqstr string = apiuri + "users/" + UserID + "/iconsets"

	// Check and apply filters to strings
	if Count != -1 || After != -1 || Premium != -1 || Vector != -1 || len(Licence) != 0 {
		reqstr += "?"
	}
	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if After != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + strconv.Itoa(int(After))
		prefix = true
	}

	if Premium != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "premium=" + strconv.Itoa(int(Premium))
		prefix = true
	}

	if Vector != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "vector=" + strconv.Itoa(int(Vector))
		prefix = true
	}

	if len(Licence) > 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "Licence=" + Licence
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicosets IconSets
	json.Unmarshal(bodyBytes, &nwicosets)

	return &nwicosets, nil
}

// ListIconSetsOfCategory returns IconSets by a specific User
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListIconSetsOfCategory(CatagoryID string, Count int32, After int32, Premium int8, Vector int8, Licence string) (*IconSets, error) {
	var reqstr string = apiuri + "categories/" + CatagoryID + "/iconsets"

	// Check and apply filters to strings
	if Count != -1 || After != -1 || Premium != -1 || Vector != -1 || len(Licence) != 0 {
		reqstr += "?"
	}

	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if After != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + strconv.Itoa(int(After))
		prefix = true
	}

	if Premium != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "premium=" + strconv.Itoa(int(Premium))
		prefix = true
	}

	if Vector != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "vector=" + strconv.Itoa(int(Vector))
		prefix = true
	}

	if len(Licence) > 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "Licence=" + Licence
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicosets IconSets
	json.Unmarshal(bodyBytes, &nwicosets)

	return &nwicosets, nil
}

// ListPublicIconSets List all public icon sets in descending order of when they were published.
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) ListPublicIconSets(Count int32, After int32, Premium int8, Vector int8, Licence string) (*IconSets, error) {
	var reqstr string = apiuri + "iconsets"

	// Check and apply filters to strings
	if Count != -1 || After != -1 || Premium != -1 || Vector != -1 || len(Licence) != 0 {
		reqstr += "?"
	}

	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if After != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "after=" + strconv.Itoa(int(After))
		prefix = true
	}

	if Premium != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "premium=" + strconv.Itoa(int(Premium))
		prefix = true
	}

	if Vector != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "vector=" + strconv.Itoa(int(Vector))
		prefix = true
	}

	if len(Licence) > 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "Licence=" + Licence
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicosets IconSets
	json.Unmarshal(bodyBytes, &nwicosets)

	return &nwicosets, nil
}

// Image Holds an images URL's and its format
type Image struct {
	PreviewURL  string `json:"preview_url"`
	Format      string `json:"format"`
	DownloadURL string `json:"download_url"`
}

// Sizes Holds all images formats and definse the size of the format
type Sizes struct {
	SizeHeight  int     `json:"size_height"`
	Formats     []Image `json:"formats"`
	Size        int     `json:"size"`
	SizeWidth   int     `json:"size_width"`
	TargetSizes []int   `json:"target_sizes"`
}

// Icon Holds details about an icon
type Icon struct {
	IconID      int        `json:"icon_id"`
	Tags        []string   `json:"tags"`
	Style       []Style    `json:"Styles"`
	Catagories  []Category `json:"categories"`
	IsPremium   bool       `json:"is_premium"`
	Prices      []Price    `json:"prices"`
	IsPurchased bool       `json:"is_purchased"`
	PublishedAt string     `json:"published_at"`
	Type        int        `json:"type"`
	IconSet     IconSet    `json:"iconset"`
	IconGlyph   bool       `json:"is_icon_glyph"`
	Containers  []Image    `json:"containers"`
	Vectors     []Sizes    `json:"vector_sizes"`
	Rasters     []Sizes    `json:"raster_sizes"`
}

// Icons holds multiple icon s
type Icons struct {
	TotalCount int    `json:"total_count"`
	Icons      []Icon `json:"icons"`
}

// IconsInSet lists all icons in a set
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) IconsInSet(IconsetID int32, Query string, Count int32, Offset int32, Vector int8) (*Icons, error) {
	var reqstr string = apiuri + "iconsets/" + strconv.Itoa(int(IconsetID)) + "/icons"

	// Check and apply filters to strings
	if Count != -1 || Offset != -1 || Vector != -1 || len(Query) != 0 {
		reqstr += "?"
	}

	prefix := false
	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "count=" + strconv.Itoa(int(Count))
		prefix = true
	}

	if Offset != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "offset=" + strconv.Itoa(int(Offset))
		prefix = true
	}

	if Vector != -1 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "vector=" + strconv.Itoa(int(Vector))
		prefix = true
	}

	if len(Query) > 0 {
		if prefix {
			reqstr += "&"
		}
		reqstr += "query=" + Query
	}
	fmt.Println(reqstr)
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicos Icons
	json.Unmarshal(bodyBytes, &nwicos)

	return &nwicos, nil
}

// GetIconDetails gets all details for a specific icon
func (icofdr *Iconfinder) GetIconDetails(IconID int32) (*Icon, error) {
	req, err := http.NewRequest("GET", apiuri+"icons/"+strconv.Itoa(int(IconID)), nil)
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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}
	var nwico Icon
	json.Unmarshal(bodyBytes, &nwico)

	return &nwico, nil
}

// SearchIcons Look for a icon
// Count is a range of 0 - 100 or -1
// Set unused filters to -1 for ints and "" for strings
func (icofdr *Iconfinder) SearchIcons(Query string, Count int32, Offset int32, Premium int8, Vector int8, Licence string, Category string, Style string) (*Icons, error) {
	var reqstr string = apiuri + "icons/search?query=" + Query

	if Count != -1 {
		if Count > 100 || Count < 0 {
			return nil, errors.New("OutOfBound")
		}
		reqstr += "&count=" + strconv.Itoa(int(Count))
	}

	if Offset != -1 {
		reqstr += "&offset=" + strconv.Itoa(int(Offset))
	}

	if Premium != -1 {
		reqstr += "&premium=" + strconv.Itoa(int(Premium))
	}

	if Vector != -1 {
		reqstr += "&vector=" + strconv.Itoa(int(Vector))
	}

	if len(Licence) > 0 {
		reqstr += "&licence=" + Licence
	}

	if len(Category) > 0 {
		reqstr += "&category=" + Category
	}

	if len(Style) > 0 {
		reqstr += "&style=" + Style
	}

	// fmt.Println(reqstr)

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

	var badreq badrequest
	json.Unmarshal(bodyBytes, &badreq)
	if badreq.Code == "not_found" {
		fmt.Println(badreq)
		return nil, errors.New("Bad Request")
	}

	var nwicos Icons
	json.Unmarshal(bodyBytes, &nwicos)

	return &nwicos, nil
}
