package responseTypes

import "time"

type Response struct {
	Message   string     `json:"message"`
	Data      *UserData  `json:"data"`
	Time      time.Time  `json:"time"`
	AccessKey *AccessKey `json:"access_key"`
}

func NewResponse(message string, data *UserData, access_key *AccessKey) *Response {
	return &Response{
		Message:   message,
		Data:      data,
		Time:      time.Now().UTC(),
		AccessKey: access_key,
	}
}
