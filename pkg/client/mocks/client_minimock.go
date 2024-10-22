// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/alisher-baizhumanov/word-of-wisdom/pkg/client.Client -o client_minimock.go -n ClientMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	mm_client "github.com/alisher-baizhumanov/word-of-wisdom/pkg/client"
	"github.com/gojuno/minimock/v3"
)

// ClientMock implements mm_client.Client
type ClientMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetChallenge          func() (ba1 []byte, u1 uint8, err error)
	funcGetChallengeOrigin    string
	inspectFuncGetChallenge   func()
	afterGetChallengeCounter  uint64
	beforeGetChallengeCounter uint64
	GetChallengeMock          mClientMockGetChallenge

	funcSubmitSolution          func(challenge []byte, solution []byte) (q1 mm_client.Quote, err error)
	funcSubmitSolutionOrigin    string
	inspectFuncSubmitSolution   func(challenge []byte, solution []byte)
	afterSubmitSolutionCounter  uint64
	beforeSubmitSolutionCounter uint64
	SubmitSolutionMock          mClientMockSubmitSolution
}

// NewClientMock returns a mock for mm_client.Client
func NewClientMock(t minimock.Tester) *ClientMock {
	m := &ClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetChallengeMock = mClientMockGetChallenge{mock: m}

	m.SubmitSolutionMock = mClientMockSubmitSolution{mock: m}
	m.SubmitSolutionMock.callArgs = []*ClientMockSubmitSolutionParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mClientMockGetChallenge struct {
	optional           bool
	mock               *ClientMock
	defaultExpectation *ClientMockGetChallengeExpectation
	expectations       []*ClientMockGetChallengeExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ClientMockGetChallengeExpectation specifies expectation struct of the Client.GetChallenge
type ClientMockGetChallengeExpectation struct {
	mock *ClientMock

	results      *ClientMockGetChallengeResults
	returnOrigin string
	Counter      uint64
}

// ClientMockGetChallengeResults contains results of the Client.GetChallenge
type ClientMockGetChallengeResults struct {
	ba1 []byte
	u1  uint8
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetChallenge *mClientMockGetChallenge) Optional() *mClientMockGetChallenge {
	mmGetChallenge.optional = true
	return mmGetChallenge
}

// Expect sets up expected params for Client.GetChallenge
func (mmGetChallenge *mClientMockGetChallenge) Expect() *mClientMockGetChallenge {
	if mmGetChallenge.mock.funcGetChallenge != nil {
		mmGetChallenge.mock.t.Fatalf("ClientMock.GetChallenge mock is already set by Set")
	}

	if mmGetChallenge.defaultExpectation == nil {
		mmGetChallenge.defaultExpectation = &ClientMockGetChallengeExpectation{}
	}

	return mmGetChallenge
}

// Inspect accepts an inspector function that has same arguments as the Client.GetChallenge
func (mmGetChallenge *mClientMockGetChallenge) Inspect(f func()) *mClientMockGetChallenge {
	if mmGetChallenge.mock.inspectFuncGetChallenge != nil {
		mmGetChallenge.mock.t.Fatalf("Inspect function is already set for ClientMock.GetChallenge")
	}

	mmGetChallenge.mock.inspectFuncGetChallenge = f

	return mmGetChallenge
}

// Return sets up results that will be returned by Client.GetChallenge
func (mmGetChallenge *mClientMockGetChallenge) Return(ba1 []byte, u1 uint8, err error) *ClientMock {
	if mmGetChallenge.mock.funcGetChallenge != nil {
		mmGetChallenge.mock.t.Fatalf("ClientMock.GetChallenge mock is already set by Set")
	}

	if mmGetChallenge.defaultExpectation == nil {
		mmGetChallenge.defaultExpectation = &ClientMockGetChallengeExpectation{mock: mmGetChallenge.mock}
	}
	mmGetChallenge.defaultExpectation.results = &ClientMockGetChallengeResults{ba1, u1, err}
	mmGetChallenge.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetChallenge.mock
}

// Set uses given function f to mock the Client.GetChallenge method
func (mmGetChallenge *mClientMockGetChallenge) Set(f func() (ba1 []byte, u1 uint8, err error)) *ClientMock {
	if mmGetChallenge.defaultExpectation != nil {
		mmGetChallenge.mock.t.Fatalf("Default expectation is already set for the Client.GetChallenge method")
	}

	if len(mmGetChallenge.expectations) > 0 {
		mmGetChallenge.mock.t.Fatalf("Some expectations are already set for the Client.GetChallenge method")
	}

	mmGetChallenge.mock.funcGetChallenge = f
	mmGetChallenge.mock.funcGetChallengeOrigin = minimock.CallerInfo(1)
	return mmGetChallenge.mock
}

// Times sets number of times Client.GetChallenge should be invoked
func (mmGetChallenge *mClientMockGetChallenge) Times(n uint64) *mClientMockGetChallenge {
	if n == 0 {
		mmGetChallenge.mock.t.Fatalf("Times of ClientMock.GetChallenge mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetChallenge.expectedInvocations, n)
	mmGetChallenge.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetChallenge
}

func (mmGetChallenge *mClientMockGetChallenge) invocationsDone() bool {
	if len(mmGetChallenge.expectations) == 0 && mmGetChallenge.defaultExpectation == nil && mmGetChallenge.mock.funcGetChallenge == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetChallenge.mock.afterGetChallengeCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetChallenge.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetChallenge implements mm_client.Client
func (mmGetChallenge *ClientMock) GetChallenge() (ba1 []byte, u1 uint8, err error) {
	mm_atomic.AddUint64(&mmGetChallenge.beforeGetChallengeCounter, 1)
	defer mm_atomic.AddUint64(&mmGetChallenge.afterGetChallengeCounter, 1)

	mmGetChallenge.t.Helper()

	if mmGetChallenge.inspectFuncGetChallenge != nil {
		mmGetChallenge.inspectFuncGetChallenge()
	}

	if mmGetChallenge.GetChallengeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetChallenge.GetChallengeMock.defaultExpectation.Counter, 1)

		mm_results := mmGetChallenge.GetChallengeMock.defaultExpectation.results
		if mm_results == nil {
			mmGetChallenge.t.Fatal("No results are set for the ClientMock.GetChallenge")
		}
		return (*mm_results).ba1, (*mm_results).u1, (*mm_results).err
	}
	if mmGetChallenge.funcGetChallenge != nil {
		return mmGetChallenge.funcGetChallenge()
	}
	mmGetChallenge.t.Fatalf("Unexpected call to ClientMock.GetChallenge.")
	return
}

// GetChallengeAfterCounter returns a count of finished ClientMock.GetChallenge invocations
func (mmGetChallenge *ClientMock) GetChallengeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetChallenge.afterGetChallengeCounter)
}

// GetChallengeBeforeCounter returns a count of ClientMock.GetChallenge invocations
func (mmGetChallenge *ClientMock) GetChallengeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetChallenge.beforeGetChallengeCounter)
}

