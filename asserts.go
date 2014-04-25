package asserts

import (
	"fmt"
	"testing"
)

// Test a string value matches another string
func Equals (t *testing.T, name string , got string, expected string ) {
   if got != expected {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, expected )
		t.FailNow()
   }
}

// Test a string value does not match another string
func NotEquals (t *testing.T, name string , got string, expected string ) {
   if got == expected {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected something different\n%s\n\n", name, got, expected )
		t.FailNow()
   }
}

// Test an int value
func IntEquals (t *testing.T, name string , got int, expected int ) {
   if got != expected {
		fmt.Printf("in IntEquals with failure for test %v, got %v, expected %v\n", name, got ,expected)
      t.Errorf("Faild %s\ngot\n%d\n\nexpected\n%d\n\n", name, got, expected )
		//t.FailNow()
   }
}

// Test a float64 value
func Float64Equals (t *testing.T, name string , got float64, expected float64 ) {
   if got != expected {
      t.Errorf("Failed %s\ngot\n%v\n\nexpected\n%v\n\n", name, got, expected )
		t.FailNow()
   }
}

// Given a value, test that it falls inclusively within (>=) start and (<=) end. This is the opposite of Outside
func Bounded( t *testing.T, name string, value ,start, end interface{}) {
	valueI := value.(int64) / 1
	startI := start.(int64) / 1
	endI := end.(int64) / 1
	if startI > endI {
		tmp := endI
		endI = startI
		startI = tmp
	}
	if valueI < startI || valueI > endI {
      t.Errorf("Failed %s\n%v is out of bounds(%v,%v)\n\n", name, value, start, end)
		t.FailNow()
	}
}

// Given a value, test that it falls exclusively outside of the interval >= start and <= end.  This is the opposite of Bounded.
func Outside( t *testing.T, name string, value ,start, end interface{}) {
	valueI := value.(int64) / 1
	startI := start.(int64) / 1
	endI := end.(int64) / 1
	if startI > endI {
		tmp := endI
		endI = startI
		startI = tmp
	}
	if valueI >= startI || valueI <= endI {
      t.Errorf("Failed %s\n%v is within bounds(%v,%v)\n\n", name, value, start, end)
		t.FailNow()
	}
}


// Test that a value is true.   Synonym of Ok and Is
func True (t *testing.T, name string , got bool ) {
   if !got {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, true)
		t.FailNow()
   }
}

// Test that a value is true.   Synonym of True and Is
var Ok func (*testing.T, string , bool ) = True

// Test that a value is true.   Synonym of True and Ok
var Is func (*testing.T, string , bool ) = True

// Test that a value is false.   Synonym of NotOk and Isnt
func False (t *testing.T, name string , got bool ) {
   if got {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, false)
		t.FailNow()
   }
}

// Test that a value is false. Synonyms with False and Isnt
var NotOk func (*testing.T, string , bool ) = False

// Test that a value is false. Synonyms with False and NotOk
var Isnt  func (*testing.T, string , bool ) = False

// Test that a value is nil.
func Nil (t *testing.T, name string , got ...interface{}) {
   if got == nil {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n%s\n\n", name, got, nil)
		t.FailNow()
   }
}

// Test that a value is not nil.
func NotNil (t *testing.T, name string , got ...interface{}) {
   if got != nil {
      t.Errorf("Failed %s\ngot\n%s\n\nexpected\n(not nil)\n\n", name, got)
		t.FailNow()
   }
}
