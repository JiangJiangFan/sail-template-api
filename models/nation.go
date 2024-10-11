package models

type Nation struct {
	// 民族ID
	NationID int `json:"nation_id"`
	// 民族名称
	NationName string `json:"nation_name"`
	// 民族代码
	NationCode string `json:"nation_code"`
}