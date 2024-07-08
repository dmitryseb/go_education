package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

type PathAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
