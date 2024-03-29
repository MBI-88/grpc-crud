// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/crud.proto

package user

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

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Dni

	// no validation rules for Phone

	// no validation rules for Website

	// no validation rules for Address

	// no validation rules for Password

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_CreateUserRequest_Name_Pattern.MatchString(m.GetName()) {
		err := CreateUserRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9-_\\\\s]{3,10}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateUserRequest_Dni_Pattern.MatchString(m.GetDni()) {
		err := CreateUserRequestValidationError{
			field:  "Dni",
			reason: "value does not match regex pattern \"^\\\\d{8}[A-Za-z]$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateUserRequest_Phone_Pattern.MatchString(m.GetPhone()) {
		err := CreateUserRequestValidationError{
			field:  "Phone",
			reason: "value does not match regex pattern \"^\\\\+598\\\\d{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateUserRequest_Website_Pattern.MatchString(m.GetWebsite()) {
		err := CreateUserRequestValidationError{
			field:  "Website",
			reason: "value does not match regex pattern \"^(https?://)[\\\\w\\\\-]+(\\\\.[\\\\w\\\\-]+)+[/#?]?.*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAddress()) > 100 {
		err := CreateUserRequestValidationError{
			field:  "Address",
			reason: "value length must be at most 100 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateUserRequest_Password_Pattern.MatchString(m.GetPassword()) {
		err := CreateUserRequestValidationError{
			field:  "Password",
			reason: "value does not match regex pattern \"^[A-Za-z\\\\d\\\\W]{8,12}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

var _CreateUserRequest_Name_Pattern = regexp.MustCompile("^[a-zA-Z0-9-_\\s]{3,10}$")

var _CreateUserRequest_Dni_Pattern = regexp.MustCompile("^\\d{8}[A-Za-z]$")

var _CreateUserRequest_Phone_Pattern = regexp.MustCompile("^\\+598\\d{8}$")

var _CreateUserRequest_Website_Pattern = regexp.MustCompile("^(https?://)[\\w\\-]+(\\.[\\w\\-]+)+[/#?]?.*$")

var _CreateUserRequest_Password_Pattern = regexp.MustCompile("^[A-Za-z\\d\\W]{8,12}$")

// Validate checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserResponseMultiError, or nil if none found.
func (m *CreateUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return CreateUserResponseMultiError(errors)
	}

	return nil
}

// CreateUserResponseMultiError is an error wrapping multiple validation errors
// returned by CreateUserResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserResponseMultiError) AllErrors() []error { return m }

// CreateUserResponseValidationError is the validation error returned by
// CreateUserResponse.Validate if the designated constraints aren't met.
type CreateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResponseValidationError) ErrorName() string {
	return "CreateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResponseValidationError{}

// Validate checks the field values on ListUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserRequestMultiError, or nil if none found.
func (m *ListUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUserRequestMultiError(errors)
	}

	return nil
}

// ListUserRequestMultiError is an error wrapping multiple validation errors
// returned by ListUserRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserRequestMultiError) AllErrors() []error { return m }

// ListUserRequestValidationError is the validation error returned by
// ListUserRequest.Validate if the designated constraints aren't met.
type ListUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserRequestValidationError) ErrorName() string { return "ListUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserRequestValidationError{}

// Validate checks the field values on ListUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserResponseMultiError, or nil if none found.
func (m *ListUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUserResponseMultiError(errors)
	}

	return nil
}

// ListUserResponseMultiError is an error wrapping multiple validation errors
// returned by ListUserResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserResponseMultiError) AllErrors() []error { return m }

// ListUserResponseValidationError is the validation error returned by
// ListUserResponse.Validate if the designated constraints aren't met.
type ListUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserResponseValidationError) ErrorName() string { return "ListUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserResponseValidationError{}

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRequestMultiError, or nil if none found.
func (m *UpdateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateUserRequestMultiError(errors)
	}

	return nil
}

// UpdateUserRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRequestMultiError) AllErrors() []error { return m }

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

// Validate checks the field values on UpdateUserRespose with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRespose) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRespose with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserResposeMultiError, or nil if none found.
func (m *UpdateUserRespose) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRespose) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return UpdateUserResposeMultiError(errors)
	}

	return nil
}

// UpdateUserResposeMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRespose.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserResposeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserResposeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserResposeMultiError) AllErrors() []error { return m }

// UpdateUserResposeValidationError is the validation error returned by
// UpdateUserRespose.Validate if the designated constraints aren't met.
type UpdateUserResposeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserResposeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserResposeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserResposeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserResposeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserResposeValidationError) ErrorName() string {
	return "UpdateUserResposeValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserResposeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRespose.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserResposeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserResposeValidationError{}

// Validate checks the field values on DeleteUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserRequestMultiError, or nil if none found.
func (m *DeleteUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteUserRequestMultiError(errors)
	}

	return nil
}

// DeleteUserRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteUserRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserRequestMultiError) AllErrors() []error { return m }

// DeleteUserRequestValidationError is the validation error returned by
// DeleteUserRequest.Validate if the designated constraints aren't met.
type DeleteUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserRequestValidationError) ErrorName() string {
	return "DeleteUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserRequestValidationError{}

// Validate checks the field values on DeleteUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserResponseMultiError, or nil if none found.
func (m *DeleteUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteUserResponseMultiError(errors)
	}

	return nil
}

// DeleteUserResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteUserResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserResponseMultiError) AllErrors() []error { return m }

// DeleteUserResponseValidationError is the validation error returned by
// DeleteUserResponse.Validate if the designated constraints aren't met.
type DeleteUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserResponseValidationError) ErrorName() string {
	return "DeleteUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserResponseValidationError{}
