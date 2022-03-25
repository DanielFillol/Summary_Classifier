package Struct

type Infered_decision struct {
	Summary    string `json:"summary,omitempty"`
	Text       string `json:"text,omitempty"`
	Class      string `json:"classificcao,omitempty"`
	Identifier string `json:"Identifier,omitempty"`
	Court      string `json:"Court,omitempty"`
}

type Raw_decision struct {
	Summary    string `json:"summary,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Court      string `json:"court,omitempty"`
}
