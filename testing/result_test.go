package result_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	rpb "github.com/marcushines/openresult"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		desc string
		want string
	}{{
		desc: "empty proto",
		want: "",
	}, {
		desc: "empty report",
		want: `id: "foo"
timestamp: "20160101Z2300.1000000"
`,
	}, {
		desc: "report",
		want: `id: "foo"
test_suites: <
  id: "Suite 1"
>
timestamp: "20160101Z2300.1000000"
`,
	}, {
		desc: "report with case",
		want: `id: "foo"
test_suites: <
  id: "Suite 1"
  test_cases: <
    id: "Case 1"
    results: <
      uuid: "foo"
      timestamp: 1
      duration: 10000000
      text: "This is a test result"
    >
  >
>
timestamp: "20160101Z2300.1000000"
`,
	}, {
		desc: "report with case and meta",
		want: `id: "foo"
test_suites: <
  id: "Suite 1"
  test_cases: <
    id: "Case 1"
    results: <
      uuid: "foo"
      timestamp: 1
      duration: 10000000
      text: "This is a test result"
    >
  >
  meta: <
    key: "any"
    value: <
      any_value: <
        type_url: "foo"
        value: "32"
      >
    >
  >
  meta: <
    key: "bar"
    value: <
      string_value: "asdf"
    >
  >
  meta: <
    key: "foo"
    value: <
      int_value: 34
    >
  >
>
timestamp: "20160101Z2300.1000000"
`,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			r := &rpb.Report{}
			err := proto.UnmarshalText(tt.want, r)
			if err != nil {
				t.Errorf("proto.Marshal(%q) failed: got %v", r, err)
				return
			}
			got := proto.MarshalTextString(r)
			if err != nil {
				t.Errorf("proto.Marshal(%q) failed: got %v", r, err)
				return
			}
			if got != tt.want {
				t.Errorf("proto.Marshal(%q) failed:\ngot:\n%v\nwant:\n%v\n", r, got, tt.want)
				return
			}
		})
	}
}
