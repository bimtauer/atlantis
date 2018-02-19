package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	webhooks "github.com/runatlantis/atlantis/server/events/webhooks"
)

func AnyWebhooksApplyResult() webhooks.ApplyResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(webhooks.ApplyResult))(nil)).Elem()))
	var nullValue webhooks.ApplyResult
	return nullValue
}

func EqWebhooksApplyResult(value webhooks.ApplyResult) webhooks.ApplyResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue webhooks.ApplyResult
	return nullValue
}
