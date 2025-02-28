package layer4

import "errors"

var (

	// Functions
	goodApplyFunc = func() (*interface{}, error) {
		return nil, nil
	}
	goodRevertFunc = func() error {
		return nil
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
	badRevertChange = &Change{
		applyFunc: goodApplyFunc,
		revertFunc: func() error {
			return errors.New("error")
		},
	}
	goodRevertedChange = &Change{
		applyFunc: goodApplyFunc,
		Reverted:  true,
	}
	goodNotRevertedChange = &Change{
		applyFunc: goodApplyFunc,
		Applied:   true,
	}
	noApplyChange = &Change{
		revertFunc: goodRevertFunc,
	}
	noRevertChange = &Change{
		applyFunc: goodApplyFunc,
	}

	// Assessments
	passingAssessment = Assessment{
		Changes: map[string]*Change{
			"pendingChange":         pendingChange,
			"appliedRevertedChange": appliedRevertedChange,
		},
		Steps: []AssessmentStep{passingAssessmentStep},
	}
	failingAssessment = Assessment{
		Steps: []AssessmentStep{failingAssessmentStep},
	}
	badRevertPassingAssessment = Assessment{
		Steps: []AssessmentStep{passingAssessmentStep},
		Changes: map[string]*Change{
			"pendingChange": badRevertChange,
		},
	}
)
