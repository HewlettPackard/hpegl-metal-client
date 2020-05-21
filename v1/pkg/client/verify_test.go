package client

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/go-test/deep"

	"github.com/quattronetworks/quake-client/v1/model"
)

// TestModel verifies that the Model and Client structure definitions are the same.
// Add any new types that need to be verified to the validation.
func TestModel(t *testing.T) {
	compare(t, "NewSshKey", reflect.TypeOf(model.NewSSHKey{}), reflect.TypeOf(NewSshKey{}))
	compare(t, "SshKey", reflect.TypeOf(model.SSHKey{}), reflect.TypeOf(SshKey{}))
	compare(t, "AvailableResources", reflect.TypeOf(model.AvailableResources{}), reflect.TypeOf(AvailableResources{}))
	compare(t, "NewHost", reflect.TypeOf(model.NewHost{}), reflect.TypeOf(NewHost{}))
	compare(t, "Host", reflect.TypeOf(model.Host{}), reflect.TypeOf(Host{}))
	compare(t, "NewNetwork", reflect.TypeOf(model.NewNetwork{}), reflect.TypeOf(NewNetwork{}))
	compare(t, "Network", reflect.TypeOf(model.Network{}), reflect.TypeOf(Network{}))
	compare(t, "NewProject", reflect.TypeOf(model.NewProject{}), reflect.TypeOf(NewProject{}))
	compare(t, "Project", reflect.TypeOf(model.Project{}), reflect.TypeOf(Project{}))
	compare(t, "UsageReportRequest", reflect.TypeOf(model.UsageReportRequest{}), reflect.TypeOf(UsageReportRequest{}))
	compare(t, "UsageReport", reflect.TypeOf(model.UsageReport{}), reflect.TypeOf(UsageReport{}))
	compare(t, "VolumeAttachment", reflect.TypeOf(model.VolumeAttachment{}), reflect.TypeOf(VolumeAttachment{}))
	compare(t, "NewVolume", reflect.TypeOf(model.NewVolume{}), reflect.TypeOf(NewVolume{}))
	compare(t, "Volume", reflect.TypeOf(model.Volume{}), reflect.TypeOf(Volume{}))
	compare(t, "AddVolume", reflect.TypeOf(model.AddVolume{}), reflect.TypeOf(AddVolume{}))
}

// compare fills in a structure of the given Client type, json encodes it, decodes it into
// the Model type, re-encodes it into json, and decodes it back into the Client type. If
// the original structure and the decoded structure are the same, then we're good. If not,
// then something is wrong, and the output might give you a clue.
func compare(t *testing.T, typeName string, m, c reflect.Type) {
	mv := reflect.New(m)
	cv := reflect.New(c)
	cvCheck := reflect.New(c)

	fillStruct(t, cv)

	data, err := json.Marshal(cv.Interface())
	if err != nil {
		t.Fatalf("Unable to marshal data (%s): %v\n", typeName, err)
	}
	err = json.Unmarshal(data, cv.Interface())
	if err != nil {
		t.Fatalf("Unable to marshal data (%s): %v\n", typeName, err)
	}

	err = json.Unmarshal(data, mv.Interface())
	if err != nil {
		t.Fatalf("Unable to unmarshal data (%s): %v\n", typeName, err)
	}
	data, err = json.Marshal(mv.Interface())
	if err != nil {
		t.Fatalf("Unable to marshal data (%s): %v\n", typeName, err)
	}
	err = json.Unmarshal(data, cvCheck.Interface())
	if err != nil {
		t.Fatalf("Unable to unmarshal data (%s): %v\n", typeName, err)
	}
	if !reflect.DeepEqual(cv.Interface(), cvCheck.Interface()) {
		t.Logf("%#v\n%#v\n", cv.Interface(), cvCheck.Interface())
		diffs := deep.Equal(cv.Interface(), cvCheck.Interface())
		for _, diff := range diffs {
			t.Logf("%s\n", diff)
		}
		t.Fatalf("Comparison of model and client type %s failed", typeName)
	}
}

// fillStruct fills a structure with nonsense data.
func fillStruct(t *testing.T, v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	ty := v.Type()

	n := ty.NumField()

	for i := 0; i < n; i++ {
		fill(t, v.Field(i))
	}

}

var u64Val uint64 = 2002
var i64Val int64 = 1002

// fill fills a field, structure, etc. with nonsense data.
func fill(t *testing.T, v reflect.Value) {
	switch v.Interface().(type) {
	case time.Time:
		v.Set(reflect.ValueOf(time.Now()))
		return
	}

	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64Val++
		v.SetInt(i64Val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u64Val++
		v.SetUint(u64Val)
	case reflect.Uintptr:
		t.Logf("Uintptr are not supported in this verification test\n")
	case reflect.Float32:
		t.Logf("Float32 are not supported in this verification test\n")
	case reflect.Float64:
		t.Logf("Float64 are not supported in this verification test\n")
	case reflect.Complex64:
		t.Logf("Complex64 are not supported in this verification test\n")
	case reflect.Complex128:
		t.Logf("Complex128 are not supported in this verification test\n")
	case reflect.Array:
	case reflect.Chan:
		t.Logf("Chans are not supported in this verification test\n")
	case reflect.Func:
		t.Logf("Funcs are not supported in this verification test\n")
	case reflect.Interface:
		t.Logf("Interfaces are not supported in this verification test\n")
	case reflect.Map:
		t.Logf("Maps are not supported in this verification test\n")
	case reflect.Ptr:
		fill(t, v.Elem())
	case reflect.Slice:
		fillSlice(t, v)
	case reflect.String:
		u64Val++
		v.SetString(fmt.Sprintf("%d", u64Val))
	case reflect.Struct:
		fillStruct(t, v)
	}
}

// fillSlice fills a slice with nonsense data. It fills the slice with two elements.
func fillSlice(t *testing.T, v reflect.Value) {
	s := reflect.MakeSlice(v.Type(), 0, 2)
	x := reflect.New(v.Type().Elem())
	fill(t, x)
	s = reflect.Append(s, x.Elem())
	x = reflect.New(v.Type().Elem())
	fill(t, x)
	s = reflect.Append(s, x.Elem())
	v.Set(s)
}
