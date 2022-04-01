package main

import (
	"fmt"
	"go-sandbox/simplemath"
	"net/http"
	"strings"
)

func main() {
	answer, err := simplemath.Divide(6.0, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}

	sv := NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	println(sv.String())
	println(SemanticVersionString(sv))

	//------------------

	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "https://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	//-------
	func() {
		println("This is an anonymous func")
	}()
	//-----------

	//-------
	a := func(name string) {
		println("another anonymous func with a param of %s", name)
	}

	a("f1")
	a("f2")
	a("f3")

	//-------

}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(r *http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil

}

func mathExpression() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}
