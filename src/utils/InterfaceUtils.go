package utils

//接口数组转字符串数组
func InterfaceArray2StringArray(interfaceArray []interface{}) []string {
	var stringArray []string
	for _, param := range interfaceArray {
		stringArray = append(stringArray, param.(string))
	}
	return stringArray
}
