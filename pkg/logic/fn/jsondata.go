package logic

import (
	model "porter/model/gqlclient"
)

//匯出的結構

type JsonData struct {
	MachineStatusData []*MachineStatusData
	MappingRuleData   []*MappingRuleData
	ProfileData       []*ProfileData
	GroupData         []*GroupData
	MachineData       []*MachineData
	ParameterData     []*ParameterData
	TranslationLangs  []*TranslationLangs
}

//未來改成從model拿原始input
type MachineStatusData struct {
	Id          string
	NewId       string //改直接放Id就好(錯)
	ParentId    string
	ParentIndex int //注意這裡形態要對否則會拿到空值
	Name        string
	Names       []Name
	Index       int
	Color       string
	Depth       int
}
type Name struct {
	Text string
	Lang string
}

type MappingRuleData struct {
	Id     string
	NewId  string
	Name   string
	PType  string
	Detail []Detail
}

type Detail struct { //這裡的大小寫會影響到json大小寫
	Id       string
	Code     string
	Message  string
	StatusId string
	Lang     string
	Text     string
	Messages []Message //fix
}

//fix
type Message struct {
	Lang string
	Text string
}

// 匯入profile第二步時候 需要帶profileMachineId 但這個id會在匯入profile第一步的時候才會產生 因此要在第一步時抓下來
// 所以我們可以直接忽略export導出的machineId
type ProfileData struct {
	model.ProfileMachine
}

type GroupData struct {
	model.Groups
}

type MachineData struct {
	model.Machines
}

type ParameterData struct {
	model.QueryParametersOb
}

type TranslationLangs struct {
	model.TranslationLangs
}
