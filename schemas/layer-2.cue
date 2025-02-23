// Top level schema //

metadata?: #Metadata

"control-families"?: [...#ControlFamily]
threats?: [...#Threat]
capabilities?: [...#Capability]

"shared-controls"?: [...#Mapping]
"shared-threats"?: [...#Threat]
"shared-capabilities"?: [...#Capability]

"ruleset-mappings"?: [...#MappingReference]

// Resuable types //

#Metadata: {
    id: string
    title: string
    description: string
    version?: string
    "last-modified"?: string
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
    requirements: [...#Requirement]

    "rule-mappings"?: [...#Mapping]
    "threat-mappings"?: [...#Mapping]
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
    url?: =~"^https?://[^\\s]+$"
}

#Mapping: {
    "reference-id": string
    identifiers: [...string]
}

#Requirement: {
    id: string
    text: string
    applicability: [...string]

    recommendation?: string
}
