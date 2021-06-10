package model

//for query, no need json tag------
type Machines struct { //這裡一定要大寫
	Id          string `json:"id,omitempty"` //裡面的欄位名稱一定要大寫開頭, 而且型態要正確!
	GroupId     string `json:"groupId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	IsStation   *bool  `json:"isStation,omitempty"` //omitempty的代表參數不一定要放或不用放
	//新增多語言資料
	Names        []Name        `json:"names"`
	Descriptions []Description `json:"descriptions"`
	ImageUrls    []ImageUrl    `json:"imageUrls"`
}

var QueryMachines struct { //這行可隨便定義
	Machines []Machines //為graphql規格定義
}

//add後返回的查詢 注意和input是兩回事
var AddMachine struct { //這行可隨便定義 但盡量和下面同名
	AddMachine struct { //這裡只用來對應addMachine 沒其他作用
		Machine Machines //Machine這個名稱是對應到gql規格
	} `graphql:"addMachine(input: $input)"`
}

//for input,  need json tag------
type AddMachineInput struct { //參數
	GroupId     string `json:"groupId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	IsStation   *bool  `json:"isStation,omitempty"`
}
