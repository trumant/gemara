package layer4

import "testing"

var changesTestData = []struct {
	testName string
	change   *Change
}{
	{
		testName: "Change not yet applied",
		change:   pendingChange,
	},
	{
		testName: "Change already applied and not yet reverted",
		change:   appliedNotRevertedChange,
	},
	{
		testName: "Change already applied and reverted",
		change:   appliedRevertedChange,
	},
	{
		testName: "No revert function specified",
		change:   noRevertChange,
	},
	{
		testName: "No apply function specified",
		change:   noApplyChange,
	},
	{
		testName: "Neither function specified",
		change:   &Change{},
	},
	{
		testName: "Change is not allowed to execute",
		change:   disallowedChange,
	},
}

func TestApply(t *testing.T) {
	for _, test := range changesTestData {
		t.Run(test.testName, func(t *testing.T) {

			test.change.Apply()

			if test.change.applyFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil applyFunc, but it was not")
			}
			if test.change.revertFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil revertFunc, but it was not")
			}
			if test.change.applyFunc != nil && test.change.revertFunc != nil {
				if !test.change.disallowed && !test.change.Applied {
					t.Errorf("Expected change to be applied, but it was not")
				} else if test.change.disallowed && test.change.Applied {
					t.Errorf("Expected change to not be applied, but it was")
				}

				test.change.Revert()

				if !test.change.disallowed && !test.change.Applied {
					t.Errorf("Reverting shound not erase the record that the change was applied")
				}
			}
		})
	}
}

func TestDisallow(t *testing.T) {
	for _, test := range changesTestData {
		t.Run(test.testName, func(t *testing.T) {
			if test.change.Applied {
				return // not applicable
			}
			test.change.Disallow()

			if !test.change.disallowed {
				t.Errorf("Expected change to be disallowed, but it was not")
			}
			test.change.Apply()

			if test.change.Applied {
				t.Errorf("Expected change to not be applied, but it was")
			}
		})
	}
}

func TestRevert(t *testing.T) {
	for _, test := range changesTestData {
		t.Run(test.testName, func(t *testing.T) {

			test.change.Revert()

			if test.change.applyFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil applyFunc, but it was not")
			}
			if test.change.revertFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil revertFunc, but it was not")
			}
			if test.change.applyFunc != nil && test.change.revertFunc != nil {
				if test.change.Applied && !test.change.Reverted {
					t.Errorf("Expected change to be reverted, but it was not")
				}
				if !test.change.Applied && test.change.Reverted {
					t.Errorf("Reverting should not be recorded if a change was not applied to revert")
				}
				test.change.Apply()

				if test.change.Reverted {
					t.Errorf("Applying further times shound mark the change as not reverted")
				}
			}

		})
	}
}
