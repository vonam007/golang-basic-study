package structtags

// GetJSONTags trả về map[fieldName]jsonTag cho struct.
// Nếu field không có json tag, bỏ qua.
func GetJSONTags(v interface{}) map[string]string {
	// TODO: Implement this using reflect
	return nil
}

// ToMap chuyển struct thành map[string]interface{} dùng json tag names.
// Fields không có json tag dùng field name.
// Fields có tag "-" bị skip.
func ToMap(v interface{}) map[string]interface{} {
	// TODO: Implement this
	return nil
}

// HasTag kiểm tra struct field có tag key cụ thể không.
func HasTag(v interface{}, fieldName, tagKey string) bool {
	// TODO: Implement this
	return false
}
