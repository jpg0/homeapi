package main

func list(req *ActionRequest) (*ActionResponse, error) {


	//period :=


	return &ActionResponse{
		Message: "some stuff",
	}, nil
}


func InitListActions() {
	Register("list", list)
}