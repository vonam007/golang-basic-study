package constants

// Day type đại diện cho ngày trong tuần
type Day int

// Dùng iota để define các hằng số cho ngày trong tuần.
// Sunday = 0, Monday = 1, ..., Saturday = 6
const (
	// TODO: Define days using iota
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// DayName trả về tên tiếng Anh của ngày.
// Ví dụ: DayName(Sunday) → "Sunday", DayName(Monday) → "Monday"
// Nếu day không hợp lệ (< 0 hoặc > 6), trả về "Unknown".
func DayName(d Day) string {
	// TODO: Implement this
	return ""
}

// ByteSize type dùng iota với bit shifting
type ByteSize uint64

const (
	_  ByteSize = iota             // ignore 0
	KB ByteSize = 1 << (10 * iota) // 1024
	MB                             // 1048576
	GB                             // 1073741824
	TB                             // 1099511627776
)

// ToBytes chuyển đổi giá trị size theo đơn vị unit sang bytes.
// unit: "KB", "MB", "GB", "TB"
// Ví dụ: ToBytes(5, "KB") → 5120
// Nếu unit không hợp lệ, trả về 0.
func ToBytes(size uint64, unit string) uint64 {
	// TODO: Implement this
	return 0
}

// IsWeekend kiểm tra ngày có phải cuối tuần không (Saturday hoặc Sunday).
func IsWeekend(d Day) bool {
	// TODO: Implement this
	return false
}
