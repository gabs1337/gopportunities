package handler

import "fmt"

func errParamIsRequerid(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s)is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request body os empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequerid("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequerid("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequerid("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequerid("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequerid("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequerid("salary", "int64")
	}
	return nil
}
