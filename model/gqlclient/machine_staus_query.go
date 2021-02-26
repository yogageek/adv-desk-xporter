package model

import "github.com/shurcooL/graphql"

var QueryMachineStatuses struct { //如果不要都叫Query, 可自定義名稱包在外面，這裡大小寫沒差
	// MachineStatuses []struct { //這裡一定要大寫
	// 	// ~layer1
	// 	Name  graphql.String //裡面的欄位名稱一定要大寫開頭
	// 	Index interface{}
	// 	Depth graphql.Int
	// 	Names []struct {
	// 		Lang interface{}
	// 		Text interface{}
	// 	}
	// 	Color     interface{}
	// 	IsDefault interface{}
	// 	//Children //沒辦法這樣用,會跟下面重命衝突, 上面還是要寫

	// 	Children []Children //ok

	// } `graphql:"machineStatuses(layer1Only: $layer1Only)"` //傳參數

	//上面的struct可以用外嵌方式像這樣(為了合gql規格 所以MachineStatuses不用在加一次s)
	MachineStatuses []MachineStatuses `graphql:"machineStatuses(layer1Only: $layer1Only)"` //傳參數
}

type MachineStatuses struct { //這裡一定要大寫
	// ~layer1
	Id    graphql.String
	Name  graphql.String //裡面的欄位名稱一定要大寫開頭, 而且型態要正確!
	Index graphql.Int
	Depth graphql.Int
	Names []struct {
		Lang graphql.String
		Text graphql.String
	}
	Color     graphql.String
	IsDefault graphql.Boolean
	//Children //沒辦法這樣用,會跟下面重命衝突, 上面還是要寫

	Children []Children //ok
}

type Parent struct {
	Index int
}

//目前desk有規定最多就三層
type Children struct {
	Id     graphql.String
	Parent Parent
	// ~layer2
	ParentId graphql.String
	Index    graphql.Int
	Color    graphql.String
	Depth    graphql.Int
	Name     graphql.String //裡面的欄位名稱一定要大寫開頭
	Names    []struct {
		Lang graphql.String
		Text graphql.String
	}
	Children []Children2
}

type Children2 struct {
	Id     graphql.String
	Parent Parent
	// ~layer3
	ParentId graphql.String
	Index    graphql.Int
	Color    graphql.String
	Depth    graphql.Int
	Name     graphql.String //裡面的欄位名稱一定要大寫開頭
	Names    []struct {
		Lang graphql.String
		Text graphql.String
	}
}

//-----------------------

/*
//cch style (fail)

var MachineStatusHierarchy struct {
	MachineStatuses []struct {
		Children Childrens //not work
	}
}

type Childrens []Children

type Children struct {
	// # layer1
	Index interface{}
	Depth interface{}
	Name  interface{}
	Names []struct {
		Lang interface{}
		Text interface{}
	}
	Children Childrens
}
*/
