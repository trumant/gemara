package layer4

import (
	"testing"
)

var assessmentsTestData = []struct {
	testName           string
	assessment         Assessment
	numberOfSteps      int
	numberOfStepsToRun int
	expectedResult     Result
}{
	{
		testName:   "Assessment with no steps",
		assessment: Assessment{},
	},
	{
		testName:           "Assessment with one step",
		assessment:         passingAssessment,
		numberOfSteps:      1,
		numberOfStepsToRun: 1,
		expectedResult:     Passed,
	},
	{
		testName:           "Assessment with two steps",
		assessment:         failingAssessment,
		numberOfSteps:      2,
		numberOfStepsToRun: 1,
		expectedResult:     Failed,
	},
	{
		testName:           "Assessment with three steps",
		assessment:         needsReviewAssessment,
		numberOfSteps:      3,
		numberOfStepsToRun: 3,
		expectedResult:     NeedsReview,
	},
	{
		testName:           "Assessment with four steps",
		assessment:         badRevertPassingAssessment,
		numberOfSteps:      4,
		numberOfStepsToRun: 4,
		expectedResult:     Passed,
	},
}

// TestNewStep ensures that NewStep queues a new step in the Assessment
func TestAddStep(t *testing.T) {
	for _, test := range assessmentsTestData {
		t.Run(test.testName, func(t *testing.T) {
			if len(test.assessment.Steps) != test.numberOfSteps {
				t.Errorf("Bad test data: expected to start with %d, got %d", test.numberOfSteps, len(test.assessment.Steps))
			}
			test.assessment.AddStep(passingAssessmentStep)
			if len(test.assessment.Steps) != test.numberOfSteps+1 {
				t.Errorf("expected %d, got %d", test.numberOfSteps, len(test.assessment.Steps))
			}
		})
	}
}

// TestRunStep ensures that runStep runs the step and updates the Assessment
func TestRunStep(t *testing.T) {
	stepsTestData := []struct {
		testName string
		step     AssessmentStep
		result   Result
	}{
		{
			testName: "Failing step",
			step:     failingAssessmentStep,
			result:   Failed,
		},
		{
			testName: "Passing step",
			step:     passingAssessmentStep,
			result:   Passed,
		},
		{
			testName: "Needs review step",
			step:     needsReviewAssessmentStep,
			result:   NeedsReview,
		},
		{
			testName: "Unknown step",
			step:     unknownAssessmentStep,
			result:   Unknown,
		},
	}
	for _, test := range stepsTestData {
		t.Run(test.testName, func(t *testing.T) {
			anyOldAssessment := Assessment{}
			result := anyOldAssessment.runStep(nil, test.step)
			if result != test.result {
				t.Errorf("expected %s, got %s", test.result, result)
			}
			if anyOldAssessment.Result != test.result {
				t.Errorf("expected %s, got %s", test.result, anyOldAssessment.Result)
			}
		})
	}
}

// TestRun ensures that Run executes all steps, halting if any step does not return Passed
func TestRun(t *testing.T) {
	for _, data := range assessmentsTestData {
		t.Run(data.testName, func(t *testing.T) {
			result := data.assessment.Run(nil, testingApplicabilityString)
			if result != data.assessment.Result {
				t.Errorf("expected match between Run return value (%s) and assessment Result value (%s)", result, data.expectedResult)
			}
			if data.assessment.Steps_Executed != data.numberOfStepsToRun {
				t.Errorf("expected to run %d tests, got %d", data.numberOfStepsToRun, data.assessment.Steps_Executed)
			}
		})
	}
}

// TestRunTolerateFailures ensures that RunTolerateFailures executes all steps, halting only if a step returns an unknown result
func TestRunTolerateFailures(t *testing.T) {
	for _, d := range assessmentsTestData {
		data := d
		t.Run(data.testName, func(t *testing.T) {
			result := data.assessment.RunTolerateFailures(nil, testingApplicabilityString)
			if result != data.assessment.Result {
				t.Errorf("expected match between RunTolerateFailures return value (%s) and assessment Result value (%s)", result, data.expectedResult)
			}
			if data.assessment.Steps_Executed != data.numberOfSteps {
				if result != Unknown {
					t.Errorf("expected to run %d tests, got %d", data.numberOfSteps, data.assessment.Steps_Executed)
				}
			}
		})
		data = d
		t.Run(data.testName+"_not-applicable", func(t *testing.T) {
			result := data.assessment.RunTolerateFailures(nil, "not a real applicability")
			if len(data.assessment.Steps) > 0 && result != NotApplicable {
				t.Errorf("expected fake applicability value to return value %s but got %s", NotApplicable, result)
			}
			if data.assessment.Steps_Executed != 0 {
				t.Errorf("expected no steps to be executed, got %d", data.assessment.Steps_Executed)
			}
		})
	}
}

