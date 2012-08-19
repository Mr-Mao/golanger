package models

import (
	"golanger/utils"
)

var (
	ColModule = utils.M{
		"name":  "module",
		"index": []string{"path", "name", "order", "status"},
	}
)

/*
模块表
module {
    "name" : <name>,
    "path" : <path>,
    "order" : <order>,
    "status" : <status>,
    "create_time" : <create_time>,
    "update_time" : <update_time>,
}
*/
type ModelModule struct {
	Name        string
	Path        string
	Order       string
	Status      byte
	Create_time int64
	Update_time int64
}
