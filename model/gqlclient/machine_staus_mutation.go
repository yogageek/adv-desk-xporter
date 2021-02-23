package model

import "github.com/shurcooL/graphql"

var AddMachineStatus struct {
	AddMachineStatus struct {
		//裡面放甚麼 當api成功後就會返回給你你要的資料

		// MachineStatus struct {
		// 	Id    graphql.String
		// 	Name  graphql.String
		// 	Index graphql.Int
		// 	Color graphql.String
		// }
		MachineStatus MachineStatus //如果要把上面的struct外嵌 需這樣寫

		//也可以這樣寫
		// MachineStatus struct {
		// 	MachineStatus
		// }

	} `graphql:"addMachineStatus(input: $input)"`
}

type MachineStatus struct {
	Id    graphql.String //`json:"id"`
	Name  graphql.String //`json:"name"`
	Index graphql.Int    //`json:"index"`
	Color graphql.String //`json:"color"`
}

type AddMachineStatusInput struct { //參數
	//要塞的參數 大小寫要注意
	ParentId interface{} `json:"parentId"` //tag重要! 攸關轉gql後的大小寫
	Name     interface{} `json:"name"`
	Index    interface{} `json:"index"`
	Color    interface{} `json:"color"`
}

var UpdateMachineStatus struct {
	UpdateMachineStatus struct {
		MachineStatus struct {
			ID    graphql.String
			Name  graphql.String
			Color graphql.String
		}
	} `graphql:"updateMachineStatus(input: $input)"`
}

type UpdateMachineStatusInput struct {
	Id    interface{} `json:"id"`
	Name  interface{} `json:"name"` //tag重要! 攸關轉gql後的大小寫
	Color interface{} `json:"color"`
}

//---

// var Query struct {
// 	Groups []struct {
// 		Name graphql.String
// 	}
// }
