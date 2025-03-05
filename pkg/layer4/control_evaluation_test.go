package layer4

import "testing"

var controlEvaluationTestData = []struct {
	testName          string
	control           *ControlEvaluation
	failBeforePass    bool
	expectedResult    Result
	expectedCorrupted bool
}{
	{
		testName:          "ControlEvaluation with no Assessments",
		expectedResult:    NeedsReview,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{},
		},
	},
	{
		testName:          "ControlEvaluation with one passing Assessment",
		expectedResult:    Passed,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{&passingAssessment},
		},
	},
	{
		testName:          "ControlEvaluation with one failing Assessment",
		expectedResult:    Failed,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{&failingAssessment},
		},
	},
	{
		testName:          "ControlEvaluation with one NeedsReview Assessment",
		expectedResult:    NeedsReview,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{&needsReviewAssessment},
		},
	},
	{
		testName:          "ControlEvaluation with one Unknown Assessment",
		expectedResult:    Unknown,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{&unknownAssessment},
		},
	},
	{
		testName:          "ControlEvaluation with first NeedsReview and then Unknown Assessment",
		expectedResult:    Unknown,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				&needsReviewAssessment,
				&unknownAssessment,
			},
		},
	},
	{
		testName:          "ControlEvaluation with first Unknown and then NeedsReview Assessment",
		expectedResult:    Unknown,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				&unknownAssessment,
				&needsReviewAssessment,
			},
		},
	},
	{
		testName:          "ControlEvaluation with first Failed and then NeedsReview Assessment",
		expectedResult:    Failed,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				&failingAssessment,
				&needsReviewAssessment,
			},
		},
	},
	{
		testName:          "ControlEvaluation with first Failing and then Passing Assessment",
		expectedResult:    Failed,
		failBeforePass:    true,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				&failingAssessment,
				&passingAssessment,
			},
		},
	},
}

// TestEvaluate runs a series of tests on the ControlEvaluation.Evaluate method
func TestEvaluate(t *testing.T) {
	for _, c := range controlEvaluationTestData {
		t.Run(c.testName, func(t *testing.T) {

			c.control.Evaluate(nil, testingApplicability)

			if c.control.Result != c.expectedResult {
				t.Errorf("Expected Result to be %v, but it was %v", c.expectedResult, c.control.Result)
			}

			if c.control.Corrupted_State != c.expectedCorrupted {
				t.Errorf("Expected Corrupted_State to be %v, but it was %v", c.expectedCorrupted, c.control.Corrupted_State)
			}
		})
	}
}

func TestAddAssesment(t *testing.T) {

	controlEvaluationTestData[0].control.AddAssessment("test", "test", []string{}, []AssessmentStep{})

	if controlEvaluationTestData[0].control.Result != Failed {
		t.Errorf("Expected Result to be Failed, but it was %v", controlEvaluationTestData[0].control.Result)
	}

	if controlEvaluationTestData[0].control.Message != "expected all Assessment fields to have a value, but got: requirementId=len(4), description=len=(4), applicability=len(0), steps=len(0)" {
		t.Errorf("Expected error message to be 'expected all Assessment fields to have a value, but got: requirementId=len(4), description=len=(4), applicability=len(0), steps=len(0)', but instead it was '%v'", controlEvaluationTestData[0].control.Message)
	}

}
