package schemas

#Layer2: {
    metadata?: #Metadata
    control_families?: [...#ControlFamily] @go(ControlFamilies)
    threats?: [...#Threat]
    capabilities?: [...#Capability]
    shared_controls?: [...#Mapping] @go(SharedControls)
    shared_threats?: [...#Mapping] @go(SharedThreats)
    shared_capabilities?: [...#Mapping] @go(SharedCapabilities)
}

#Metadata: {
    id: string
    title: string
    description: string
    version?: string
    last_modified?: string @go(LastModified)
    applicability_categories?: [...#Category] @go(ApplicabilityCategories)
    mapping_references?: [...#MappingReference] @go(MappingReferences)
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
    assessment_requirements: [...#AssessmentRequirement] @go(AssessmentRequirements)
    guideline_mappings?: [...#Mapping] @go(GuidelineMappings)
    threat_mappings?: [...#Mapping] @go(ThreatMappings)
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
    reference_id: string @go(ReferenceID)
    identifiers: [...string]
}

#AssessmentRequirement: {
    id: string
    text: string
    applicability: [...string]

    recommendation?: string
}
