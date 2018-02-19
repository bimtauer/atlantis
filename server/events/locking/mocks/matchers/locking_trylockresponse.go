package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	locking "github.com/runatlantis/atlantis/server/events/locking"
)

func AnyLockingTryLockResponse() locking.TryLockResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(locking.TryLockResponse))(nil)).Elem()))
	var nullValue locking.TryLockResponse
	return nullValue
}

func EqLockingTryLockResponse(value locking.TryLockResponse) locking.TryLockResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue locking.TryLockResponse
	return nullValue
}
