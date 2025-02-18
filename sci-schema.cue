#Control: {
    id: string
    title: string
    objective: string
    family: string
    threats: [...string]
    "assessment-requirements": [...#Requirement]

    mappings?: [...#Mapping]
}

#Mapping: {
    framework: string
    version: string
    "control-ids": [...string]
}

#Requirement: {
    id: string
    text: string
    applicability: [...string]
    "recommended-implementation"?: string
}

"common-controls"?: [...#Mapping]

"controls"?: [...#Control]
