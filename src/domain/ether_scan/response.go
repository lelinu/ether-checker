package ether_scan

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (r *Response) IsError() bool{
	if r.Status == "0" {
		return true
	}

	return false
}
