package model

type LogLine struct {
	Time int64 `json:"time"`
	Server string `json:"server"`
	File string `json:"file"`
	Line string `json:"line"`
}
