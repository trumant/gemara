package layer4

// This file is for reusable test data to help seed ideas and reduce duplication.

import "errors"

var (
	// Generic applicability for testing
	testingApplicability = []string{"test"}

	// Functions
	goodApplyFunc = func() (*interface{}, error) {
		return nil, nil
	}
	goodRevertFunc = func() error {
		return nil
	}
	badApplyFunc = func() (*interface{}, error) {
		return nil, errors.New("error")
	}
	badRevertFunc = func() error {
		return errors.New("error")
	}

	// Assessment Results
	passingAssessmentStep = func(targetData interface{}, a *Assessment) (Result, string) {
		return Passed, ""
	}
	failingAssessmentStep = func(targetData interface{}, a *Assessment) (Result, string) {
		return Failed, ""
	}
	needsReviewAssessmentStep = func(targetData interface{}, a *Assessment) (Result, string) {
		return NeedsReview, ""
	}
	unknownAssessmentStep = func(targetData interface{}, a *Assessment) (Result, string) {
		return Unknown, ""
	}

	// Changes
	pendingChange = &Change{
		applyFunc:  goodApplyFunc,
		revertFunc: goodRevertFunc,
	}
	appliedRevertedChange = &Change{
		applyFunc:  goodApplyFunc,
		revertFunc: goodRevertFunc,
		Applied:    true,
		Reverted:   true,
	}
	appliedNotRevertedChange = &Change{
		applyFunc:  goodApplyFunc,
		revertFunc: goodRevertFunc,
		Applied:    true,
	}
	badApplyChange = &Change{
		applyFunc:  badApplyFunc,
		revertFunc: goodRevertFunc,
	}
	badRevertChange = &Change{
		applyFunc:  goodApplyFunc,
		revertFunc: badRevertFunc,
	}
	goodRevertedChange = &Change{
		applyFunc:  goodApplyFunc,
		revertFunc: goodRevertFunc,
		Reverted:   true,
	}
	goodNotRevertedChange = &Change{
		applyFunc:  goodApplyFunc,
		revertFunc: goodRevertFunc,
		Applied:    true,
	}
	noApplyChange = &Change{
		revertFunc: goodRevertFunc,
	}
	noRevertChange = &Change{
		applyFunc: goodApplyFunc,
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
		Changes: map[string]*Change{
			"pendingChange":         pendingChange,
			"appliedRevertedChange": appliedRevertedChange,
		},
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
