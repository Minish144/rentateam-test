// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: posts.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _posts_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Posts with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Posts) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PostsValidationError is the validation error returned by Posts.Validate if
// the designated constraints aren't met.
type PostsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostsValidationError) ErrorName() string { return "PostsValidationError" }

// Error satisfies the builtin error interface
func (e PostsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosts.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostsValidationError{}

// Validate checks the field values on Posts_CreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Posts_CreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetHeader()) < 5 {
		return Posts_CreateRequestValidationError{
			field:  "Header",
			reason: "value length must be at least 5 runes",
		}
	}

	if utf8.RuneCountInString(m.GetBody()) < 30 {
		return Posts_CreateRequestValidationError{
			field:  "Body",
			reason: "value length must be at least 30 runes",
		}
	}

	return nil
}

// Posts_CreateRequestValidationError is the validation error returned by
// Posts_CreateRequest.Validate if the designated constraints aren't met.
type Posts_CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Posts_CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Posts_CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Posts_CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Posts_CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Posts_CreateRequestValidationError) ErrorName() string {
	return "Posts_CreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e Posts_CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosts_CreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Posts_CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Posts_CreateRequestValidationError{}

// Validate checks the field values on Posts_CreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Posts_CreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Posts_CreateResponseValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Posts_CreateResponseValidationError is the validation error returned by
// Posts_CreateResponse.Validate if the designated constraints aren't met.
type Posts_CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Posts_CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Posts_CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Posts_CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Posts_CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Posts_CreateResponseValidationError) ErrorName() string {
	return "Posts_CreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e Posts_CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosts_CreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Posts_CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Posts_CreateResponseValidationError{}

// Validate checks the field values on Posts_ListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Posts_ListRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// Posts_ListRequestValidationError is the validation error returned by
// Posts_ListRequest.Validate if the designated constraints aren't met.
type Posts_ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Posts_ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Posts_ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Posts_ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Posts_ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Posts_ListRequestValidationError) ErrorName() string {
	return "Posts_ListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e Posts_ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosts_ListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Posts_ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Posts_ListRequestValidationError{}

// Validate checks the field values on Posts_ListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Posts_ListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Posts_ListResponseValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Count

	return nil
}

// Posts_ListResponseValidationError is the validation error returned by
// Posts_ListResponse.Validate if the designated constraints aren't met.
type Posts_ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Posts_ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Posts_ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Posts_ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Posts_ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Posts_ListResponseValidationError) ErrorName() string {
	return "Posts_ListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e Posts_ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosts_ListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Posts_ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Posts_ListResponseValidationError{}
