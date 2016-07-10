package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

//json resume is based on https://jsonresume.org/
type ResumeJson struct {
	Basics       BasicsInfo
	Work         []WorkInfo
	Volunteer    []VolunteerInfo
	Education    []EducationInfo
	Awards       []AwardsInfo
	Publications []PublicationsInfo
	Skills       []SkillsInfo
	Languages    []LaunguagesInfo
	Interests    []InterestsInfo
	References   []ReferencesInfo
}

type BasicsInfo struct {
	Name     string
	Label    string
	Picture  string
	Email    string
	Phone    string
	Website  string
	Summary  string
	Location LocationInfo
	Profiles []ProfilesInfo
}

type LocationInfo struct {
	Address     string
	PostalCode  string
	City        string
	CountryCode string
	Region      string
}

type ProfilesInfo struct {
	Network  string
	Username string
	Url      string
}

type WorkInfo struct {
	Company    string
	Position   string
	Website    string
	StartDate  string
	EndDate    string
	Summary    string
	Highlights HighlightsInfo
}

type HighlightsInfo struct {
	Highlight string
}

type VolunteerInfo struct {
	Organization string
	Position     string
	Website      string
	StartDate    string
	EndDate      string
	Summary      string
	Highlights   []HighlightsInfo
}

type EducationInfo struct {
	Institution string
	Area        string
	StudyType   string
	StartDate   string
	EndDate     string
	Gpa         float32
	Courses     []CoursesInfo
}

type CoursesInfo struct {
	Course string
}

type AwardsInfo struct {
	Title   string
	Date    string
	Awarded string
	Summary string
}

type PublicationsInfo struct {
	Name        string
	Publisher   string
	ReleaseDate string
	Website     string
	Summary     string
}

type SkillsInfo struct {
	Name     string
	Keywords string
}

type LaunguagesInfo struct {
	Name  string
	Level string
}

type InterestsInfo struct {
	Name     string
	Keywords string
}

type ReferencesInfo struct {
	Name      string
	Reference string
}

func New() ResumeJson {

	response, err := http.Get("http://golang.org")
	if err != nil {
		panic(err)
	}
	log.Info(response)
	var resume ResumeJson
	return resume
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
