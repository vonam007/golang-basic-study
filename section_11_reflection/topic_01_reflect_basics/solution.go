package reflectbasics

import "reflect"

// TypeInfo trả về tên kiểu và Kind của giá trị.
func TypeInfo(v interface{}) (typeName string, kind reflect.Kind) {
	// TODO: Implement this
	return "", reflect.Invalid
}

// FieldNames trả về slice tên các exported fields của struct.
// Nếu v không phải struct, trả về nil.
func FieldNames(v interface{}) []string {
	// TODO: Implement this
	return nil
}

// GetField lấy giá trị field theo tên từ struct.
// Trả về (value, true) hoặc (nil, false) nếu field không tồn tại.
func GetField(v interface{}, fieldName string) (interface{}, bool) {
	// TODO: Implement this
	return nil, false
}

// IsZero kiểm tra giá trị có phải zero value không.
func IsZero(v interface{}) bool {
	// TODO: Implement this
	return false
}
