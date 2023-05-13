// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gtool.proto

package pb

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

// Validate checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServerMultiError, or nil if none found.
func (m *Server) ValidateAll() error {
	return m.validate(true)
}

func (m *Server) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Key != nil {
		// no validation rules for Key
	}

	if m.Url != nil {
		// no validation rules for Url
	}

	if len(errors) > 0 {
		return ServerMultiError(errors)
	}

	return nil
}

// ServerMultiError is an error wrapping multiple validation errors returned by
// Server.ValidateAll() if the designated constraints aren't met.
type ServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerMultiError) AllErrors() []error { return m }

// ServerValidationError is the validation error returned by Server.Validate if
// the designated constraints aren't met.
type ServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerValidationError) ErrorName() string { return "ServerValidationError" }

// Error satisfies the builtin error interface
func (e ServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Service) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServiceMultiError, or nil if none found.
func (m *Service) ValidateAll() error {
	return m.validate(true)
}

func (m *Service) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMethods() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServiceValidationError{
						field:  fmt.Sprintf("Methods[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServiceValidationError{
						field:  fmt.Sprintf("Methods[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceValidationError{
					field:  fmt.Sprintf("Methods[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Key != nil {
		// no validation rules for Key
	}

	if m.ServiceName != nil {
		// no validation rules for ServiceName
	}

	if len(errors) > 0 {
		return ServiceMultiError(errors)
	}

	return nil
}

// ServiceMultiError is an error wrapping multiple validation errors returned
// by Service.ValidateAll() if the designated constraints aren't met.
type ServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceMultiError) AllErrors() []error { return m }

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on Method with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Method) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Method with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MethodMultiError, or nil if none found.
func (m *Method) ValidateAll() error {
	return m.validate(true)
}

func (m *Method) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InputType

	// no validation rules for OutputType

	if m.Key != nil {
		// no validation rules for Key
	}

	if m.MethodName != nil {
		// no validation rules for MethodName
	}

	if m.MethodType != nil {
		// no validation rules for MethodType
	}

	if len(errors) > 0 {
		return MethodMultiError(errors)
	}

	return nil
}

// MethodMultiError is an error wrapping multiple validation errors returned by
// Method.ValidateAll() if the designated constraints aren't met.
type MethodMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MethodMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MethodMultiError) AllErrors() []error { return m }

// MethodValidationError is the validation error returned by Method.Validate if
// the designated constraints aren't met.
type MethodValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MethodValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MethodValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MethodValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MethodValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MethodValidationError) ErrorName() string { return "MethodValidationError" }

