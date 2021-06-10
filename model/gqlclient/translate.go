package model

//add後返回的查詢 注意和input是兩回事
var TranslateGroup struct { //這行可隨便定義 但盡量和下面同名
	TranslateGroup struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		Group struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateGroup(input: $input)"`
}

type TranslateGroupInput struct { //參數
	Id          string `json:"id"`
	Name        string `json:"name"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
}

//----------------------

var TranslateMachine struct { //這行可隨便定義 但盡量和下面同名
	TranslateMachine struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		Machine struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateMachine(input: $input)"`
}

type TranslateMachineInput struct { //參數
	Id          string `json:"id"`
	Name        string `json:"name"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

//----------------------

var TranslateMachineStatus struct { //這行可隨便定義 但盡量和下面同名
	TranslateMachineStatus struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		MachineStatus struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateMachineStatus(input: $input)"`
}

type TranslateMachineStatusInput struct { //參數
	Id   string `json:"id"`
	Lang string `json:"lang"`
	Name string `json:"name"`
}

//----------------------

var TranslateParameterMappingCode struct { //這行可隨便定義 但盡量和下面同名
	TranslateParameterMappingCode struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		MappingCode struct {
			Id string
		} //為graphql規格定義
	} `graphql:"translateParameterMappingCode(input: $input)"`
}

type TranslateParameterMappingCodeInput struct { //參數
	Id      string `json:"id"`
	Lang    string `json:"lang"`
	Message string `json:"message"`
}

//----------------------

var TranslateParameter struct { //這行可隨便定義 但盡量和下面同名
	TranslateParameter struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		Parameter struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateParameter(input: $input)"`
}

type TranslateParameterInput struct { //參數
	Id          string `json:"id"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
}

//----------------------

var TranslateProfileMachine struct { //這行可隨便定義 但盡量和下面同名
	TranslateProfileMachine struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		ProfileMachine struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateProfileMachine(input: $input)"`
}

// input
//使用上面 TranslateMachineInput

//----------------------

var TranslateProfileParameter struct { //這行可隨便定義 但盡量和下面同名
	TranslateProfileParameter struct { //這裡只用來對應addParameterMappingRule 沒其他作用
		ProfileParameter struct {
			Name string
		} //為graphql規格定義
	} `graphql:"translateProfileParameter(input: $input)"`
}

// input
//使用上面 TranslateParameterInput
