package model

//for query, no need json tag------
type TranslationLangs struct { //這裡一定要大寫
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Lang      string `json:"lang"`
	IsDefault bool   `json:"isDefault"`
}

var QueryTranslationLangs struct { //這行可隨便定義
	TranslationLangs []TranslationLangs //為graphql規格定義
}
