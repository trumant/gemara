package schemas

#Layer2: {
    metadata?: #Metadata
    control_families?: [...#ControlFamily] @go(ControlFamilies)
    threats?: [...#Threat]
    capabilities?: [...#Capability]

    "shared-controls"?: [...#Mapping] @go(-)
    "shared-threats"?: [...#Mapping] @go(-)
    "shared-capabilities"?: [...#Mapping] @go(-)
}

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
    requirements: [...#AssessmentRequirement]

    "guideline-mappings"?: [...#Mapping] @go(-)
    "threat-mappings"?: [...#Mapping] @go(-)
}

#Threat: {
    id: string
    title: string
    description: string
    capabilities: [...#Mapping]
    mappings?: [...#Mapping]
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
