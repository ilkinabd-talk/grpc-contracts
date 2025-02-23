// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: google/rpc/context/attribute_context.proto

package attribute_context

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AttributeContext with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContextMultiError, or nil if none found.
func (m *AttributeContext) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrigin()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Origin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Origin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrigin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Origin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Source",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Source",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDestination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Destination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Destination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDestination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Destination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetApi()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Api",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContextValidationError{
					field:  "Api",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApi()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Api",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetExtensions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttributeContextValidationError{
						field:  fmt.Sprintf("Extensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttributeContextValidationError{
						field:  fmt.Sprintf("Extensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttributeContextValidationError{
					field:  fmt.Sprintf("Extensions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AttributeContextMultiError(errors)
	}

	return nil
}

// AttributeContextMultiError is an error wrapping multiple validation errors
// returned by AttributeContext.ValidateAll() if the designated constraints
// aren't met.
type AttributeContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContextMultiError) AllErrors() []error { return m }

// AttributeContextValidationError is the validation error returned by
// AttributeContext.Validate if the designated constraints aren't met.
type AttributeContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContextValidationError) ErrorName() string { return "AttributeContextValidationError" }

// Error satisfies the builtin error interface
func (e AttributeContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContextValidationError{}

// Validate checks the field values on AttributeContext_Peer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext_Peer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext_Peer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContext_PeerMultiError, or nil if none found.
func (m *AttributeContext_Peer) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext_Peer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ip

	// no validation rules for Port

	// no validation rules for Labels

	// no validation rules for Principal

	// no validation rules for RegionCode

	if len(errors) > 0 {
		return AttributeContext_PeerMultiError(errors)
	}

	return nil
}

// AttributeContext_PeerMultiError is an error wrapping multiple validation
// errors returned by AttributeContext_Peer.ValidateAll() if the designated
// constraints aren't met.
type AttributeContext_PeerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContext_PeerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContext_PeerMultiError) AllErrors() []error { return m }

// AttributeContext_PeerValidationError is the validation error returned by
// AttributeContext_Peer.Validate if the designated constraints aren't met.
type AttributeContext_PeerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_PeerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_PeerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_PeerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_PeerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_PeerValidationError) ErrorName() string {
	return "AttributeContext_PeerValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_PeerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Peer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_PeerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_PeerValidationError{}

// Validate checks the field values on AttributeContext_Api with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext_Api) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext_Api with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContext_ApiMultiError, or nil if none found.
func (m *AttributeContext_Api) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext_Api) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	// no validation rules for Operation

	// no validation rules for Protocol

	// no validation rules for Version

	if len(errors) > 0 {
		return AttributeContext_ApiMultiError(errors)
	}

	return nil
}

// AttributeContext_ApiMultiError is an error wrapping multiple validation
// errors returned by AttributeContext_Api.ValidateAll() if the designated
// constraints aren't met.
type AttributeContext_ApiMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContext_ApiMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContext_ApiMultiError) AllErrors() []error { return m }