// MinimockGetChallengeDone returns true if the count of the GetChallenge invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockGetChallengeDone() bool {
	if m.GetChallengeMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetChallengeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetChallengeMock.invocationsDone()
}

// MinimockGetChallengeInspect logs each unmet expectation
func (m *ClientMock) MinimockGetChallengeInspect() {
	for _, e := range m.GetChallengeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ClientMock.GetChallenge")
		}
	}

	afterGetChallengeCounter := mm_atomic.LoadUint64(&m.afterGetChallengeCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetChallengeMock.defaultExpectation != nil && afterGetChallengeCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.GetChallenge at\n%s", m.GetChallengeMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetChallenge != nil && afterGetChallengeCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.GetChallenge at\n%s", m.funcGetChallengeOrigin)
	}

	if !m.GetChallengeMock.invocationsDone() && afterGetChallengeCounter > 0 {
		m.t.Errorf("Expected %d calls to ClientMock.GetChallenge at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetChallengeMock.expectedInvocations), m.GetChallengeMock.expectedInvocationsOrigin, afterGetChallengeCounter)
	}
}

type mClientMockSubmitSolution struct {
	optional           bool
	mock               *ClientMock
	defaultExpectation *ClientMockSubmitSolutionExpectation
	expectations       []*ClientMockSubmitSolutionExpectation

	callArgs []*ClientMockSubmitSolutionParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ClientMockSubmitSolutionExpectation specifies expectation struct of the Client.SubmitSolution
type ClientMockSubmitSolutionExpectation struct {
	mock               *ClientMock
	params             *ClientMockSubmitSolutionParams
	paramPtrs          *ClientMockSubmitSolutionParamPtrs
	expectationOrigins ClientMockSubmitSolutionExpectationOrigins
	results            *ClientMockSubmitSolutionResults
	returnOrigin       string
	Counter            uint64
}

// ClientMockSubmitSolutionParams contains parameters of the Client.SubmitSolution
type ClientMockSubmitSolutionParams struct {
	challenge []byte
	solution  []byte
}

// ClientMockSubmitSolutionParamPtrs contains pointers to parameters of the Client.SubmitSolution
type ClientMockSubmitSolutionParamPtrs struct {
	challenge *[]byte
	solution  *[]byte
}

// ClientMockSubmitSolutionResults contains results of the Client.SubmitSolution
type ClientMockSubmitSolutionResults struct {
	q1  mm_client.Quote
	err error
}

// ClientMockSubmitSolutionOrigins contains origins of expectations of the Client.SubmitSolution
type ClientMockSubmitSolutionExpectationOrigins struct {
	origin          string
	originChallenge string
	originSolution  string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSubmitSolution *mClientMockSubmitSolution) Optional() *mClientMockSubmitSolution {
	mmSubmitSolution.optional = true
	return mmSubmitSolution
}

// Expect sets up expected params for Client.SubmitSolution
func (mmSubmitSolution *mClientMockSubmitSolution) Expect(challenge []byte, solution []byte) *mClientMockSubmitSolution {
	if mmSubmitSolution.mock.funcSubmitSolution != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Set")
	}

	if mmSubmitSolution.defaultExpectation == nil {
		mmSubmitSolution.defaultExpectation = &ClientMockSubmitSolutionExpectation{}
	}

	if mmSubmitSolution.defaultExpectation.paramPtrs != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by ExpectParams functions")
	}

	mmSubmitSolution.defaultExpectation.params = &ClientMockSubmitSolutionParams{challenge, solution}
	mmSubmitSolution.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSubmitSolution.expectations {
		if minimock.Equal(e.params, mmSubmitSolution.defaultExpectation.params) {
			mmSubmitSolution.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSubmitSolution.defaultExpectation.params)
		}
	}

	return mmSubmitSolution
}

