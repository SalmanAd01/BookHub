package models

type Book struct {
	Authorname     string `json:"authorname"`
	Subjectname    string `json:"subjectname"`
	Semnumber      string `json:"semnumber"`
	Branch         string `json:"branch"`
	Universityname string `json:"universityname"`
	Bookfile       string `json:"bookfile"`
}
