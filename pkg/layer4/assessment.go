package layer4

// TestResult is a struct that contains the results of a single step within a testSet
type Assessment struct {
	Requirement_Id string             // Requirement_ID is the unique identifier for the requirement being tested
	Description    string             // Description is a human-readable description of the test
	Result         Result             // Passed is true if the test passed
	Message        string             // Message is the human-readable result of the test
	Steps          []AssessmentStep   // Steps is a slice of steps that were executed during the test
	StepsExecuted  int                // StepsExecuted is the number of steps that were executed during the test
	Value          interface{}        // Value is the object that was returned during the test
	Changes        map[string]*Change // Changes is a slice of changes that were made during the test
}

// AssessmentStep is a function type that inspects the provided targetData and returns a Result with a message
// The message may be an error string or other descriptive text
type AssessmentStep func(targetData interface{}, a *Assessment) (Result, string)

// NewStep queues a new step in the Assessment
func (t *Assessment) NewStep(step AssessmentStep) {
	t.Steps = append(t.Steps, step)
}

func (a *Assessment) runStep(targetData interface{}, step AssessmentStep) Result {
	a.StepsExecuted++
	result, message := step(targetData, a)
	a.Result = checkResultOverride(a.Result, result)
	a.Message = message
	return result
}

// Run will execute all steps, halting if any step does not return layer4.Passed
func (a *Assessment) Run(targetData interface{}) Result {
	for _, step := range a.Steps {
		if a.runStep(targetData, step) == Failed {
			return Failed
		}
	}
	return a.Result
}

// RunTolerateFailures will execute all steps, halting only if a step
// returns an unknown result
func (a *Assessment) RunTolerateFailures(targetData interface{}) Result {
	for _, step := range a.Steps {
		a.runStep(targetData, step)
	}
	return a.Result
}

// NewChange creates a new Change object and adds it to the Assessment
func (t *Assessment) NewChange(changeName string, targetName string, targetObject *interface{}, applyFunc ApplyFunc, revertFunc RevertFunc) *Change {
	if t.Changes == nil {
		t.Changes = make(map[string]*Change)
	}
	t.Changes[changeName] = &Change{
		Target_Name:   targetName,
		Target_Object: targetObject,
		applyFunc:     applyFunc,
		revertFunc:    revertFunc,
	}

	return t.Changes[changeName]
}

func (t *Assessment) RevertChanges() (corrupted bool) {
	for _, change := range t.Changes {
		if !corrupted && (change.Applied || change.Error != nil) {
			if !change.Reverted {
				change.Revert()
			}
			if change.Error != nil || !change.Reverted {
				corrupted = true // do not break loop here; continue attempting to revert all changes
			}
		}
	}
	return
}
