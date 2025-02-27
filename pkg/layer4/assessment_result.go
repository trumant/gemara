package layer4

// TestResult is a struct that contains the results of a single step within a testSet
type AssessmentResult struct {
	Passed      bool               // Passed is true if the test passed
	Description string             // Description is a human-readable description of the test
	Message     string             // Message is a human-readable description of the test result
	Function    string             // Function is the name of the code that was executed
	Value       interface{}        // Value is the object that was returned during the test
	Changes     map[string]*Change // Changes is a slice of changes that were made during the test
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

func (t *AssessmentResult) NewChange(changeName string, targetName string, targetObject *interface{}, applyFunc ApplyFunc, revertFunc RevertFunc) *Change {
	t.Changes[changeName] = &Change{
		Target_Name:   targetName,
		Target_Object: targetObject,
		applyFunc:     applyFunc,
		revertFunc:    revertFunc,
	}

	return t.Changes[changeName]
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