// Error satisfies the builtin error interface
func (e MethodValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMethod.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MethodValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MethodValidationError{}

// Validate checks the field values on ServerInfoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServerInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerInfoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServerInfoReqMultiError, or
// nil if none found.
func (m *ServerInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Url != nil {

		if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 100 {
			err := ServerInfoReqValidationError{
				field:  "Url",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ServerInfoReqMultiError(errors)
	}

	return nil
}

// ServerInfoReqMultiError is an error wrapping multiple validation errors
// returned by ServerInfoReq.ValidateAll() if the designated constraints
// aren't met.
type ServerInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerInfoReqMultiError) AllErrors() []error { return m }

// ServerInfoReqValidationError is the validation error returned by
// ServerInfoReq.Validate if the designated constraints aren't met.
type ServerInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoReqValidationError) ErrorName() string { return "ServerInfoReqValidationError" }

// Error satisfies the builtin error interface
func (e ServerInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoReqValidationError{}

// Validate checks the field values on ServerInfoRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServerInfoRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerInfoRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServerInfoRspMultiError, or
// nil if none found.
func (m *ServerInfoRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerInfoRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Code != nil {
		// no validation rules for Code
	}

	if m.Message != nil {
		// no validation rules for Message
	}

	if m.Data != nil {

		if all {
			switch v := interface{}(m.GetData()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerInfoRspValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerInfoRspValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerInfoRspValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ServerInfoRspMultiError(errors)
	}

	return nil
}

// ServerInfoRspMultiError is an error wrapping multiple validation errors
// returned by ServerInfoRsp.ValidateAll() if the designated constraints
// aren't met.
type ServerInfoRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerInfoRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerInfoRspMultiError) AllErrors() []error { return m }

// ServerInfoRspValidationError is the validation error returned by
// ServerInfoRsp.Validate if the designated constraints aren't met.
type ServerInfoRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoRspValidationError) ErrorName() string { return "ServerInfoRspValidationError" }

// Error satisfies the builtin error interface
func (e ServerInfoRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfoRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoRspValidationError{}

// Validate checks the field values on MethodParamReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MethodParamReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MethodParamReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MethodParamReqMultiError,
// or nil if none found.
func (m *MethodParamReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MethodParamReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Url != nil {

		if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 100 {
			err := MethodParamReqValidationError{
				field:  "Url",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.ServiceName != nil {

		if l := utf8.RuneCountInString(m.GetServiceName()); l < 1 || l > 100 {
			err := MethodParamReqValidationError{
				field:  "ServiceName",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.MethodName != nil {

		if l := utf8.RuneCountInString(m.GetMethodName()); l < 1 || l > 100 {
			err := MethodParamReqValidationError{
				field:  "MethodName",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return MethodParamReqMultiError(errors)
	}

	return nil
}

// MethodParamReqMultiError is an error wrapping multiple validation errors
// returned by MethodParamReq.ValidateAll() if the designated constraints
// aren't met.
type MethodParamReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MethodParamReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MethodParamReqMultiError) AllErrors() []error { return m }

// MethodParamReqValidationError is the validation error returned by
// MethodParamReq.Validate if the designated constraints aren't met.
type MethodParamReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MethodParamReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MethodParamReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MethodParamReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MethodParamReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MethodParamReqValidationError) ErrorName() string { return "MethodParamReqValidationError" }

// Error satisfies the builtin error interface
func (e MethodParamReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMethodParamReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MethodParamReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MethodParamReqValidationError{}

// Validate checks the field values on MethodParamRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MethodParamRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MethodParamRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MethodParamRspMultiError,
// or nil if none found.
func (m *MethodParamRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *MethodParamRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Code != nil {
		// no validation rules for Code
	}

	if m.Message != nil {
		// no validation rules for Message
	}

	if m.Data != nil {
		// no validation rules for Data
	}

	if len(errors) > 0 {
		return MethodParamRspMultiError(errors)
	}

	return nil
}

// MethodParamRspMultiError is an error wrapping multiple validation errors
// returned by MethodParamRsp.ValidateAll() if the designated constraints
// aren't met.
type MethodParamRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MethodParamRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MethodParamRspMultiError) AllErrors() []error { return m }

// MethodParamRspValidationError is the validation error returned by
// MethodParamRsp.Validate if the designated constraints aren't met.
type MethodParamRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MethodParamRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MethodParamRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MethodParamRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MethodParamRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MethodParamRspValidationError) ErrorName() string { return "MethodParamRspValidationError" }

// Error satisfies the builtin error interface
func (e MethodParamRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMethodParamRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MethodParamRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MethodParamRspValidationError{}

// Validate checks the field values on CallMethodReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CallMethodReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallMethodReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CallMethodReqMultiError, or
// nil if none found.
func (m *CallMethodReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CallMethodReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Url != nil {
		// no validation rules for Url
	}

	if m.ServiceName != nil {
		// no validation rules for ServiceName
	}

	if m.MethodName != nil {
		// no validation rules for MethodName
	}

	if m.Data != nil {
		// no validation rules for Data
	}

	if len(errors) > 0 {
		return CallMethodReqMultiError(errors)
	}

	return nil
}

// CallMethodReqMultiError is an error wrapping multiple validation errors
// returned by CallMethodReq.ValidateAll() if the designated constraints
// aren't met.
type CallMethodReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallMethodReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallMethodReqMultiError) AllErrors() []error { return m }

// CallMethodReqValidationError is the validation error returned by
// CallMethodReq.Validate if the designated constraints aren't met.
type CallMethodReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallMethodReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallMethodReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallMethodReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallMethodReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallMethodReqValidationError) ErrorName() string { return "CallMethodReqValidationError" }

// Error satisfies the builtin error interface
func (e CallMethodReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallMethodReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallMethodReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallMethodReqValidationError{}

// Validate checks the field values on CallMethodRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CallMethodRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallMethodRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CallMethodRspMultiError, or
// nil if none found.
func (m *CallMethodRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *CallMethodRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Code != nil {
		// no validation rules for Code
	}

	if m.Message != nil {
		// no validation rules for Message
	}

	if m.Data != nil {
		// no validation rules for Data
	}

	if len(errors) > 0 {
		return CallMethodRspMultiError(errors)
	}

	return nil
}

// CallMethodRspMultiError is an error wrapping multiple validation errors
// returned by CallMethodRsp.ValidateAll() if the designated constraints
// aren't met.
type CallMethodRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallMethodRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallMethodRspMultiError) AllErrors() []error { return m }

// CallMethodRspValidationError is the validation error returned by
// CallMethodRsp.Validate if the designated constraints aren't met.
type CallMethodRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallMethodRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallMethodRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallMethodRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallMethodRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallMethodRspValidationError) ErrorName() string { return "CallMethodRspValidationError" }

// Error satisfies the builtin error interface
func (e CallMethodRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallMethodRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallMethodRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallMethodRspValidationError{}

// Validate checks the field values on CallServerStreamMethodRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CallServerStreamMethodRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallServerStreamMethodRsp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CallServerStreamMethodRspMultiError, or nil if none found.
func (m *CallServerStreamMethodRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *CallServerStreamMethodRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Code != nil {
		// no validation rules for Code
	}

	if m.Message != nil {
		// no validation rules for Message
	}

	if len(errors) > 0 {
		return CallServerStreamMethodRspMultiError(errors)
	}

	return nil
}

// CallServerStreamMethodRspMultiError is an error wrapping multiple validation
// errors returned by CallServerStreamMethodRsp.ValidateAll() if the
// designated constraints aren't met.
type CallServerStreamMethodRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallServerStreamMethodRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallServerStreamMethodRspMultiError) AllErrors() []error { return m }

// CallServerStreamMethodRspValidationError is the validation error returned by
// CallServerStreamMethodRsp.Validate if the designated constraints aren't met.
type CallServerStreamMethodRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallServerStreamMethodRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallServerStreamMethodRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallServerStreamMethodRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallServerStreamMethodRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallServerStreamMethodRspValidationError) ErrorName() string {
	return "CallServerStreamMethodRspValidationError"
}

// Error satisfies the builtin error interface
func (e CallServerStreamMethodRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallServerStreamMethodRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallServerStreamMethodRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallServerStreamMethodRspValidationError{}

// Validate checks the field values on CallClientStreamMethodReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CallClientStreamMethodReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallClientStreamMethodReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CallClientStreamMethodReqMultiError, or nil if none found.
func (m *CallClientStreamMethodReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CallClientStreamMethodReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Url != nil {
		// no validation rules for Url
	}

	if m.ServiceName != nil {
		// no validation rules for ServiceName
	}

	if m.MethodName != nil {
		// no validation rules for MethodName
	}

	if len(errors) > 0 {
		return CallClientStreamMethodReqMultiError(errors)
	}

	return nil
}

// CallClientStreamMethodReqMultiError is an error wrapping multiple validation
// errors returned by CallClientStreamMethodReq.ValidateAll() if the
// designated constraints aren't met.
type CallClientStreamMethodReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallClientStreamMethodReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallClientStreamMethodReqMultiError) AllErrors() []error { return m }

// CallClientStreamMethodReqValidationError is the validation error returned by
// CallClientStreamMethodReq.Validate if the designated constraints aren't met.
type CallClientStreamMethodReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallClientStreamMethodReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallClientStreamMethodReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallClientStreamMethodReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallClientStreamMethodReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallClientStreamMethodReqValidationError) ErrorName() string {
	return "CallClientStreamMethodReqValidationError"
}

// Error satisfies the builtin error interface
func (e CallClientStreamMethodReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallClientStreamMethodReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallClientStreamMethodReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallClientStreamMethodReqValidationError{}