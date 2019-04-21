package person

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-ee/utils/net"
)

type ChurchImporter struct {
	client *http.Client
	url    string
}

func NewChurchImporter(url string) *ChurchImporter {
	return &ChurchImporter{client: &http.Client{},
		url: fmt.Sprintf("%v/%v", url, "person/churches")}
}

func (o *ChurchImporter) ImportJSON(fileJSON string) (err error) {
	var churches []*Church
	if churches, err = ReadChurchesFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(churches)
	return
}

func (o *ChurchImporter) Create(churches []*Church) (err error) {
	for _, church := range churches {
		net.PostById(church, church.Id, o.url, o.client)
	}
	return
}

func ReadChurchesFileJSON(file string) (churches []*Church, err error) {
	jsonBytes, _ := ioutil.ReadFile(file)

	if err = json.Unmarshal(jsonBytes, &churches); err != nil {
		log.Fatal("Cannot unmarshal", err)
	}
	return
}
