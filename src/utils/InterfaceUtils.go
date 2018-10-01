package utils
func InterfaceArray2StringArray(interfaceArray []interface{})[]string{
	var stringArray []string
	for _, param := range interfaceArray {
		stringArray = append(stringArray, param.(string))
	}
	return stringArray
}
