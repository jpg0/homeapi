package main

import (
	"fmt"
)

func InitSelectShows() {
	RegisterHandler("select_show", &SelectShowController{
	})
}

type SelectShowController struct {

}

func (ssc *SelectShowController) Run(req *APIAIRequest, config *Configuration) (*APIAIResponse, error) {
	idxStr := req.Result.Parameters["show_index"]

	//idx, err := strconv.Atoi(idxStr)
	//
	//if err != nil {
	//	//re-ask
	//	panic("no implemented")
	//}

	ctx := req.GetContext("show_options")

	if ctx == nil {
		//re-ask
		panic("no implemented")
	}

	showI, show_present := ctx.Parameters[idxStr]

	if !show_present {
		panic("no show")
	}

	show := showI.(map[string]interface{})

	newCtx := NewAPIAIContext("show", 5)
	newCtx.Parameters["showid"] = show["showid"]
	newCtx.Parameters["showtype"] = show["showtype"]
	newCtx.Parameters["showname"] = show["title"]

	res :=  NewAPIAIResponse(fmt.Sprintf("Confirm download %v?", show["title"]))

	res.AddContext(newCtx)

	return res, nil
}