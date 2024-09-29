package validator

// Validator is a simple struct to hold validation errors.
type Validator struct {
	Errors map[string]string
}

// New creates a new Validator instance with an empty errors map.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid returns true if the errors map doesn't contain any entries.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the map, using the specified key, if one
// doesn't already exist.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message if a validation check 'ok' is false.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Unique returns true if all elements in the input slice are unique.
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
