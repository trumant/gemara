package layer4

// TestResult is a struct that contains the results of a single step within a testSet
type AssessmentResult struct {
	Passed      bool               `json:"passed"`      // Passed is true if the test passed
	Description string             `json:"description"` // Description is a human-readable description of the test
	Message     string             `json:"message"`     // Message is a human-readable description of the test result
	Function    string             `json:"function"`    // Function is the name of the code that was executed
	Value       interface{}        `json:"value"`       // Value is the object that was returned during the test
	Changes     map[string]*Change `json:"changes"`     // Changes is a slice of changes that were made during the test
}

func (t *AssessmentResult) SetPass(message string, value interface{}) {
	t.Passed = true
	t.Message = message
	t.Value = value
}

func (t *AssessmentResult) SetFail(message string, value interface{}) {
	t.Passed = false
	t.Message = message
	t.Value = value
}

func (t *AssessmentResult) SetReview(message string, value interface{}) {
	t.Passed = false
	t.Message = message
	t.Value = value
}

func (t *AssessmentResult) RevertChanges() (badStateAlert bool) {
	for _, change := range t.Changes {
		if !badStateAlert && (change.Applied || change.Error != nil) {
			if !change.Reverted {
				change.Revert()
			}
			if change.Error != nil || !change.Reverted {
				badStateAlert = true // do not break loop here; continue reverting all changes
			}
		}
	}
	return
}
