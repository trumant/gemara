package main

type Layer2 struct {
	// All are optional, as these may be used to compile into another complete Layer2 object
	Metadata           Metadata           `json:"metadata" cue:"#Metadata"`
	Controls           []Control          `json:"controls" cue:"[...]#Control"`
	Threats            []Threat           `json:"threats" cue:"[...]#Threat"`
	Capabilities       []Capability       `json:"capabilities" cue:"[...]#Capability"`
	SharedControls     []Mapping          `json:"shared-controls" cue:"[...]#Mapping"`
	SharedThreats      []Threat           `json:"shared-threats" cue:"[...]#Threat"`
	SharedCapabilities []Capability       `json:"shared-capabilities" cue:"[...]#Capability"`
	MappingReferences  []MappingReference `json:"mapping-references" cue:"[...]#MappingReference"`
}

type Metadata struct {
	ID           string `json:"id" cue:"string"`
	Title        string `json:"title" cue:"string"`
	Description  string `json:"description" cue:"string"`
	Version      string `json:"version" cue:"string"`
	LastModified string `json:"last-modified" cue:"string"`
}

type Control struct {
	ID                     string        `json:"id" cue:"string"`
	Title                  string        `json:"title" cue:"string"`
	Objective              string        `json:"objective" cue:"string"`
	Family                 string        `json:"family" cue:"string"`
	AssessmentRequirements []Requirement `json:"assessment-requirements" cue:"[...]#Requirement"`

	// optional
	Category string    `json:"category" cue:"string"`
	Mappings []Mapping `json:"mappings" cue:"[...]#Mapping"`
}

type Threat struct {
	ID           string    `json:"id" cue:"string"`
	Title        string    `json:"title" cue:"string"`
	Description  string    `json:"description" cue:"string"`
	Capabilities []Mapping `json:"capabilities" cue:"[...]#Mapping"`

	// optional
	Category string    `json:"category" cue:"string"`
	Mappings []Mapping `json:"mappings" cue:"[...]#Mapping"`
}

type Capability struct {
	ID          string `json:"id" cue:"string"`
	Title       string `json:"title" cue:"string"`
	Description string `json:"description" cue:"string"`

	// optional
	Category string `json:"category" cue:"string"`
}

type MappingReference struct {
	ID      string `json:"id" cue:"string"`
	Title   string `json:"title" cue:"string"`
	Version string `json:"version" cue:"string"`

	// optional
	URL string `json:"url" cue:"=~\"^https?://[^\\s]+$\""`
}

type Mapping struct {
	ReferenceID string   `json:"reference-id" cue:"string"`
	Identifiers []string `json:"identifiers" cue:"[...]string"`
}

type Requirement struct {
	ID            string   `json:"id" cue:"string"`
	Text          string   `json:"text" cue:"string"`
	Applicability []string `json:"applicability" cue:"[...]string"`

	// optional
	Recommendation string `json:"recommendation" cue:"string"`
}
