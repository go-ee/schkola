package library

import (
    "encoding/json"
    "github.com/go-ee/utils/net"
    "io/ioutil"
    "net/http"
)
type BookClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewBookClient(url string, client *http.Client) (ret *BookClient) {
    url = url + "/" + "books"
    ret = &BookClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *BookClient) ImportJSON(fileJSON string) (err error) {
    var items []*Book
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *BookClient) Create(items []*Book) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *BookClient) ReadFileJSON(fileJSON string) (ret []*Book, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type LibraryClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
    BookClient *BookClient `json:"bookClient" eh:"optional"`
}

func NewLibraryClient(url string, client *http.Client) (ret *LibraryClient) {
    url = url + "/" + "library"
    bookClient := NewBookClient(url, client)
    ret = &LibraryClient{
        Url: url,
        Client: client,
        BookClient: bookClient,
    }
    return
}









