package helper

import "fmt"

// GetMapValueSafe is a function that safely retrieves a value from a map based on a provided key and type.
// It returns the value if it exists and matches the provided type, otherwise it returns an error.
//
// Parameters:
// m: The map from which to retrieve the value.
// param: The key of the value to retrieve.
// t: The expected type of the value ("string" or "int").
func GetMapValueSafe(m map[string]interface{}, param string, t string) (interface{}, error) {
	//
	// Check if the key exists in the map
	//
	if validated, ok := m[param]; ok {
		//
		// If the key exists, check the type of the value
		//
		switch t {
		case "string":
			//
			// If the expected type is string, try to cast the value to string
			//
			if str, ok := validated.(string); ok {
				// If the cast is successful, return the string value
				return str, nil
			}
		case "int":
			//
			// If the expected type is int, try to cast the value to int
			//
			if i, ok := validated.(int); ok {
				// If the cast is successful, return the int value
				return i, nil
			}
		default:
			//
			// If the expected type is not supported, return an error
			//
			return nil, fmt.Errorf("Unsupported type %s", t)
		}
	}
	//
	// If the key does not exist in the map, return an error
	//
	return nil, fmt.Errorf("Key %s not found", param)
}
