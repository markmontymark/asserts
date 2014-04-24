package asserts

import (
	"fmt"
	"testing"
)

func Equals (t *testing.T, name string , got string, expected string ) {
   if got != expected {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, expected )
		t.FailNow()
   }
}

func NotEquals (t *testing.T, name string , got string, expected string ) {
   if got == expected {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected something different\n%s\n\n", name, got, expected )
		t.FailNow()
   }
}

func IntEquals (t *testing.T, name string , got int, expected int ) {
   if got != expected {
		fmt.Printf("in IntEquals with failure for test %v, got %v, expected %v\n", name, got ,expected)
      t.Errorf("Faild %s\ngot\n%d\n\nexpected\n%d\n\n", name, got, expected )
		//t.FailNow()
   }
}

func Float64Equals (t *testing.T, name string , got float64, expected float64 ) {
   if got != expected {
      t.Errorf("Failed %s\ngot\n%v\n\nexpected\n%v\n\n", name, got, expected )
		t.FailNow()
   }
}


func True (t *testing.T, name string , got bool ) {
   if !got {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, true)
		t.FailNow()
   }
}
var Ok func (*testing.T, string , bool ) = True

func False (t *testing.T, name string , got bool ) {
   if got {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, false)
		t.FailNow()
   }
}
var NotOk func (*testing.T, string , bool ) = False

func Nil (t *testing.T, name string , got ...interface{}) {
   if got == nil {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, nil)
		t.FailNow()
   }
}

func NotNil (t *testing.T, name string , got ...interface{}) {
   if got != nil {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n(not nil)\n\n", name, got)
		t.FailNow()
   }
}
