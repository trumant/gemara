package layer4

import (
	"errors"
	"testing"
)

// TestTestResult is a test function for TestResult
func TestAssessmentResult(t *testing.T) {
	testData := []struct {
		testName     string
		expectedPass bool
		value        interface{}
	}{
		{
			testName:     "Pass with message and value of type string",
			expectedPass: true,
			value:        "value",
		},
		{
			testName:     "Pass with message and value of type int",
			expectedPass: true,
			value:        1,
		},
		{
			testName:     "Pass with message and value of type string slice",
			expectedPass: true,
			value:        []string{"value"},
		},
		{
			testName:     "Fail with message and value of type string",
			expectedPass: false,
			value:        "value",
		},
		{
			testName:     "Fail with message and value of type int",
			expectedPass: false,
			value:        1,
		},
		{
			testName:     "Fail with message and value of type string slice",
			expectedPass: false,
			value:        []string{"value"},
		},
	}
	t.Run("TestResult", func(t *testing.T) {
		for _, tt := range testData {
			t.Run(tt.testName, func(t *testing.T) {
				testResult := AssessmentResult{}
				testMessage := "this should never change"

				if tt.expectedPass {
					testResult.SetPass(testMessage, tt.value)
				} else {
					testResult.SetFail(testMessage, tt.value)
				}

				if testResult.Passed != tt.expectedPass {
					t.Errorf("Expected test to pass: %t, got: %t", tt.expectedPass, testResult.Passed)
				}

				if testResult.Message != testMessage {
					t.Errorf("Expected message: %s, got: %s", testMessage, testResult.Message)
				}

				switch testResult.Value.(type) {
				case string:
					if testResult.Value.(string) != tt.value {
						t.Errorf("Expected value: %v, got: %v", tt.value, testResult.Value)
					}
				case int:
					if testResult.Value.(int) != tt.value {
						t.Errorf("Expected value: %v, got: %v", tt.value, testResult.Value)
					}
				case []string:
					if testResult.Value.([]string)[0] != tt.value.([]string)[0] {
						t.Errorf("Expected value: %v, got: %v", tt.value, testResult.Value)
					}
				}
			})
		}
	})
}

// TestRevertChanges is a test function for RevertChanges
func TestRevertChanges(t *testing.T) {
	reusableTestRevertChanges(t)
}

// This is also used by TestControlEvaluationCleanup
func reusableTestRevertChanges(t *testing.T) {
	goodChange := &Change{
		Applied: true,
		revertFunc: func() error {
			return nil
		},
	}
	badRevertChange := &Change{
		Applied: true,
		revertFunc: func() error {
			return errors.New("error")
		},
	}
	testData := []struct {
		testName              string
		expectedPass          bool
		expectedBadStateAlert bool
		changes               map[string]*Change
	}{

		{
			testName:     "RevertChanges with good changes",
			expectedPass: true,
		},
		{
			testName:              "RevertChanges with good and bad changes",
			expectedBadStateAlert: true,
			changes: map[string]*Change{
				"goodChange":      goodChange,
				"badRevertChange": badRevertChange,
			},
		},
		{
			testName:              "RevertChanges with good and bad changes (2)",
			expectedBadStateAlert: true,
			changes: map[string]*Change{
				"badApplyChange": badRevertChange,
				"goodChange":     goodChange,
			},
		},
		{
			testName:              "RevertChanges with good and bad changes (3)",
			expectedBadStateAlert: true,
			changes: map[string]*Change{
				"goodChange":      goodChange,
				"badRevertChange": badRevertChange,
				"goodChange2":     goodChange,
			},
		},
		{
			testName:              "RevertChanges with good and bad changes (4)",
			expectedBadStateAlert: true,
			changes: map[string]*Change{
				"badRevertChange": badRevertChange,
			},
		},
		{
			testName:              "RevertChanges with good and bad changes (5)",
			expectedBadStateAlert: true,
			changes: map[string]*Change{
				"badRevertChange":  badRevertChange,
				"badRevertChange2": badRevertChange,
			},
		},
	}
	t.Run("RevertChanges", func(t *testing.T) {
		for _, tt := range testData {
			t.Run(tt.testName, func(t *testing.T) {
				testResult := AssessmentResult{
					Changes: tt.changes,
				}
				badStateAlert := testResult.RevertChanges()
				if badStateAlert != tt.expectedBadStateAlert {
					t.Errorf("Expected badStateAlert: %t, got: %t", tt.expectedBadStateAlert, badStateAlert)
				}
			})
		}
	})
}
