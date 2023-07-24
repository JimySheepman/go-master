package main

import (
	"encoding/json"

	"github.com/golang/glog"
)

func ToJSON(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		glog.Warningf("ToJSON encode error: %+v data: %+v", err, v)
	}

	return string(data)
}
