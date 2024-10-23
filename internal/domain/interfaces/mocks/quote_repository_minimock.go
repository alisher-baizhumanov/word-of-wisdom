// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/interfaces.QuoteRepository -o quote_repository_minimock.go -n QuoteRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/entity"
	"github.com/gojuno/minimock/v3"
)

// QuoteRepositoryMock implements mm_interfaces.QuoteRepository
type QuoteRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetRandomQuote          func(ctx context.Context) (q1 entity.Quote, err error)
	funcGetRandomQuoteOrigin    string
	inspectFuncGetRandomQuote   func(ctx context.Context)
	afterGetRandomQuoteCounter  uint64
	beforeGetRandomQuoteCounter uint64
	GetRandomQuoteMock          mQuoteRepositoryMockGetRandomQuote
}

// NewQuoteRepositoryMock returns a mock for mm_interfaces.QuoteRepository
func NewQuoteRepositoryMock(t minimock.Tester) *QuoteRepositoryMock {
	m := &QuoteRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetRandomQuoteMock = mQuoteRepositoryMockGetRandomQuote{mock: m}
	m.GetRandomQuoteMock.callArgs = []*QuoteRepositoryMockGetRandomQuoteParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mQuoteRepositoryMockGetRandomQuote struct {
	optional           bool
	mock               *QuoteRepositoryMock
	defaultExpectation *QuoteRepositoryMockGetRandomQuoteExpectation
	expectations       []*QuoteRepositoryMockGetRandomQuoteExpectation

	callArgs []*QuoteRepositoryMockGetRandomQuoteParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// QuoteRepositoryMockGetRandomQuoteExpectation specifies expectation struct of the QuoteRepository.GetRandomQuote
type QuoteRepositoryMockGetRandomQuoteExpectation struct {
	mock               *QuoteRepositoryMock
	params             *QuoteRepositoryMockGetRandomQuoteParams
	paramPtrs          *QuoteRepositoryMockGetRandomQuoteParamPtrs
	expectationOrigins QuoteRepositoryMockGetRandomQuoteExpectationOrigins
	results            *QuoteRepositoryMockGetRandomQuoteResults
	returnOrigin       string
	Counter            uint64
}

// QuoteRepositoryMockGetRandomQuoteParams contains parameters of the QuoteRepository.GetRandomQuote
type QuoteRepositoryMockGetRandomQuoteParams struct {
	ctx context.Context
}

// QuoteRepositoryMockGetRandomQuoteParamPtrs contains pointers to parameters of the QuoteRepository.GetRandomQuote
type QuoteRepositoryMockGetRandomQuoteParamPtrs struct {
	ctx *context.Context
}

// QuoteRepositoryMockGetRandomQuoteResults contains results of the QuoteRepository.GetRandomQuote
type QuoteRepositoryMockGetRandomQuoteResults struct {
	q1  entity.Quote
	err error
}

// QuoteRepositoryMockGetRandomQuoteOrigins contains origins of expectations of the QuoteRepository.GetRandomQuote
type QuoteRepositoryMockGetRandomQuoteExpectationOrigins struct {
	origin    string
	originCtx string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Optional() *mQuoteRepositoryMockGetRandomQuote {
	mmGetRandomQuote.optional = true
	return mmGetRandomQuote
}

// Expect sets up expected params for QuoteRepository.GetRandomQuote
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Expect(ctx context.Context) *mQuoteRepositoryMockGetRandomQuote {
	if mmGetRandomQuote.mock.funcGetRandomQuote != nil {
		mmGetRandomQuote.mock.t.Fatalf("QuoteRepositoryMock.GetRandomQuote mock is already set by Set")
	}

	if mmGetRandomQuote.defaultExpectation == nil {
		mmGetRandomQuote.defaultExpectation = &QuoteRepositoryMockGetRandomQuoteExpectation{}
	}

	if mmGetRandomQuote.defaultExpectation.paramPtrs != nil {
		mmGetRandomQuote.mock.t.Fatalf("QuoteRepositoryMock.GetRandomQuote mock is already set by ExpectParams functions")
	}

	mmGetRandomQuote.defaultExpectation.params = &QuoteRepositoryMockGetRandomQuoteParams{ctx}
	mmGetRandomQuote.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGetRandomQuote.expectations {
		if minimock.Equal(e.params, mmGetRandomQuote.defaultExpectation.params) {
			mmGetRandomQuote.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetRandomQuote.defaultExpectation.params)
		}
	}

	return mmGetRandomQuote
}

// ExpectCtxParam1 sets up expected param ctx for QuoteRepository.GetRandomQuote
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) ExpectCtxParam1(ctx context.Context) *mQuoteRepositoryMockGetRandomQuote {
	if mmGetRandomQuote.mock.funcGetRandomQuote != nil {
		mmGetRandomQuote.mock.t.Fatalf("QuoteRepositoryMock.GetRandomQuote mock is already set by Set")
	}

	if mmGetRandomQuote.defaultExpectation == nil {
		mmGetRandomQuote.defaultExpectation = &QuoteRepositoryMockGetRandomQuoteExpectation{}
	}

	if mmGetRandomQuote.defaultExpectation.params != nil {
		mmGetRandomQuote.mock.t.Fatalf("QuoteRepositoryMock.GetRandomQuote mock is already set by Expect")
	}

	if mmGetRandomQuote.defaultExpectation.paramPtrs == nil {
		mmGetRandomQuote.defaultExpectation.paramPtrs = &QuoteRepositoryMockGetRandomQuoteParamPtrs{}
	}
	mmGetRandomQuote.defaultExpectation.paramPtrs.ctx = &ctx
	mmGetRandomQuote.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGetRandomQuote
}

// Inspect accepts an inspector function that has same arguments as the QuoteRepository.GetRandomQuote
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Inspect(f func(ctx context.Context)) *mQuoteRepositoryMockGetRandomQuote {
	if mmGetRandomQuote.mock.inspectFuncGetRandomQuote != nil {
		mmGetRandomQuote.mock.t.Fatalf("Inspect function is already set for QuoteRepositoryMock.GetRandomQuote")
	}

	mmGetRandomQuote.mock.inspectFuncGetRandomQuote = f

	return mmGetRandomQuote
}

// Return sets up results that will be returned by QuoteRepository.GetRandomQuote
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Return(q1 entity.Quote, err error) *QuoteRepositoryMock {
	if mmGetRandomQuote.mock.funcGetRandomQuote != nil {
		mmGetRandomQuote.mock.t.Fatalf("QuoteRepositoryMock.GetRandomQuote mock is already set by Set")
	}

	if mmGetRandomQuote.defaultExpectation == nil {
		mmGetRandomQuote.defaultExpectation = &QuoteRepositoryMockGetRandomQuoteExpectation{mock: mmGetRandomQuote.mock}
	}
	mmGetRandomQuote.defaultExpectation.results = &QuoteRepositoryMockGetRandomQuoteResults{q1, err}
	mmGetRandomQuote.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetRandomQuote.mock
}

// Set uses given function f to mock the QuoteRepository.GetRandomQuote method
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Set(f func(ctx context.Context) (q1 entity.Quote, err error)) *QuoteRepositoryMock {
	if mmGetRandomQuote.defaultExpectation != nil {
		mmGetRandomQuote.mock.t.Fatalf("Default expectation is already set for the QuoteRepository.GetRandomQuote method")
	}

	if len(mmGetRandomQuote.expectations) > 0 {
		mmGetRandomQuote.mock.t.Fatalf("Some expectations are already set for the QuoteRepository.GetRandomQuote method")
	}

	mmGetRandomQuote.mock.funcGetRandomQuote = f
	mmGetRandomQuote.mock.funcGetRandomQuoteOrigin = minimock.CallerInfo(1)
	return mmGetRandomQuote.mock
}

// When sets expectation for the QuoteRepository.GetRandomQuote which will trigger the result defined by the following
// Then helper
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) When(ctx context.Context) *QuoteRepositoryMockGetRandomQuoteExpectation {
	if mmGetRandomQuote.mock.funcGetRandomQuote != nil {
		mmGetRandomQuote.mock.t.Fatalf("QuoteRepositoryMock.GetRandomQuote mock is already set by Set")
	}

	expectation := &QuoteRepositoryMockGetRandomQuoteExpectation{
		mock:               mmGetRandomQuote.mock,
		params:             &QuoteRepositoryMockGetRandomQuoteParams{ctx},
		expectationOrigins: QuoteRepositoryMockGetRandomQuoteExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGetRandomQuote.expectations = append(mmGetRandomQuote.expectations, expectation)
	return expectation
}

// Then sets up QuoteRepository.GetRandomQuote return parameters for the expectation previously defined by the When method
func (e *QuoteRepositoryMockGetRandomQuoteExpectation) Then(q1 entity.Quote, err error) *QuoteRepositoryMock {
	e.results = &QuoteRepositoryMockGetRandomQuoteResults{q1, err}
	return e.mock
}

// Times sets number of times QuoteRepository.GetRandomQuote should be invoked
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Times(n uint64) *mQuoteRepositoryMockGetRandomQuote {
	if n == 0 {
		mmGetRandomQuote.mock.t.Fatalf("Times of QuoteRepositoryMock.GetRandomQuote mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetRandomQuote.expectedInvocations, n)
	mmGetRandomQuote.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetRandomQuote
}

func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) invocationsDone() bool {
	if len(mmGetRandomQuote.expectations) == 0 && mmGetRandomQuote.defaultExpectation == nil && mmGetRandomQuote.mock.funcGetRandomQuote == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetRandomQuote.mock.afterGetRandomQuoteCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetRandomQuote.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetRandomQuote implements mm_interfaces.QuoteRepository
func (mmGetRandomQuote *QuoteRepositoryMock) GetRandomQuote(ctx context.Context) (q1 entity.Quote, err error) {
	mm_atomic.AddUint64(&mmGetRandomQuote.beforeGetRandomQuoteCounter, 1)
	defer mm_atomic.AddUint64(&mmGetRandomQuote.afterGetRandomQuoteCounter, 1)

	mmGetRandomQuote.t.Helper()

	if mmGetRandomQuote.inspectFuncGetRandomQuote != nil {
		mmGetRandomQuote.inspectFuncGetRandomQuote(ctx)
	}

	mm_params := QuoteRepositoryMockGetRandomQuoteParams{ctx}

	// Record call args
	mmGetRandomQuote.GetRandomQuoteMock.mutex.Lock()
	mmGetRandomQuote.GetRandomQuoteMock.callArgs = append(mmGetRandomQuote.GetRandomQuoteMock.callArgs, &mm_params)
	mmGetRandomQuote.GetRandomQuoteMock.mutex.Unlock()

	for _, e := range mmGetRandomQuote.GetRandomQuoteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.q1, e.results.err
		}
	}

	if mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation.Counter, 1)
		mm_want := mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation.params
		mm_want_ptrs := mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation.paramPtrs

		mm_got := QuoteRepositoryMockGetRandomQuoteParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGetRandomQuote.t.Errorf("QuoteRepositoryMock.GetRandomQuote got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetRandomQuote.t.Errorf("QuoteRepositoryMock.GetRandomQuote got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetRandomQuote.GetRandomQuoteMock.defaultExpectation.results
		if mm_results == nil {
			mmGetRandomQuote.t.Fatal("No results are set for the QuoteRepositoryMock.GetRandomQuote")
		}
		return (*mm_results).q1, (*mm_results).err
	}
	if mmGetRandomQuote.funcGetRandomQuote != nil {
		return mmGetRandomQuote.funcGetRandomQuote(ctx)
	}
	mmGetRandomQuote.t.Fatalf("Unexpected call to QuoteRepositoryMock.GetRandomQuote. %v", ctx)
	return
}