// ExpectChallengeParam1 sets up expected param challenge for Client.SubmitSolution
func (mmSubmitSolution *mClientMockSubmitSolution) ExpectChallengeParam1(challenge []byte) *mClientMockSubmitSolution {
	if mmSubmitSolution.mock.funcSubmitSolution != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Set")
	}

	if mmSubmitSolution.defaultExpectation == nil {
		mmSubmitSolution.defaultExpectation = &ClientMockSubmitSolutionExpectation{}
	}

	if mmSubmitSolution.defaultExpectation.params != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Expect")
	}

	if mmSubmitSolution.defaultExpectation.paramPtrs == nil {
		mmSubmitSolution.defaultExpectation.paramPtrs = &ClientMockSubmitSolutionParamPtrs{}
	}
	mmSubmitSolution.defaultExpectation.paramPtrs.challenge = &challenge
	mmSubmitSolution.defaultExpectation.expectationOrigins.originChallenge = minimock.CallerInfo(1)

	return mmSubmitSolution
}

// ExpectSolutionParam2 sets up expected param solution for Client.SubmitSolution
func (mmSubmitSolution *mClientMockSubmitSolution) ExpectSolutionParam2(solution []byte) *mClientMockSubmitSolution {
	if mmSubmitSolution.mock.funcSubmitSolution != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Set")
	}

	if mmSubmitSolution.defaultExpectation == nil {
		mmSubmitSolution.defaultExpectation = &ClientMockSubmitSolutionExpectation{}
	}

	if mmSubmitSolution.defaultExpectation.params != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Expect")
	}

	if mmSubmitSolution.defaultExpectation.paramPtrs == nil {
		mmSubmitSolution.defaultExpectation.paramPtrs = &ClientMockSubmitSolutionParamPtrs{}
	}
	mmSubmitSolution.defaultExpectation.paramPtrs.solution = &solution
	mmSubmitSolution.defaultExpectation.expectationOrigins.originSolution = minimock.CallerInfo(1)

	return mmSubmitSolution
}

// Inspect accepts an inspector function that has same arguments as the Client.SubmitSolution
func (mmSubmitSolution *mClientMockSubmitSolution) Inspect(f func(challenge []byte, solution []byte)) *mClientMockSubmitSolution {
	if mmSubmitSolution.mock.inspectFuncSubmitSolution != nil {
		mmSubmitSolution.mock.t.Fatalf("Inspect function is already set for ClientMock.SubmitSolution")
	}

	mmSubmitSolution.mock.inspectFuncSubmitSolution = f

	return mmSubmitSolution
}

