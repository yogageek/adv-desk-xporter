package logic

import (
	model "porter/model/gqlclient"
	logic "porter/pkg/logic/client"
	"time"

	// . "porter/util"
	util "porter/util"
)

func Run() {
	// logic.LoopRefreshToken()
	time.Sleep(1 * time.Second)

	logic.PrepareGQLClient()
	a := QueryGroups()
	util.PrintJson(a)
	// b := QueryMachineStatuses()
	// util.PrintJson(b)
	// c := QueryMachines()
	// util.PrintJson(c)

	input := model.TranslateGroupInput{
		Id:          "R3JvdXA.YEXwS8GunQAGpEa1",
		Name:        "TestTrans",
		Lang:        "en-US",
		Description: "kekekeke",
	}
	gqlQuery := model.TranslateGroup

	Mutate(input, &gqlQuery)
}
