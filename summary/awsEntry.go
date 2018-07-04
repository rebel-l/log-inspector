package summary

import "strconv"

type AwsEntry struct {
	Uri string
	Count int
	StatusCodes []int
	UserAgents []string
}

func newAwsEntry(uri string, statusCode int, userAgent string) *AwsEntry {
	a := new(AwsEntry)
	a.Uri = uri
	a.Count = 1
	a.StatusCodes = make([]int, 1)
	a.StatusCodes[0] = statusCode
	a.UserAgents = make([]string, 1)
	a.UserAgents[0] = userAgent
	return a
}

func (a *AwsEntry) GetUserAgents() string {
	ua := ""
	for k, v := range a.UserAgents {
		if k > 0 {
			ua += ", "
		}
		ua += v
	}
	return ua
}

func (a *AwsEntry) AddUserAgent(userAgent string) {
	for _, v := range a.UserAgents {
		if userAgent == v {
			return
		}
	}
	a.UserAgents = append(a.UserAgents, userAgent)
}

func (a *AwsEntry) GetStatusCodes() string {
	codes := ""
	for k, v := range a.StatusCodes {
		if k > 0 {
			codes += ", "
		}
		codes += strconv.Itoa(v)
	}
	return codes
}

func (a *AwsEntry) AddStatusCode(statusCode int) {
	for _, v := range a.StatusCodes {
		if statusCode == v {
			return
		}
	}
	a.StatusCodes = append(a.StatusCodes, statusCode)
}