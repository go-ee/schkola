package finance

import (
    "encoding/json"
    "github.com/go-ee/utils/net"
    "io/ioutil"
    "net/http"
)
type ExpenseClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewExpenseClient(url string, client *http.Client) (ret *ExpenseClient) {
    url = url + "/" + "expenses"
    ret = &ExpenseClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *ExpenseClient) ImportJSON(fileJSON string) (err error) {
    var items []*Expense
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *ExpenseClient) Create(items []*Expense) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *ExpenseClient) ReadFileJSON(fileJSON string) (ret []*Expense, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type ExpensePurposeClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewExpensePurposeClient(url string, client *http.Client) (ret *ExpensePurposeClient) {
    url = url + "/" + "expensePurposes"
    ret = &ExpensePurposeClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *ExpensePurposeClient) ImportJSON(fileJSON string) (err error) {
    var items []*ExpensePurpose
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *ExpensePurposeClient) Create(items []*ExpensePurpose) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *ExpensePurposeClient) ReadFileJSON(fileJSON string) (ret []*ExpensePurpose, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type FeeClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewFeeClient(url string, client *http.Client) (ret *FeeClient) {
    url = url + "/" + "fees"
    ret = &FeeClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *FeeClient) ImportJSON(fileJSON string) (err error) {
    var items []*Fee
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *FeeClient) Create(items []*Fee) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *FeeClient) ReadFileJSON(fileJSON string) (ret []*Fee, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type FeeKindClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewFeeKindClient(url string, client *http.Client) (ret *FeeKindClient) {
    url = url + "/" + "feeKinds"
    ret = &FeeKindClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *FeeKindClient) ImportJSON(fileJSON string) (err error) {
    var items []*FeeKind
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *FeeKindClient) Create(items []*FeeKind) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *FeeKindClient) ReadFileJSON(fileJSON string) (ret []*FeeKind, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type FinanceClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
    ExpenseClient *ExpenseClient `json:"expenseClient" eh:"optional"`
    ExpensePurposeClient *ExpensePurposeClient `json:"expensePurposeClient" eh:"optional"`
    FeeClient *FeeClient `json:"feeClient" eh:"optional"`
    FeeKindClient *FeeKindClient `json:"feeKindClient" eh:"optional"`
}

func NewFinanceClient(url string, client *http.Client) (ret *FinanceClient) {
    url = url + "/" + "finance"
    expenseClient := NewExpenseClient(url, client)
    expensePurposeClient := NewExpensePurposeClient(url, client)
    feeClient := NewFeeClient(url, client)
    feeKindClient := NewFeeKindClient(url, client)
    ret = &FinanceClient{
        Url: url,
        Client: client,
        ExpenseClient: expenseClient,
        ExpensePurposeClient: expensePurposeClient,
        FeeClient: feeClient,
        FeeKindClient: feeKindClient,
    }
    return
}









