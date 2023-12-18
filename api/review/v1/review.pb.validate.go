// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/review/v1/review.proto

package v1

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

// Validate checks the field values on CreateReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateReviewRequestMultiError, or nil if none found.
func (m *CreateReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserID() <= 0 {
		err := CreateReviewRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOrderID() <= 0 {
		err := CreateReviewRequestValidationError{
			field:  "OrderID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _CreateReviewRequest_Score_InLookup[m.GetScore()]; !ok {
		err := CreateReviewRequestValidationError{
			field:  "Score",
			reason: "value must be in list [1 2 3 4 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _CreateReviewRequest_ServiceScore_InLookup[m.GetServiceScore()]; !ok {
		err := CreateReviewRequestValidationError{
			field:  "ServiceScore",
			reason: "value must be in list [1 2 3 4 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _CreateReviewRequest_ExpressScore_InLookup[m.GetExpressScore()]; !ok {
		err := CreateReviewRequestValidationError{
			field:  "ExpressScore",
			reason: "value must be in list [1 2 3 4 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 8 || l > 255 {
		err := CreateReviewRequestValidationError{
			field:  "Content",
			reason: "value length must be between 8 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PicInfo

	// no validation rules for VideoInfo

	// no validation rules for Anonymous

	if len(errors) > 0 {
		return CreateReviewRequestMultiError(errors)
	}

	return nil
}

// CreateReviewRequestMultiError is an error wrapping multiple validation
// errors returned by CreateReviewRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateReviewRequestMultiError) AllErrors() []error { return m }

// CreateReviewRequestValidationError is the validation error returned by
// CreateReviewRequest.Validate if the designated constraints aren't met.
type CreateReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReviewRequestValidationError) ErrorName() string {
	return "CreateReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReviewRequestValidationError{}

var _CreateReviewRequest_Score_InLookup = map[int32]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

var _CreateReviewRequest_ServiceScore_InLookup = map[int32]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

var _CreateReviewRequest_ExpressScore_InLookup = map[int32]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

// Validate checks the field values on CreateReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateReviewReplyMultiError, or nil if none found.
func (m *CreateReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReviewID

	if len(errors) > 0 {
		return CreateReviewReplyMultiError(errors)
	}

	return nil
}

// CreateReviewReplyMultiError is an error wrapping multiple validation errors
// returned by CreateReviewReply.ValidateAll() if the designated constraints
// aren't met.
type CreateReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateReviewReplyMultiError) AllErrors() []error { return m }

// CreateReviewReplyValidationError is the validation error returned by
// CreateReviewReply.Validate if the designated constraints aren't met.
type CreateReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReviewReplyValidationError) ErrorName() string {
	return "CreateReviewReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReviewReplyValidationError{}

// Validate checks the field values on GetReviewRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReviewRequestMultiError, or nil if none found.
func (m *GetReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetReviewID() <= 0 {
		err := GetReviewRequestValidationError{
			field:  "ReviewID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetReviewRequestMultiError(errors)
	}

	return nil
}

// GetReviewRequestMultiError is an error wrapping multiple validation errors
// returned by GetReviewRequest.ValidateAll() if the designated constraints
// aren't met.
type GetReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReviewRequestMultiError) AllErrors() []error { return m }

// GetReviewRequestValidationError is the validation error returned by
// GetReviewRequest.Validate if the designated constraints aren't met.
type GetReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReviewRequestValidationError) ErrorName() string { return "GetReviewRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReviewRequestValidationError{}

// Validate checks the field values on GetReviewReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetReviewReplyMultiError,
// or nil if none found.
func (m *GetReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetReviewReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetReviewReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetReviewReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetReviewReplyMultiError(errors)
	}

	return nil
}

// GetReviewReplyMultiError is an error wrapping multiple validation errors
// returned by GetReviewReply.ValidateAll() if the designated constraints
// aren't met.
type GetReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReviewReplyMultiError) AllErrors() []error { return m }

// GetReviewReplyValidationError is the validation error returned by
// GetReviewReply.Validate if the designated constraints aren't met.
type GetReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReviewReplyValidationError) ErrorName() string { return "GetReviewReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReviewReplyValidationError{}

// Validate checks the field values on ReviewInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReviewInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReviewInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReviewInfoMultiError, or
// nil if none found.
func (m *ReviewInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ReviewInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReviewID

	// no validation rules for UserID

	// no validation rules for OrderID

	// no validation rules for Score

	// no validation rules for ServiceScore

	// no validation rules for ExpressScore

	// no validation rules for Content

	// no validation rules for PicInfo

	// no validation rules for VideoInfo

	// no validation rules for Status

	if len(errors) > 0 {
		return ReviewInfoMultiError(errors)
	}

	return nil
}

// ReviewInfoMultiError is an error wrapping multiple validation errors
// returned by ReviewInfo.ValidateAll() if the designated constraints aren't met.
type ReviewInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReviewInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReviewInfoMultiError) AllErrors() []error { return m }

// ReviewInfoValidationError is the validation error returned by
// ReviewInfo.Validate if the designated constraints aren't met.
type ReviewInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReviewInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReviewInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReviewInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReviewInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReviewInfoValidationError) ErrorName() string { return "ReviewInfoValidationError" }

// Error satisfies the builtin error interface
func (e ReviewInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReviewInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReviewInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReviewInfoValidationError{}

// Validate checks the field values on ReplyReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReplyReviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReplyReviewRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReplyReviewRequestMultiError, or nil if none found.
func (m *ReplyReviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReplyReviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetReviewID() <= 0 {
		err := ReplyReviewRequestValidationError{
			field:  "ReviewID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStoreID() <= 0 {
		err := ReplyReviewRequestValidationError{
			field:  "StoreID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 2 || l > 200 {
		err := ReplyReviewRequestValidationError{
			field:  "Content",
			reason: "value length must be between 2 and 200 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PicInfo

	// no validation rules for VideoInfo

	if len(errors) > 0 {
		return ReplyReviewRequestMultiError(errors)
	}

	return nil
}

// ReplyReviewRequestMultiError is an error wrapping multiple validation errors
// returned by ReplyReviewRequest.ValidateAll() if the designated constraints
// aren't met.
type ReplyReviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReplyReviewRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReplyReviewRequestMultiError) AllErrors() []error { return m }

// ReplyReviewRequestValidationError is the validation error returned by
// ReplyReviewRequest.Validate if the designated constraints aren't met.
type ReplyReviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReplyReviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReplyReviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReplyReviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReplyReviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReplyReviewRequestValidationError) ErrorName() string {
	return "ReplyReviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReplyReviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReplyReviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReplyReviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReplyReviewRequestValidationError{}

// Validate checks the field values on ReplyReviewReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReplyReviewReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReplyReviewReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReplyReviewReplyMultiError, or nil if none found.
func (m *ReplyReviewReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ReplyReviewReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReplyID

	if len(errors) > 0 {
		return ReplyReviewReplyMultiError(errors)
	}

	return nil
}

// ReplyReviewReplyMultiError is an error wrapping multiple validation errors
// returned by ReplyReviewReply.ValidateAll() if the designated constraints
// aren't met.
type ReplyReviewReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReplyReviewReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReplyReviewReplyMultiError) AllErrors() []error { return m }

// ReplyReviewReplyValidationError is the validation error returned by
// ReplyReviewReply.Validate if the designated constraints aren't met.
type ReplyReviewReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReplyReviewReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReplyReviewReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReplyReviewReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReplyReviewReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReplyReviewReplyValidationError) ErrorName() string { return "ReplyReviewReplyValidationError" }

// Error satisfies the builtin error interface
func (e ReplyReviewReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReplyReviewReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReplyReviewReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReplyReviewReplyValidationError{}

// Validate checks the field values on ListReviewByUserIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListReviewByUserIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReviewByUserIDRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReviewByUserIDRequestMultiError, or nil if none found.
func (m *ListReviewByUserIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReviewByUserIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserID() <= 0 {
		err := ListReviewByUserIDRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListReviewByUserIDRequestMultiError(errors)
	}

	return nil
}

// ListReviewByUserIDRequestMultiError is an error wrapping multiple validation
// errors returned by ListReviewByUserIDRequest.ValidateAll() if the
// designated constraints aren't met.
type ListReviewByUserIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReviewByUserIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReviewByUserIDRequestMultiError) AllErrors() []error { return m }

// ListReviewByUserIDRequestValidationError is the validation error returned by
// ListReviewByUserIDRequest.Validate if the designated constraints aren't met.
type ListReviewByUserIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReviewByUserIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReviewByUserIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReviewByUserIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReviewByUserIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReviewByUserIDRequestValidationError) ErrorName() string {
	return "ListReviewByUserIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListReviewByUserIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReviewByUserIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReviewByUserIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReviewByUserIDRequestValidationError{}

// Validate checks the field values on ListReviewByUserIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListReviewByUserIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReviewByUserIDReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReviewByUserIDReplyMultiError, or nil if none found.
func (m *ListReviewByUserIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReviewByUserIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListReviewByUserIDReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReviewByUserIDReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReviewByUserIDReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListReviewByUserIDReplyMultiError(errors)
	}

	return nil
}

// ListReviewByUserIDReplyMultiError is an error wrapping multiple validation
// errors returned by ListReviewByUserIDReply.ValidateAll() if the designated
// constraints aren't met.
type ListReviewByUserIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReviewByUserIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReviewByUserIDReplyMultiError) AllErrors() []error { return m }

// ListReviewByUserIDReplyValidationError is the validation error returned by
// ListReviewByUserIDReply.Validate if the designated constraints aren't met.
type ListReviewByUserIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReviewByUserIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReviewByUserIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReviewByUserIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReviewByUserIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReviewByUserIDReplyValidationError) ErrorName() string {
	return "ListReviewByUserIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListReviewByUserIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReviewByUserIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReviewByUserIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReviewByUserIDReplyValidationError{}
