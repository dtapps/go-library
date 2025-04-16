package framework

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Validator struct {
	once        sync.Once
	validate    *validator.Validate
	validateTag string
}

func NewValidator() *Validator {
	vd := &Validator{}
	vd.lazyinit()
	return vd
}

type SliceValidationError []error

// Error concatenates all error elements in SliceValidationError into a single string separated by \n.
func (err SliceValidationError) Error() string {
	n := len(err)
	switch n {
	case 0:
		return ""
	default:
		var b strings.Builder
		if err[0] != nil {
			fmt.Fprintf(&b, "[%d]: %s", 0, err[0].Error())
		}
		if n > 1 {
			for i := 1; i < n; i++ {
				if err[i] != nil {
					b.WriteString("\n")
					fmt.Fprintf(&b, "[%d]: %s", i, err[i].Error())
				}
			}
		}
		return b.String()
	}
}

func (m *Validator) ValidateStruct(obj interface{}) error {
	if obj == nil {
		return nil
	}
	value := reflect.Value{}
	if val, ok := obj.(reflect.Value); ok {
		value = val
		obj = val.Interface()
	} else {
		value = reflect.ValueOf(obj)
	}

	// Handle default values before validation
	if err := m.setDefaults(value); err != nil {
		return err
	}

	switch value.Kind() {
	case reflect.Ptr:
		return m.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return m.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(SliceValidationError, 0)
		for i := 0; i < count; i++ {
			if err := m.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}

// setDefaults sets default values for fields with the "default" tag
func (m *Validator) setDefaults(value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil
	}

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := t.Field(i)

		// Get the tag and split it into parts
		tag := fieldType.Tag.Get(m.ValidateTag())
		if tag == "" {
			continue
		}

		// Parse the "default" value from the tag
		defaultValue := parseDefaultValue(tag)
		if defaultValue == "" {
			continue
		}

		// Set the default value if the field is empty
		if field.CanSet() && isEmptyValue(field) {
			if err := setFieldValue(field, defaultValue); err != nil {
				return fmt.Errorf("failed to set default value for field '%s': %w", fieldType.Name, err)
			}
		}
	}
	return nil
}

// parseDefaultValue extracts the "default" value from the tag
func parseDefaultValue(tag string) string {
	parts := strings.Split(tag, ",")
	for _, part := range parts {
		if strings.HasPrefix(part, "default=") {
			return strings.TrimPrefix(part, "default=")
		}
	}
	return ""
}

// isEmptyValue checks if a field is empty (zero value)
func isEmptyValue(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		return field.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return field.Float() == 0
	case reflect.Bool:
		return !field.Bool()
	case reflect.Slice, reflect.Array, reflect.Map:
		return field.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return field.IsNil()
	default:
		return false
	}
}

// setFieldValue sets the value of a field based on its type
func setFieldValue(field reflect.Value, defaultValue string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(defaultValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(defaultValue, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(defaultValue, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(val)
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(defaultValue, 64)
		if err != nil {
			return err
		}
		field.SetFloat(val)
	case reflect.Bool:
		val, err := strconv.ParseBool(defaultValue)
		if err != nil {
			return err
		}
		field.SetBool(val)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
	return nil
}

// validateStruct receives struct type
func (m *Validator) validateStruct(obj interface{}) error {
	m.lazyinit()
	err := m.validate.Struct(obj)
	if err == nil {
		return nil
	}

	// Handle validation errors
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return CustomValidationError(ve)
	}
	return err
}

// CustomValidationError formats validation errors into a custom format
func CustomValidationError(errs validator.ValidationErrors) error {
	var errMsgs []string
	for _, e := range errs {
		field := e.Field() // Field name
		tag := e.Tag()     // Validation rule
		param := e.Param() // Parameter value
		msg := generateErrorMessage(field, tag, param)
		errMsgs = append(errMsgs, msg)
	}
	return errors.New(strings.Join(errMsgs, ""))
}

func (m *Validator) Engine() interface{} {
	m.lazyinit()
	return m.validate
}

func (m *Validator) lazyinit() {
	if m.validate == nil {
		m.once.Do(func() {
			m.validate = validator.New()
			m.validate.SetTagName("binding")
			m.validateTag = "binding"
		})
	}
}

func (m *Validator) ValidateTag() string {
	return m.validateTag
}

func (m *Validator) SetValidateTag(tag string) {
	vdEngine := m.Engine().(*validator.Validate)
	vdEngine.SetTagName(tag)
	m.validateTag = tag
}
