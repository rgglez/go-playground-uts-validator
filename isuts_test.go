/*
   Copyright 2026 Rodolfo González González

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package isuts

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
)

//-----------------------------------------------------------------------------
// Helpers
//-----------------------------------------------------------------------------

func newValidator(t *testing.T, tag string) *validator.Validate {
	t.Helper()
	v := validator.New()
	if err := RegisterUTSValidator(v, tag); err != nil {
		t.Fatalf("RegisterUTSValidator: %v", err)
	}
	return v
}

//-----------------------------------------------------------------------------
// RegisterUTSValidator
//-----------------------------------------------------------------------------

func TestRegisterUTSValidator_DefaultTag(t *testing.T) {
	v := validator.New()
	if err := RegisterUTSValidator(v, ""); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestRegisterUTSValidator_CustomTag(t *testing.T) {
	v := validator.New()
	if err := RegisterUTSValidator(v, "mytag"); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

//-----------------------------------------------------------------------------
// ValidateUnixTimestamp — string fields
//-----------------------------------------------------------------------------

func TestValidateUnixTimestamp_StringValid(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS string `validate:"uts"`
	}
	cases := []string{"0", "1", "1700000000", "-1", "9223372036854775807"}
	for _, c := range cases {
		if err := v.Struct(S{TS: c}); err != nil {
			t.Errorf("expected valid for %q, got: %v", c, err)
		}
	}
}

func TestValidateUnixTimestamp_StringInvalid(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS string `validate:"uts"`
	}
	cases := []string{"", "abc", "1.5", "2006-01-02", " 123"}
	for _, c := range cases {
		if err := v.Struct(S{TS: c}); err == nil {
			t.Errorf("expected invalid for %q, but validation passed", c)
		}
	}
}

//-----------------------------------------------------------------------------
// ValidateUnixTimestamp — integer fields
//-----------------------------------------------------------------------------

func TestValidateUnixTimestamp_IntTypes(t *testing.T) {
	v := newValidator(t, "")

	type SI struct {
		TS int `validate:"uts"`
	}
	type SI8 struct {
		TS int8 `validate:"uts"`
	}
	type SI16 struct {
		TS int16 `validate:"uts"`
	}
	type SI32 struct {
		TS int32 `validate:"uts"`
	}
	type SI64 struct {
		TS int64 `validate:"uts"`
	}

	if err := v.Struct(SI{TS: 1700000000}); err != nil {
		t.Errorf("int: %v", err)
	}
	if err := v.Struct(SI8{TS: 127}); err != nil {
		t.Errorf("int8: %v", err)
	}
	if err := v.Struct(SI16{TS: 32767}); err != nil {
		t.Errorf("int16: %v", err)
	}
	if err := v.Struct(SI32{TS: 2147483647}); err != nil {
		t.Errorf("int32: %v", err)
	}
	if err := v.Struct(SI64{TS: 1700000000}); err != nil {
		t.Errorf("int64: %v", err)
	}
}

func TestValidateUnixTimestamp_UintTypes(t *testing.T) {
	v := newValidator(t, "")

	type SU struct {
		TS uint `validate:"uts"`
	}
	type SU8 struct {
		TS uint8 `validate:"uts"`
	}
	type SU16 struct {
		TS uint16 `validate:"uts"`
	}
	type SU32 struct {
		TS uint32 `validate:"uts"`
	}
	type SU64 struct {
		TS uint64 `validate:"uts"`
	}

	if err := v.Struct(SU{TS: 1700000000}); err != nil {
		t.Errorf("uint: %v", err)
	}
	if err := v.Struct(SU8{TS: 255}); err != nil {
		t.Errorf("uint8: %v", err)
	}
	if err := v.Struct(SU16{TS: 65535}); err != nil {
		t.Errorf("uint16: %v", err)
	}
	if err := v.Struct(SU32{TS: 4294967295}); err != nil {
		t.Errorf("uint32: %v", err)
	}
	if err := v.Struct(SU64{TS: 1700000000}); err != nil {
		t.Errorf("uint64: %v", err)
	}
}

//-----------------------------------------------------------------------------
// ValidateUnixTimestamp — time.Time fields
//-----------------------------------------------------------------------------

func TestValidateUnixTimestamp_TimeValid(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS time.Time `validate:"uts"`
	}
	cases := []time.Time{
		time.Now(),
		time.Unix(1700000000, 0),
		time.Unix(0, 1), // non-zero nanosecond
	}
	for _, c := range cases {
		if err := v.Struct(S{TS: c}); err != nil {
			t.Errorf("expected valid for %v, got: %v", c, err)
		}
	}
}

func TestValidateUnixTimestamp_TimeZero(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS time.Time `validate:"uts"`
	}
	if err := v.Struct(S{TS: time.Time{}}); err == nil {
		t.Error("expected invalid for zero time.Time, but validation passed")
	}
}

//-----------------------------------------------------------------------------
// ValidateUnixTimestamp — pointer fields
//-----------------------------------------------------------------------------

func TestValidateUnixTimestamp_PointerNil(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS *int64 `validate:"uts"`
	}
	if err := v.Struct(S{TS: nil}); err == nil {
		t.Error("expected invalid for nil pointer, but validation passed")
	}
}

func TestValidateUnixTimestamp_PointerNonNil(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS *int64 `validate:"uts"`
	}
	val := int64(1700000000)
	if err := v.Struct(S{TS: &val}); err != nil {
		t.Errorf("expected valid for non-nil int64 pointer, got: %v", err)
	}
}

//-----------------------------------------------------------------------------
// ValidateUnixTimestamp — unsupported kinds
//-----------------------------------------------------------------------------

func TestValidateUnixTimestamp_FloatInvalid(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS float64 `validate:"uts"`
	}
	if err := v.Struct(S{TS: 1700000000.0}); err == nil {
		t.Error("expected invalid for float64, but validation passed")
	}
}

func TestValidateUnixTimestamp_BoolInvalid(t *testing.T) {
	v := newValidator(t, "")
	type S struct {
		TS bool `validate:"uts"`
	}
	if err := v.Struct(S{TS: true}); err == nil {
		t.Error("expected invalid for bool, but validation passed")
	}
}
