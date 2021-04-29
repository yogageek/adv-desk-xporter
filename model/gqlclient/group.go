package model

//for query, no need json tag------
type Groups struct { //這裡一定要大寫
	ParentId    string      `json:"parentId,omitempty"`
	Id          string      `json:"id,omitempty"` //裡面的欄位名稱一定要大寫開頭, 而且型態要正確!
	Name        string      `json:"name"`
	Description string      `json:"description"`
	TimeZone    string      `json:"timeZone"`
	Coordinate  *Coordinate `json:"coordinate"`
	//新增多語言資料
	Names        []Name        `json:"names"`
	Descriptions []Description `json:"descriptions"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

var QueryGroups struct { //這行可隨便定義
	Groups []Groups //為graphql規格定義
}

//add後返回的查詢 注意和input是兩回事
var AddGroup struct { //這行可隨便定義 但盡量和下面同名
	AddGroup struct { //這裡只用來對應addGroup 沒其他作用
		Group Groups
	} `graphql:"addGroup(input: $input)"`
}

//for input,  need json tag------
type AddGroupInput struct { //參數
	Groups
}
