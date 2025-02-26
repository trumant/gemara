// Top level schema //

"evaluation-plans": [...#EvaluationPlan]

// Types

#EvaluationPlan {
    name: string
    "start-time": string
    "end-time": string
    result: #Result
    "corrupted-state"?: bool
    "evaluation-results": [...#Evaluation]
}

#Evaluation: {
    name: string
    "control-id": string
    result: #Result
    message: string
    "documentation-url"?: =~"^https?://[^\\s]+$"
    "corrupted-state"?: bool
    "test-results"?: [...#TestResult]
}

#AssessmentResult: {
    result: #Result
    name: string
    description: string
    message: string
    "function-address" string
}

#Result: "Passed" | "Failed" | "Needs Review"