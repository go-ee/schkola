package person

import (
    "encoding/json"
    "github.com/go-ee/utils/net"
    "io/ioutil"
    "net/http"
)
type ChurchClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewChurchClient(url string, client *http.Client) (ret *ChurchClient) {
    url = url + "/" + "churches"
    ret = &ChurchClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *ChurchClient) ImportJSON(fileJSON string) (err error) {
    var items []*Church
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *ChurchClient) Create(items []*Church) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *ChurchClient) ReadFileJSON(fileJSON string) (ret []*Church, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type GraduationClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewGraduationClient(url string, client *http.Client) (ret *GraduationClient) {
    url = url + "/" + "graduations"
    ret = &GraduationClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *GraduationClient) ImportJSON(fileJSON string) (err error) {
    var items []*Graduation
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *GraduationClient) Create(items []*Graduation) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *GraduationClient) ReadFileJSON(fileJSON string) (ret []*Graduation, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type ProfileClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewProfileClient(url string, client *http.Client) (ret *ProfileClient) {
    url = url + "/" + "profiles"
    ret = &ProfileClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *ProfileClient) ImportJSON(fileJSON string) (err error) {
    var items []*Profile
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *ProfileClient) Create(items []*Profile) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *ProfileClient) ReadFileJSON(fileJSON string) (ret []*Profile, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type PersonClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
    ChurchClient *ChurchClient `json:"churchClient" eh:"optional"`
    GraduationClient *GraduationClient `json:"graduationClient" eh:"optional"`
    ProfileClient *ProfileClient `json:"profileClient" eh:"optional"`
}

func NewPersonClient(url string, client *http.Client) (ret *PersonClient) {
    url = url + "/" + "person"
    churchClient := NewChurchClient(url, client)
    graduationClient := NewGraduationClient(url, client)
    profileClient := NewProfileClient(url, client)
    ret = &PersonClient{
        Url: url,
        Client: client,
        ChurchClient: churchClient,
        GraduationClient: graduationClient,
        ProfileClient: profileClient,
    }
    return
}









