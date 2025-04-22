package schemas

#Layer2: {
    metadata?: #Metadata
    // TODO: replace multiple `@go(-)` when https://github.com/cue-lang/cue/commit/93c1421c23ac8d5ddc8910a9186f5b94e5252ea9 releases in cue > v0.12.1
    "control-families"?: [...#ControlFamily] @go(-)
    threats?: [...#Threat]
    capabilities?: [...#Capability]

    "shared-controls"?: [...#Mapping] @go(-)
    "shared-threats"?: [...#Mapping] @go(-)
    "shared-capabilities"?: [...#Mapping] @go(-)
}

// Resuable types //

#Metadata: {
    id: string
    title: string
    description: string
    version?: string
    "last-modified"?: string @go(-)
    "applicability-categories"?: [...#Category] @go(-)
    "mapping-references"?: [...#MappingReference] @go(-)
}

#Category: {
    id: string
    title: string
    description: string
}

#ControlFamily: {
    title: string
    description: string
    controls: [...#Control]
}

#Control: {
    id: string
    title: string
    objective: string
    "assessment-requirements": [...#AssessmentRequirement] @go(-)

    "guideline-mappings"?: [...#Mapping] @go(-)
    "threat-mappings"?: [...#Mapping] @go(-)
}

#Threat: {
    id: string
    title: string
    description: string
    capabilities: [...#Mapping]

    "external-mappings"?: [...#Mapping] @go(-)
}

#Capability: {
    id: string
    title: string
    description: string
}

#MappingReference: {
    id: string
    title: string
    version: string
    description?: string
    url?: =~"^https?://[^\\s]+$"
}

#Mapping: {
    "reference-id": string @go(-)
    identifiers: [...string]
}

#AssessmentRequirement: {
    id: string
    text: string
    applicability: [...string]

    recommendation?: string
}
