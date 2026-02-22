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
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

//-----------------------------------------------------------------------------

const DefaultTagName = "uts"

//-----------------------------------------------------------------------------

var timeType = reflect.TypeOf(time.Time{})

//-----------------------------------------------------------------------------

// ValidateUnixTimestamp checks whether the field value is a valid Unix timestamp.
// It supports time.Time(), string and integer kinds.
func ValidateUnixTimestamp(fl validator.FieldLevel) bool {
	field := fl.Field()
	if !field.IsValid() {
		return false
	}

	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			return false
		}
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.String:
		value := field.String()
		if value == "" {
			return false
		}
		_, err := strconv.ParseInt(value, 10, 64)
		return err == nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Struct:
		if field.Type() == timeType {
			t, ok := field.Interface().(time.Time)
			return ok && !t.IsZero()
		}
		return false
	default:
		return false
	}
}

//-----------------------------------------------------------------------------

// RegisterUTSValidator registers the Unix timestamp validator on the given tag.
// If fieldTag is empty, DefaultTagName is used.
func RegisterUTSValidator(v *validator.Validate, fieldTag string) error {
	if fieldTag == "" {
		fieldTag = DefaultTagName
	}
	return v.RegisterValidation(fieldTag, ValidateUnixTimestamp)
}
