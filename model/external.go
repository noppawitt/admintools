package model

// ExternalStoredProcedure is a external stored procedure model
type ExternalStoredProcedure struct {
	Name string `json:"name"`
}

// ExternalView is a external view model
type ExternalView struct {
	Name string `json:"name"`
}

// ExternalParameter is a external parameter model
type ExternalParameter struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Length int    `json:"length"`
}
