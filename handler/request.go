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
		return errParamIsRequerid("location", "string")
	}
	if r.Link == "" {
		return fmt.Errorf("request body os empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequerid("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequerid("company", "string")
	}
	if r.Location == "" {
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

// UpdateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided , validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided in the request")
}
