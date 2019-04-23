package student

import (
    "encoding/json"
    "github.com/go-ee/utils/net"
    "io/ioutil"
    "net/http"
)
type AttendanceClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewAttendanceClient(url string, client *http.Client) (ret *AttendanceClient) {
    url = url + "/" + "attendances"
    ret = &AttendanceClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *AttendanceClient) ImportJSON(fileJSON string) (err error) {
    var items []*Attendance
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *AttendanceClient) Create(items []*Attendance) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *AttendanceClient) ReadFileJSON(fileJSON string) (ret []*Attendance, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type CourseClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewCourseClient(url string, client *http.Client) (ret *CourseClient) {
    url = url + "/" + "courses"
    ret = &CourseClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *CourseClient) ImportJSON(fileJSON string) (err error) {
    var items []*Course
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *CourseClient) Create(items []*Course) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *CourseClient) ReadFileJSON(fileJSON string) (ret []*Course, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type GradeClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewGradeClient(url string, client *http.Client) (ret *GradeClient) {
    url = url + "/" + "grades"
    ret = &GradeClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *GradeClient) ImportJSON(fileJSON string) (err error) {
    var items []*Grade
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *GradeClient) Create(items []*Grade) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *GradeClient) ReadFileJSON(fileJSON string) (ret []*Grade, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type GroupClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewGroupClient(url string, client *http.Client) (ret *GroupClient) {
    url = url + "/" + "groups"
    ret = &GroupClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *GroupClient) ImportJSON(fileJSON string) (err error) {
    var items []*Group
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *GroupClient) Create(items []*Group) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *GroupClient) ReadFileJSON(fileJSON string) (ret []*Group, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type SchoolApplicationClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewSchoolApplicationClient(url string, client *http.Client) (ret *SchoolApplicationClient) {
    url = url + "/" + "schoolApplications"
    ret = &SchoolApplicationClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *SchoolApplicationClient) ImportJSON(fileJSON string) (err error) {
    var items []*SchoolApplication
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *SchoolApplicationClient) Create(items []*SchoolApplication) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *SchoolApplicationClient) ReadFileJSON(fileJSON string) (ret []*SchoolApplication, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type SchoolYearClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
}

func NewSchoolYearClient(url string, client *http.Client) (ret *SchoolYearClient) {
    url = url + "/" + "schoolYears"
    ret = &SchoolYearClient{
        Url: url,
        Client: client,
    }
    return
}

func (o *SchoolYearClient) ImportJSON(fileJSON string) (err error) {
    var items []*SchoolYear
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
    return
}

func (o *SchoolYearClient) Create(items []*SchoolYear) (err error) {
    for _, item := range items {
		net.PostById(item, item.Id, o.Url, o.Client)
	}
    return
}

func (o *SchoolYearClient) ReadFileJSON(fileJSON string) (ret []*SchoolYear, err error) {
    jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
    return
}


type StudentClient struct {
    Url string `json:"url" eh:"optional"`
    Client *http.Client `json:"client" eh:"optional"`
    AttendanceClient *AttendanceClient `json:"attendanceClient" eh:"optional"`
    CourseClient *CourseClient `json:"courseClient" eh:"optional"`
    GradeClient *GradeClient `json:"gradeClient" eh:"optional"`
    GroupClient *GroupClient `json:"groupClient" eh:"optional"`
    SchoolApplicationClient *SchoolApplicationClient `json:"schoolApplicationClient" eh:"optional"`
    SchoolYearClient *SchoolYearClient `json:"schoolYearClient" eh:"optional"`
}

func NewStudentClient(url string, client *http.Client) (ret *StudentClient) {
    url = url + "/" + "student"
    attendanceClient := NewAttendanceClient(url, client)
    courseClient := NewCourseClient(url, client)
    gradeClient := NewGradeClient(url, client)
    groupClient := NewGroupClient(url, client)
    schoolApplicationClient := NewSchoolApplicationClient(url, client)
    schoolYearClient := NewSchoolYearClient(url, client)
    ret = &StudentClient{
        Url: url,
        Client: client,
        AttendanceClient: attendanceClient,
        CourseClient: courseClient,
        GradeClient: gradeClient,
        GroupClient: groupClient,
        SchoolApplicationClient: schoolApplicationClient,
        SchoolYearClient: schoolYearClient,
    }
    return
}









