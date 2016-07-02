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
	LinkId string `json:"linkId"`
	Url    string `json:"url"`
	Title  string `json:"title"`
}

type Queue struct {
	Id             string    `gorethink:"id,omitempty"`
	GroupId        string    `gorethink:"groupId"`
	UserId         string    `gorethink:"userId"`
	Status         string    `gorethink:"status"`
	CreateDateTime time.Time `gorethink:"createDateTime"`
	Email          string    `gorethink:"status"`
}

type Timing struct {
	TimingId string  `json:"timingId"`
	Url      string  `json:"url"`
	MinTime  float64 `json:"minTime"`
	MaxTime  float64 `json:"maxTime"`
	AvgTime  float64 `json:"avgTime"`
	Status   string  `json:"status"`
}

type Email struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	To        string `json:"to"`
	ReportUrl string `json:"reportUrl"`
	UserName  string `json:"userName"`
}

type Config struct {
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
