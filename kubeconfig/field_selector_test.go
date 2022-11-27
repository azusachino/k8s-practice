package kubeconfig

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/fields"
)

func TestSelector(t *testing.T) {
	selector()
}

// Selector foo=bar,baz=qux matched field set baz=qux,foo=bar
// Selector foo=bar matched field set baz=qux,foo=bar
// Selector qux!=abc matched field set baz=qux,foo=bar
// Selector foo=bar,baz=qux matched field set baz=qux,foo=bar
// Selector foo=bar matched field set baz=qux,foo=bar

func selector() {
	flds := fields.Set{
		"foo": "bar",
		"baz": "qux",
	}

	sel := fields.SelectorFromSet(flds)

	if sel.Matches(flds) {
		fmt.Printf("Selector %v matched field set %v\n", sel, flds)
	} else {
		panic("Selector should have matched field set")
	}

	// f==v selector.
	sel = fields.OneTermEqualSelector("foo", "bar")
	if sel.Matches(flds) {
		fmt.Printf("Selector %v matched field set %v\n", sel, flds)
	} else {
		panic("Selector should have matched field set")
	}

	// f!=v selector.
	sel = fields.OneTermNotEqualSelector("qux", "abc")
	if sel.Matches(flds) {
		fmt.Printf("Selector %v matched field set %v\n", sel, flds)
	} else {
		panic("Selector should have not matched field set")
	}

	// f1=v1,f2=v2
	sel = fields.AndSelectors(
		fields.OneTermEqualSelector("foo", "bar"),
		fields.OneTermEqualSelector("baz", "qux"),
	)
	if sel.Matches(flds) {
		fmt.Printf("Selector %v matched field set %v\n", sel, flds)
	} else {
		panic("Selector should have not matched field set")
	}

	// Selector from string expression.
	sel, err := fields.ParseSelector("foo==bar")
	if err != nil {
		panic(err.Error())
	}
	if sel.Matches(flds) {
		fmt.Printf("Selector %v matched field set %v\n", sel, flds)
	} else {
		panic("Selector should have matched field set")
	}
}
