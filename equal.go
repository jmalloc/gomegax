package gomegax

import (
	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"google.golang.org/protobuf/testing/protocmp"
)

// EqualX is an alternative to gomega.Equal() that uses Google's go-cmp library
// for comparison.
//
// By default it uses the protocmp transform, making it useful for comparing
// protobuf messages.
func EqualX(
	expected interface{},
	options ...cmp.Option,
) types.GomegaMatcher {
	if len(options) == 0 {
		options = append(options, protocmp.Transform())
	}

	return &equalMatcher{
		expected: expected,
		options:  options,
	}
}

type equalMatcher struct {
	expected interface{}
	options  cmp.Options
}

func (m *equalMatcher) Match(actual interface{}) (success bool, err error) {
	return cmp.Equal(actual, m.expected, m.options), nil
}

func (m *equalMatcher) FailureMessage(actual interface{}) (message string) {
	diff := cmp.Diff(actual, m.expected, m.options)
	return "Expected no difference, got:\n" + format.IndentString(diff, 1)
}

func (m *equalMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to equal", m.expected)
}
