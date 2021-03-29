package logic

import (
	model "porter/model/gqlclient"

	"github.com/golang/glog"
)

var DefaultLang string

func getSourceTranslations() []model.TranslationLangs {
	return QueryTranslationLangs(gclientQ)
}

func getEnvTranslations() []model.TranslationLangs {
	return QueryTranslationLangs(gclientM)
}

func GetDefaultLangFromJson(data *jsonData) string {
	t := data.TranslationLangs
	for _, v := range t {
		if v.IsDefault {
			return v.Lang
		}
	}
	glog.Error("can't find default lang")
	return ""
}

func GetDefaultLangFromEnv() string {
	t := getEnvTranslations()
	for _, v := range t {
		if v.IsDefault {
			return v.Lang
		}
	}
	glog.Error("can't find default lang")
	return ""
}

func SetDefaultLang() {
	DefaultLang = GetDefaultLangFromEnv()
}
