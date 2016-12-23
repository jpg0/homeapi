package main

func list(req *ActionRequest) (*ActionResponse, error) {

	period := "7 days"
	data := "some stuff"

	return &ActionResponse{
		Context: map[string]string{
			"period": period,
			"data": data,
		},
	}, nil
}


func InitListActions() {
	Register("list", list)
}