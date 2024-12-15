package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salaray  int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salaray <= 0 {
		return fmt.Errorf("request body is empty or malformmed")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Salaray <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

type UpadateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salaray  int64  `json:"salary"`
}

func (r *UpadateOpeningRequest) Validate() error {
	//if any field is provided, validation is truthy

	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salaray > 0 {
		return nil
	}

	// if none of the fields were provided
	return fmt.Errorf("at least one field must be provided")

}