// GetRandomQuoteAfterCounter returns a count of finished QuoteRepositoryMock.GetRandomQuote invocations
func (mmGetRandomQuote *QuoteRepositoryMock) GetRandomQuoteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRandomQuote.afterGetRandomQuoteCounter)
}

// GetRandomQuoteBeforeCounter returns a count of QuoteRepositoryMock.GetRandomQuote invocations
func (mmGetRandomQuote *QuoteRepositoryMock) GetRandomQuoteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetRandomQuote.beforeGetRandomQuoteCounter)
}

// Calls returns a list of arguments used in each call to QuoteRepositoryMock.GetRandomQuote.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetRandomQuote *mQuoteRepositoryMockGetRandomQuote) Calls() []*QuoteRepositoryMockGetRandomQuoteParams {
	mmGetRandomQuote.mutex.RLock()

	argCopy := make([]*QuoteRepositoryMockGetRandomQuoteParams, len(mmGetRandomQuote.callArgs))
	copy(argCopy, mmGetRandomQuote.callArgs)

	mmGetRandomQuote.mutex.RUnlock()

	return argCopy
}

// MinimockGetRandomQuoteDone returns true if the count of the GetRandomQuote invocations corresponds
// the number of defined expectations
func (m *QuoteRepositoryMock) MinimockGetRandomQuoteDone() bool {
	if m.GetRandomQuoteMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetRandomQuoteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetRandomQuoteMock.invocationsDone()
}

