package rpc

import "net/http"

var TaskClient *http.Client

func init() {
	InitClient()
}

//InitClient func for Init client
func InitClient() {
	TaskClient = &http.Client{}
}