// Return sets up results that will be returned by Client.SubmitSolution
func (mmSubmitSolution *mClientMockSubmitSolution) Return(q1 mm_client.Quote, err error) *ClientMock {
	if mmSubmitSolution.mock.funcSubmitSolution != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Set")
	}

	if mmSubmitSolution.defaultExpectation == nil {
		mmSubmitSolution.defaultExpectation = &ClientMockSubmitSolutionExpectation{mock: mmSubmitSolution.mock}
	}
	mmSubmitSolution.defaultExpectation.results = &ClientMockSubmitSolutionResults{q1, err}
	mmSubmitSolution.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSubmitSolution.mock
}

// Set uses given function f to mock the Client.SubmitSolution method
func (mmSubmitSolution *mClientMockSubmitSolution) Set(f func(challenge []byte, solution []byte) (q1 mm_client.Quote, err error)) *ClientMock {
	if mmSubmitSolution.defaultExpectation != nil {
		mmSubmitSolution.mock.t.Fatalf("Default expectation is already set for the Client.SubmitSolution method")
	}

	if len(mmSubmitSolution.expectations) > 0 {
		mmSubmitSolution.mock.t.Fatalf("Some expectations are already set for the Client.SubmitSolution method")
	}

	mmSubmitSolution.mock.funcSubmitSolution = f
	mmSubmitSolution.mock.funcSubmitSolutionOrigin = minimock.CallerInfo(1)
	return mmSubmitSolution.mock
}

// When sets expectation for the Client.SubmitSolution which will trigger the result defined by the following
// Then helper
func (mmSubmitSolution *mClientMockSubmitSolution) When(challenge []byte, solution []byte) *ClientMockSubmitSolutionExpectation {
	if mmSubmitSolution.mock.funcSubmitSolution != nil {
		mmSubmitSolution.mock.t.Fatalf("ClientMock.SubmitSolution mock is already set by Set")
	}

	expectation := &ClientMockSubmitSolutionExpectation{
		mock:               mmSubmitSolution.mock,
		params:             &ClientMockSubmitSolutionParams{challenge, solution},
		expectationOrigins: ClientMockSubmitSolutionExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSubmitSolution.expectations = append(mmSubmitSolution.expectations, expectation)
	return expectation
}

// Then sets up Client.SubmitSolution return parameters for the expectation previously defined by the When method
func (e *ClientMockSubmitSolutionExpectation) Then(q1 mm_client.Quote, err error) *ClientMock {
	e.results = &ClientMockSubmitSolutionResults{q1, err}
	return e.mock
}

// Times sets number of times Client.SubmitSolution should be invoked
func (mmSubmitSolution *mClientMockSubmitSolution) Times(n uint64) *mClientMockSubmitSolution {
	if n == 0 {
		mmSubmitSolution.mock.t.Fatalf("Times of ClientMock.SubmitSolution mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSubmitSolution.expectedInvocations, n)
	mmSubmitSolution.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSubmitSolution
}

func (mmSubmitSolution *mClientMockSubmitSolution) invocationsDone() bool {
	if len(mmSubmitSolution.expectations) == 0 && mmSubmitSolution.defaultExpectation == nil && mmSubmitSolution.mock.funcSubmitSolution == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSubmitSolution.mock.afterSubmitSolutionCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSubmitSolution.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SubmitSolution implements mm_client.Client
func (mmSubmitSolution *ClientMock) SubmitSolution(challenge []byte, solution []byte) (q1 mm_client.Quote, err error) {
	mm_atomic.AddUint64(&mmSubmitSolution.beforeSubmitSolutionCounter, 1)
	defer mm_atomic.AddUint64(&mmSubmitSolution.afterSubmitSolutionCounter, 1)

	mmSubmitSolution.t.Helper()

	if mmSubmitSolution.inspectFuncSubmitSolution != nil {
		mmSubmitSolution.inspectFuncSubmitSolution(challenge, solution)
	}

	mm_params := ClientMockSubmitSolutionParams{challenge, solution}

	// Record call args
	mmSubmitSolution.SubmitSolutionMock.mutex.Lock()
	mmSubmitSolution.SubmitSolutionMock.callArgs = append(mmSubmitSolution.SubmitSolutionMock.callArgs, &mm_params)
	mmSubmitSolution.SubmitSolutionMock.mutex.Unlock()

	for _, e := range mmSubmitSolution.SubmitSolutionMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.q1, e.results.err
		}
	}

	if mmSubmitSolution.SubmitSolutionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSubmitSolution.SubmitSolutionMock.defaultExpectation.Counter, 1)
		mm_want := mmSubmitSolution.SubmitSolutionMock.defaultExpectation.params
		mm_want_ptrs := mmSubmitSolution.SubmitSolutionMock.defaultExpectation.paramPtrs

		mm_got := ClientMockSubmitSolutionParams{challenge, solution}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.challenge != nil && !minimock.Equal(*mm_want_ptrs.challenge, mm_got.challenge) {
				mmSubmitSolution.t.Errorf("ClientMock.SubmitSolution got unexpected parameter challenge, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSubmitSolution.SubmitSolutionMock.defaultExpectation.expectationOrigins.originChallenge, *mm_want_ptrs.challenge, mm_got.challenge, minimock.Diff(*mm_want_ptrs.challenge, mm_got.challenge))
			}

			if mm_want_ptrs.solution != nil && !minimock.Equal(*mm_want_ptrs.solution, mm_got.solution) {
				mmSubmitSolution.t.Errorf("ClientMock.SubmitSolution got unexpected parameter solution, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSubmitSolution.SubmitSolutionMock.defaultExpectation.expectationOrigins.originSolution, *mm_want_ptrs.solution, mm_got.solution, minimock.Diff(*mm_want_ptrs.solution, mm_got.solution))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSubmitSolution.t.Errorf("ClientMock.SubmitSolution got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSubmitSolution.SubmitSolutionMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSubmitSolution.SubmitSolutionMock.defaultExpectation.results
		if mm_results == nil {
			mmSubmitSolution.t.Fatal("No results are set for the ClientMock.SubmitSolution")
		}
		return (*mm_results).q1, (*mm_results).err
	}
	if mmSubmitSolution.funcSubmitSolution != nil {
		return mmSubmitSolution.funcSubmitSolution(challenge, solution)
	}
	mmSubmitSolution.t.Fatalf("Unexpected call to ClientMock.SubmitSolution. %v %v", challenge, solution)
	return
}

// SubmitSolutionAfterCounter returns a count of finished ClientMock.SubmitSolution invocations
func (mmSubmitSolution *ClientMock) SubmitSolutionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSubmitSolution.afterSubmitSolutionCounter)
}

// SubmitSolutionBeforeCounter returns a count of ClientMock.SubmitSolution invocations
func (mmSubmitSolution *ClientMock) SubmitSolutionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSubmitSolution.beforeSubmitSolutionCounter)
}

