package model

type Name struct { //這裡一定要大寫
	Text string `json:"text,omitempty"`
	Lang string `json:"lang,omitempty"`
}

type Description struct { //這裡一定要大寫
	Text string `json:"text,omitempty"`
	Lang string `json:"lang,omitempty"`
}
