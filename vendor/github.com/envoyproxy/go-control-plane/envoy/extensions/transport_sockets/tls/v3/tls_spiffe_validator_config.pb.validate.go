// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tls/v3/tls_spiffe_validator_config.proto

package tlsv3

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

// Validate checks the field values on SPIFFECertValidatorConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SPIFFECertValidatorConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SPIFFECertValidatorConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SPIFFECertValidatorConfigMultiError, or nil if none found.
func (m *SPIFFECertValidatorConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SPIFFECertValidatorConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTrustDomains()) < 1 {
		err := SPIFFECertValidatorConfigValidationError{
			field:  "TrustDomains",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetTrustDomains() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SPIFFECertValidatorConfigValidationError{
						field:  fmt.Sprintf("TrustDomains[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SPIFFECertValidatorConfigValidationError{
						field:  fmt.Sprintf("TrustDomains[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SPIFFECertValidatorConfigValidationError{
					field:  fmt.Sprintf("TrustDomains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SPIFFECertValidatorConfigMultiError(errors)
	}

	return nil
}

// SPIFFECertValidatorConfigMultiError is an error wrapping multiple validation
// errors returned by SPIFFECertValidatorConfig.ValidateAll() if the
// designated constraints aren't met.
type SPIFFECertValidatorConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SPIFFECertValidatorConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SPIFFECertValidatorConfigMultiError) AllErrors() []error { return m }

// SPIFFECertValidatorConfigValidationError is the validation error returned by
// SPIFFECertValidatorConfig.Validate if the designated constraints aren't met.
type SPIFFECertValidatorConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SPIFFECertValidatorConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SPIFFECertValidatorConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SPIFFECertValidatorConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SPIFFECertValidatorConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SPIFFECertValidatorConfigValidationError) ErrorName() string {
	return "SPIFFECertValidatorConfigValidationError"
}

// Error satisfies the builtin error interface
func (e SPIFFECertValidatorConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSPIFFECertValidatorConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SPIFFECertValidatorConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SPIFFECertValidatorConfigValidationError{}

// Validate checks the field values on SPIFFECertValidatorConfig_TrustDomain
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *SPIFFECertValidatorConfig_TrustDomain) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SPIFFECertValidatorConfig_TrustDomain
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// SPIFFECertValidatorConfig_TrustDomainMultiError, or nil if none found.
func (m *SPIFFECertValidatorConfig_TrustDomain) ValidateAll() error {
	return m.validate(true)
}

func (m *SPIFFECertValidatorConfig_TrustDomain) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := SPIFFECertValidatorConfig_TrustDomainValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTrustBundle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SPIFFECertValidatorConfig_TrustDomainValidationError{
					field:  "TrustBundle",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SPIFFECertValidatorConfig_TrustDomainValidationError{
					field:  "TrustBundle",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTrustBundle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SPIFFECertValidatorConfig_TrustDomainValidationError{
				field:  "TrustBundle",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SPIFFECertValidatorConfig_TrustDomainMultiError(errors)
	}

	return nil
}

// SPIFFECertValidatorConfig_TrustDomainMultiError is an error wrapping
// multiple validation errors returned by
// SPIFFECertValidatorConfig_TrustDomain.ValidateAll() if the designated
// constraints aren't met.
type SPIFFECertValidatorConfig_TrustDomainMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SPIFFECertValidatorConfig_TrustDomainMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SPIFFECertValidatorConfig_TrustDomainMultiError) AllErrors() []error { return m }

// SPIFFECertValidatorConfig_TrustDomainValidationError is the validation error
// returned by SPIFFECertValidatorConfig_TrustDomain.Validate if the
// designated constraints aren't met.
type SPIFFECertValidatorConfig_TrustDomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) ErrorName() string {
	return "SPIFFECertValidatorConfig_TrustDomainValidationError"
}

// Error satisfies the builtin error interface
func (e SPIFFECertValidatorConfig_TrustDomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSPIFFECertValidatorConfig_TrustDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SPIFFECertValidatorConfig_TrustDomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SPIFFECertValidatorConfig_TrustDomainValidationError{}