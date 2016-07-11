package timing_structures

import (
	"time"
)

type Template struct {
	Id           string `json:"id"`
	TemplateId   string `json:"template_id"`
	Active       int    `json:"active"`
	Name         string `json:"name"`
	HtmlContent  string `json:"html_content"`
	PlainContent string `json:"plain_content"`
	Subject      string `json:"subject"`
	UpdatedAt    string `json:"updated_at"`
}

type Link struct {
	LinkId       string `json:"linkId"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	Auth         string `json:"auth"`
	AuthCookie   string `json:"authCookie"`
	AuthPassword string `json:"authPassword"`
	AuthUser     string `json:"authUser"`
}

type Queue struct {
	Id             string    `gorethink:"id,omitempty"`
	GroupId        string    `gorethink:"groupId"`
	UserId         string    `gorethink:"userId"`
	Status         string    `gorethink:"status"`
	CreateDateTime time.Time `gorethink:"createDateTime"`
	Email          string    `gorethink:"email"`
}

type Group struct {
	Id            string      `gorethink:"id,omitempty"`
	UserId        string      `gorethink:"userId"`
	Name          string      `gorethink:"name"`
	Email         string      `gorethink:"email"`
	IsEmail       bool        `gorethink:"isEmail"`
	Repeat        int         `gorethink:"repeat"`
	IsRepeat      bool        `gorethink:"isRepeat"`
	LastDateCheck interface{} `gorethink:"lastDateCheck"`
	Auth          string      `gorethink:"auth"`
	AuthCookie    string      `gorethink:"authCookie"`
	AuthPassword  string      `gorethink:"authPassword"`
	AuthUser      string      `gorethink:"authUser"`
}

type Timing struct {
	TimingId     string  `json:"timingId"`
	Url          string  `json:"url"`
	MinTime      float64 `json:"minTime"`
	MaxTime      float64 `json:"maxTime"`
	AvgTime      float64 `json:"avgTime"`
	Status       string  `json:"status"`
	Auth         string  `json:"auth"`
	AuthCookie   string  `json:"authCookie"`
	AuthPassword string  `json:"authPassword"`
	AuthUser     string  `json:"authUser"`
	NodeName     string  `json:"nodeName"`
}

type Email struct {
	Id             string `json:"id"`
	Status         string `json:"status"`
	Type           string `json:"type"`
	To             string `json:"to"`
	ReportUrl      string `json:"reportUrl"`
	UserName       string `json:"userName"`
	ContactName    string `json:"contactName"`
	ContactEmail   string `json:"contactEmail"`
	ContactMessage string `json:"contactMessage"`
}

type Config struct {
	NodeName  string
	RethinkDB struct {
		Server string
		DB     string
	}
	RabbitMQ struct {
		Server string
	}
	SendGrid struct {
		ApiKey string
		Host   string
	}
}