// Calls returns a list of arguments used in each call to ClientMock.SubmitSolution.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSubmitSolution *mClientMockSubmitSolution) Calls() []*ClientMockSubmitSolutionParams {
	mmSubmitSolution.mutex.RLock()

	argCopy := make([]*ClientMockSubmitSolutionParams, len(mmSubmitSolution.callArgs))
	copy(argCopy, mmSubmitSolution.callArgs)

	mmSubmitSolution.mutex.RUnlock()

	return argCopy
}

// MinimockSubmitSolutionDone returns true if the count of the SubmitSolution invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockSubmitSolutionDone() bool {
	if m.SubmitSolutionMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SubmitSolutionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SubmitSolutionMock.invocationsDone()
}

// MinimockSubmitSolutionInspect logs each unmet expectation
func (m *ClientMock) MinimockSubmitSolutionInspect() {
	for _, e := range m.SubmitSolutionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMock.SubmitSolution at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSubmitSolutionCounter := mm_atomic.LoadUint64(&m.afterSubmitSolutionCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SubmitSolutionMock.defaultExpectation != nil && afterSubmitSolutionCounter < 1 {
		if m.SubmitSolutionMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to ClientMock.SubmitSolution at\n%s", m.SubmitSolutionMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to ClientMock.SubmitSolution at\n%s with params: %#v", m.SubmitSolutionMock.defaultExpectation.expectationOrigins.origin, *m.SubmitSolutionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSubmitSolution != nil && afterSubmitSolutionCounter < 1 {
		m.t.Errorf("Expected call to ClientMock.SubmitSolution at\n%s", m.funcSubmitSolutionOrigin)
	}

	if !m.SubmitSolutionMock.invocationsDone() && afterSubmitSolutionCounter > 0 {
		m.t.Errorf("Expected %d calls to ClientMock.SubmitSolution at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SubmitSolutionMock.expectedInvocations), m.SubmitSolutionMock.expectedInvocationsOrigin, afterSubmitSolutionCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetChallengeInspect()

			m.MinimockSubmitSolutionInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetChallengeDone() &&
		m.MinimockSubmitSolutionDone()
}