// MinimockGetRandomQuoteInspect logs each unmet expectation
func (m *QuoteRepositoryMock) MinimockGetRandomQuoteInspect() {
	for _, e := range m.GetRandomQuoteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to QuoteRepositoryMock.GetRandomQuote at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetRandomQuoteCounter := mm_atomic.LoadUint64(&m.afterGetRandomQuoteCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetRandomQuoteMock.defaultExpectation != nil && afterGetRandomQuoteCounter < 1 {
		if m.GetRandomQuoteMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to QuoteRepositoryMock.GetRandomQuote at\n%s", m.GetRandomQuoteMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to QuoteRepositoryMock.GetRandomQuote at\n%s with params: %#v", m.GetRandomQuoteMock.defaultExpectation.expectationOrigins.origin, *m.GetRandomQuoteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetRandomQuote != nil && afterGetRandomQuoteCounter < 1 {
		m.t.Errorf("Expected call to QuoteRepositoryMock.GetRandomQuote at\n%s", m.funcGetRandomQuoteOrigin)
	}

	if !m.GetRandomQuoteMock.invocationsDone() && afterGetRandomQuoteCounter > 0 {
		m.t.Errorf("Expected %d calls to QuoteRepositoryMock.GetRandomQuote at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetRandomQuoteMock.expectedInvocations), m.GetRandomQuoteMock.expectedInvocationsOrigin, afterGetRandomQuoteCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *QuoteRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetRandomQuoteInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *QuoteRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *QuoteRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetRandomQuoteDone()
}