// TestNewChange ensures that NewChange creates a new Change object and adds it to the Assessment
func TestNewChange(t *testing.T) {
	anyOldAssessment := Assessment{}
	testName := "Add-a-new-change"
	t.Run(testName, func(t *testing.T) {
		if len(anyOldAssessment.Changes) != 0 {
			t.Errorf("Expected empty assessment object to start with 0 Change objects, got %d", len(anyOldAssessment.Changes))
		}
		change := anyOldAssessment.NewChange(testName, "", nil, goodApplyFunc, goodRevertFunc)
		if len(anyOldAssessment.Changes) != 1 {
			t.Errorf("Expected assessment object to have 1 Change object, got %d", len(anyOldAssessment.Changes))
		}
		if change == nil {
			t.Error("expected a change object to be returned by NewChange, got nil")
		}
		if change != anyOldAssessment.Changes[testName] {
			t.Errorf("Found different change object in assessment object than the one returned by NewChange")
		}

	})
}

// TestRevertChanges ensures that RevertChanges attempts to revert all changes in the Assessment
func TestRevertChanges(t *testing.T) {
	revertChangesTestData := []struct {
		testName   string
		assessment Assessment
		corrupted  bool
	}{
		{
			testName:   "No changes",
			assessment: Assessment{},
			corrupted:  false,
		},
		{
			testName:   "Change already applied and reverted",
			assessment: Assessment{Changes: map[string]*Change{"test": goodRevertedChange}},
			corrupted:  false,
		},
		{
			testName:   "Change without apply function",
			assessment: Assessment{Changes: map[string]*Change{"test": noApplyChange}},
			corrupted:  true,
		},
		{
			testName:   "Change with error from apply function",
			assessment: Assessment{Changes: map[string]*Change{"test": badApplyChange}},
			corrupted:  true,
		},
		{
			testName:   "Change with error from revert function",
			assessment: Assessment{Changes: map[string]*Change{"test": badRevertChange}},
			corrupted:  true,
		},
		{
			testName:   "Change previously applied and needs reverted",
			assessment: Assessment{Changes: map[string]*Change{"test": goodNotRevertedChange}},
			corrupted:  false,
		},
		{
			testName:   "Two changes already applied, with one already reverted",
			assessment: passingAssessment,
			corrupted:  false,
		},
	}
	for _, data := range revertChangesTestData {
		t.Run(data.testName, func(t *testing.T) {
			for _, change := range data.assessment.Changes {
				change.Apply()
			}
			corrupted := data.assessment.RevertChanges()
			if corrupted != data.corrupted {
				t.Errorf("expected corruption to be %t, got %t", data.corrupted, corrupted)
			}
		})
	}
}

func TestNewAssessment(t *testing.T) {
	newAssessmentsTestData := []struct {
		testName      string
		requirementId string
		description   string
		applicability []string
		steps         []AssessmentStep
		expectedError bool
	}{
		{
			testName:      "Empty requirementId",
			requirementId: "",
			description:   "test",
			applicability: []string{"test"},
			steps:         []AssessmentStep{passingAssessmentStep},
			expectedError: true,
		},
		{
			testName:      "Empty description",
			requirementId: "test",
			description:   "",
			applicability: []string{"test"},
			steps:         []AssessmentStep{passingAssessmentStep},
			expectedError: true,
		},
		{
			testName:      "Empty applicability",
			requirementId: "test",
			description:   "test",
			applicability: []string{},
			steps:         []AssessmentStep{passingAssessmentStep},
			expectedError: true,
		},
		{
			testName:      "Empty steps",
			requirementId: "test",
			description:   "test",
			applicability: []string{"test"},
			steps:         []AssessmentStep{},
			expectedError: true,
		},
		{
			testName:      "Good data",
			requirementId: "test",
			description:   "test",
			applicability: []string{"test"},
			steps:         []AssessmentStep{passingAssessmentStep},
			expectedError: false,
		},
	}
	for _, data := range newAssessmentsTestData {
		t.Run(data.testName, func(t *testing.T) {
			assessment, err := NewAssessment(data.requirementId, data.description, data.applicability, data.steps)
			if data.expectedError && err == nil {
				t.Error("expected error, got nil")
			}
			if !data.expectedError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if assessment == nil && !data.expectedError {
				t.Error("expected assessment object, got nil")
			}
		})
	}
}