// AttributeContext_ApiValidationError is the validation error returned by
// AttributeContext_Api.Validate if the designated constraints aren't met.
type AttributeContext_ApiValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_ApiValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_ApiValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_ApiValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_ApiValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_ApiValidationError) ErrorName() string {
	return "AttributeContext_ApiValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_ApiValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Api.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_ApiValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_ApiValidationError{}

// Validate checks the field values on AttributeContext_Auth with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext_Auth) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext_Auth with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContext_AuthMultiError, or nil if none found.
func (m *AttributeContext_Auth) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext_Auth) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Principal

	// no validation rules for Presenter

	if all {
		switch v := interface{}(m.GetClaims()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_AuthValidationError{
					field:  "Claims",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_AuthValidationError{
					field:  "Claims",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClaims()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_AuthValidationError{
				field:  "Claims",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeContext_AuthMultiError(errors)
	}

	return nil
}

// AttributeContext_AuthMultiError is an error wrapping multiple validation
// errors returned by AttributeContext_Auth.ValidateAll() if the designated
// constraints aren't met.
type AttributeContext_AuthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContext_AuthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContext_AuthMultiError) AllErrors() []error { return m }

// AttributeContext_AuthValidationError is the validation error returned by
// AttributeContext_Auth.Validate if the designated constraints aren't met.
type AttributeContext_AuthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_AuthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_AuthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_AuthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_AuthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_AuthValidationError) ErrorName() string {
	return "AttributeContext_AuthValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_AuthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Auth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_AuthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_AuthValidationError{}

// Validate checks the field values on AttributeContext_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext_Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext_Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContext_RequestMultiError, or nil if none found.
func (m *AttributeContext_Request) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext_Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Method

	// no validation rules for Headers

	// no validation rules for Path

	// no validation rules for Host

	// no validation rules for Scheme

	// no validation rules for Query

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_RequestValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_RequestValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_RequestValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Size

	// no validation rules for Protocol

	// no validation rules for Reason

	if all {
		switch v := interface{}(m.GetAuth()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_RequestValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_RequestValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_RequestValidationError{
				field:  "Auth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeContext_RequestMultiError(errors)
	}

	return nil
}

// AttributeContext_RequestMultiError is an error wrapping multiple validation
// errors returned by AttributeContext_Request.ValidateAll() if the designated
// constraints aren't met.
type AttributeContext_RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContext_RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContext_RequestMultiError) AllErrors() []error { return m }

// AttributeContext_RequestValidationError is the validation error returned by
// AttributeContext_Request.Validate if the designated constraints aren't met.
type AttributeContext_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_RequestValidationError) ErrorName() string {
	return "AttributeContext_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_RequestValidationError{}

// Validate checks the field values on AttributeContext_Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext_Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext_Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContext_ResponseMultiError, or nil if none found.
func (m *AttributeContext_Response) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext_Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Size

	// no validation rules for Headers

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_ResponseValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_ResponseValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_ResponseValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBackendLatency()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_ResponseValidationError{
					field:  "BackendLatency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_ResponseValidationError{
					field:  "BackendLatency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBackendLatency()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_ResponseValidationError{
				field:  "BackendLatency",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeContext_ResponseMultiError(errors)
	}

	return nil
}

// AttributeContext_ResponseMultiError is an error wrapping multiple validation
// errors returned by AttributeContext_Response.ValidateAll() if the
// designated constraints aren't met.
type AttributeContext_ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContext_ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContext_ResponseMultiError) AllErrors() []error { return m }

// AttributeContext_ResponseValidationError is the validation error returned by
// AttributeContext_Response.Validate if the designated constraints aren't met.
type AttributeContext_ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_ResponseValidationError) ErrorName() string {
	return "AttributeContext_ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_ResponseValidationError{}

// Validate checks the field values on AttributeContext_Resource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeContext_Resource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeContext_Resource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeContext_ResourceMultiError, or nil if none found.
func (m *AttributeContext_Resource) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeContext_Resource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for Labels

	// no validation rules for Uid

	// no validation rules for Annotations

	// no validation rules for DisplayName

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_ResourceValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_ResourceValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_ResourceValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_ResourceValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_ResourceValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_ResourceValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDeleteTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeContext_ResourceValidationError{
					field:  "DeleteTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeContext_ResourceValidationError{
					field:  "DeleteTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeleteTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_ResourceValidationError{
				field:  "DeleteTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Etag

	// no validation rules for Location

	if len(errors) > 0 {
		return AttributeContext_ResourceMultiError(errors)
	}

	return nil
}

// AttributeContext_ResourceMultiError is an error wrapping multiple validation
// errors returned by AttributeContext_Resource.ValidateAll() if the
// designated constraints aren't met.
type AttributeContext_ResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeContext_ResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeContext_ResourceMultiError) AllErrors() []error { return m }

// AttributeContext_ResourceValidationError is the validation error returned by
// AttributeContext_Resource.Validate if the designated constraints aren't met.
type AttributeContext_ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_ResourceValidationError) ErrorName() string {
	return "AttributeContext_ResourceValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Resource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_ResourceValidationError{}
