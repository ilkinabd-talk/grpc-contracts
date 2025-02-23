// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: google/api/log.proto

package serviceconfig

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

// Validate checks the field values on LogDescriptor with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogDescriptor) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogDescriptor with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogDescriptorMultiError, or
// nil if none found.
func (m *LogDescriptor) ValidateAll() error {
	return m.validate(true)
}

func (m *LogDescriptor) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	for idx, item := range m.GetLabels() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LogDescriptorValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LogDescriptorValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LogDescriptorValidationError{
					field:  fmt.Sprintf("Labels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Description

	// no validation rules for DisplayName

	if len(errors) > 0 {
		return LogDescriptorMultiError(errors)
	}

	return nil
}

// LogDescriptorMultiError is an error wrapping multiple validation errors
// returned by LogDescriptor.ValidateAll() if the designated constraints
// aren't met.
type LogDescriptorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogDescriptorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogDescriptorMultiError) AllErrors() []error { return m }

// LogDescriptorValidationError is the validation error returned by
// LogDescriptor.Validate if the designated constraints aren't met.
type LogDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogDescriptorValidationError) ErrorName() string { return "LogDescriptorValidationError" }

// Error satisfies the builtin error interface
func (e LogDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogDescriptorValidationError{}
