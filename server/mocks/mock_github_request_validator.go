// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server (interfaces: GithubRequestValidator)

package mocks

import (
	http "net/http"
	"reflect"

	pegomock "github.com/petergtz/pegomock"
)

type MockGithubRequestValidator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGithubRequestValidator() *MockGithubRequestValidator {
	return &MockGithubRequestValidator{fail: pegomock.GlobalFailHandler}
}

func (mock *MockGithubRequestValidator) Validate(r *http.Request, secret []byte) ([]byte, error) {
	params := []pegomock.Param{r, secret}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Validate", params, []reflect.Type{reflect.TypeOf((*[]byte)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []byte
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]byte)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGithubRequestValidator) VerifyWasCalledOnce() *VerifierGithubRequestValidator {
	return &VerifierGithubRequestValidator{mock, pegomock.Times(1), nil}
}

func (mock *MockGithubRequestValidator) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierGithubRequestValidator {
	return &VerifierGithubRequestValidator{mock, invocationCountMatcher, nil}
}

func (mock *MockGithubRequestValidator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierGithubRequestValidator {
	return &VerifierGithubRequestValidator{mock, invocationCountMatcher, inOrderContext}
}

type VerifierGithubRequestValidator struct {
	mock                   *MockGithubRequestValidator
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierGithubRequestValidator) Validate(r *http.Request, secret []byte) *GithubRequestValidator_Validate_OngoingVerification {
	params := []pegomock.Param{r, secret}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Validate", params)
	return &GithubRequestValidator_Validate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type GithubRequestValidator_Validate_OngoingVerification struct {
	mock              *MockGithubRequestValidator
	methodInvocations []pegomock.MethodInvocation
}

func (c *GithubRequestValidator_Validate_OngoingVerification) GetCapturedArguments() (*http.Request, []byte) {
	r, secret := c.GetAllCapturedArguments()
	return r[len(r)-1], secret[len(secret)-1]
}

func (c *GithubRequestValidator_Validate_OngoingVerification) GetAllCapturedArguments() (_param0 []*http.Request, _param1 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*http.Request, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*http.Request)
		}
		_param1 = make([][]byte, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]byte)
		}
	}
	return
}
