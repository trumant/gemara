package layer4

// This file is for reusable test data to help seed ideas and reduce duplication.

import "errors"

var (
	// Generic applicability for testing
	testingApplicability = []string{"test-applicability"}

	// Functions
	goodApplyFunc = func() (interface{}, error) {
		return nil, nil
	}
	goodRevertFunc = func() error {
		return nil
	}
	badApplyFunc = func() (interface{}, error) {
		return nil, errors.New("error")
	}
	badRevertFunc = func() error {
		return errors.New("error")
	}

	// Assessment Results
	passingAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return Passed, ""
	}
	failingAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return Failed, ""
	}
	needsReviewAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return NeedsReview, ""
	}
	unknownAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return Unknown, ""
	}

	// Changes
	pendingChange = &Change{
		Target_Name: "pendingChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
	}
	appliedRevertedChange = &Change{
		Target_Name: "appliedRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
		Reverted:    true,
	}
	appliedNotRevertedChange = &Change{
		Target_Name: "appliedNotRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
	}
	badApplyChange = &Change{
		Target_Name: "badApplyChange",
		Description: "description placeholder",
		applyFunc:   badApplyFunc,
		revertFunc:  goodRevertFunc,
	}
	badRevertChange = &Change{
		Target_Name: "badRevertChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  badRevertFunc,
	}
	goodRevertedChange = &Change{
		Target_Name: "goodRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Reverted:    true,
	}
	goodNotRevertedChange = &Change{
		Target_Name: "goodNotRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
	}
	noApplyChange = &Change{
		Target_Name: "noApplyChange",
		Description: "description placeholder",
		revertFunc:  goodRevertFunc,
	}
	noRevertChange = &Change{
		Target_Name: "noRevertChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
	}

	// Assessments
	failingAssessment = Assessment{
		Steps: []AssessmentStep{
			failingAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
	passingAssessment = Assessment{
		Steps: []AssessmentStep{
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
	needsReviewAssessment = Assessment{
		Steps: []AssessmentStep{
			passingAssessmentStep,
			needsReviewAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
	unknownAssessment = Assessment{
		Steps: []AssessmentStep{
			passingAssessmentStep,
			unknownAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
	badRevertPassingAssessment = Assessment{
		Changes: map[string]*Change{
			"badRevertChange": badRevertChange,
		},
		Steps: []AssessmentStep{
			passingAssessmentStep,
			passingAssessmentStep,
			passingAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
)
