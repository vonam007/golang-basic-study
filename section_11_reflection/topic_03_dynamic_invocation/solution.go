package dynamicinvocation

import "reflect"

// CallMethod gọi method theo tên trên object.
// Trả về kết quả (slice reflect.Value) và error nếu method không tồn tại.
func CallMethod(obj interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	// TODO: Implement this
	_ = reflect.ValueOf(obj)
	return nil, nil
}

// SetField set giá trị field trong struct pointer bằng reflection.
// Trả về error nếu field không tồn tại hoặc type mismatch.
func SetField(structPtr interface{}, fieldName string, value interface{}) error {
	// TODO: Implement this
	return nil
}

// MethodNames trả về danh sách tên methods của object.
func MethodNames(obj interface{}) []string {
	// TODO: Implement this
	return nil
}
