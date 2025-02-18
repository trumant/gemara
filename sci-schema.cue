import (
  "time"
)

#Control: {
    id: string
    title: string
    objective: string
    family: string
    threats: [...string]
    mappings: [...#Mapping]
    "assessment-requirements": [...#Requirement]
}

#Mapping: {
    framework: string
    version: string
    controls: [...string]
}

#Requirement: {
    id: string
    text: string
    "recommended-implementation"?: string
    applicability: [...string]
}

"common-controls"?: [...#Mapping]

"controls"?: [...#Control]
