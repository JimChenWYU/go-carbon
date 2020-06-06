package carbon

///////////////////////////////////////////////////////////////////
///////////////////////// TESTING AIDS ////////////////////////////
///////////////////////////////////////////////////////////////////

// A test Carbon instance to be returned when now instances are created.
var testNow *Carbon = nil

// Set a Carbon instance (real or mock) to be returned when a "now" instance is created.
func SetTestNow(c *Carbon) {
	testNow = c
}

// Get the Carbon instance (real or mock) to be returned when a "now" instance is created.
func TestNow() *Carbon {
	return testNow
}

// Determine if there is a valid test instance set. A valid test instance is anything that is not null.
func HasTestNow() bool {
	return TestNow() != nil
}

// for test
func wrapWithTestNow(fn func(), c *Carbon)  {
	test := TestNow()
	SetTestNow(c)
	fn()
	SetTestNow(test)
}

// for test
func wrapWithDefaultNowTestNow(fn func())  {
	wrapWithTestNow(fn, Now())
}
