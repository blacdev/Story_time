package file

import (
	"reflect"
	"testing"
)


func TestReadFile(t *testing.T){
 
	want := []byte(`{"intro": {"title": "The Little Blue Gopher"}}`)

	got, _ := ReadFile("sample_test.json")

	if !reflect.DeepEqual(want, got){
		t.Errorf("got %v wanted %v", got, want)
	}
}