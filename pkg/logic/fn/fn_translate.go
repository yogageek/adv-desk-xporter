package logic

import (
	model "porter/model/gqlclient"
	. "porter/pkg/logic/client"
	. "porter/pkg/logic/gql"

	"github.com/golang/glog"
)

func getEnvTranslations() []model.TranslationLangs {
	return QueryTranslationLangs(GclientM)
}

func GetDefaultLangFromJson(data *JsonData) string {
	t := data.TranslationLangsData
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
