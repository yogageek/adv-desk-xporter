package logic

import (
	"context"
	"encoding/json"
	model "porter/model/gqlclient"
	"porter/util"

	"github.com/golang/glog"
	"github.com/shurcooL/graphql"
)

func QueryTranslationLangs(gclient *graphql.Client) []model.TranslationLangs {
	gqlQuery := model.QueryTranslationLangs

	variables := map[string]interface{}{} //if no variables

	err := gclient.Query(context.Background(), &gqlQuery, variables)
	if err != nil {
		glog.Error(err)
	}

	//debugging
	b, _ := json.MarshalIndent(gqlQuery, "", " ")
	util.PrintGreen(b)

	return gqlQuery.TranslationLangs
}
