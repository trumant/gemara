// Top level schema //

metadata?: #Metadata

controls?: [...#Control]
threats?: [...#Threat]
capabilities?: [...#Capability]

"shared-controls"?: [...#Mapping]
"shared-threats"?: [...#Threat]
"shared-capabilities"?: [...#Capability]

"mapping-references"?: [...#MappingReference]

// Resuable types //

#Metadata: {
    id: string
    title: string
    description: string
    version?: string
    "last-modified"?: string
}

#Control: {
    id: string
    title: string
    objective: string
    family: string
    "assessment-requirements": [...#Requirement]

    mappings?: [...#Mapping]
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
