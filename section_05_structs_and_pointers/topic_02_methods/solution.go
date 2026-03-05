package methods

// BankAccount đại diện tài khoản ngân hàng.
type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit thêm tiền vào tài khoản. Amount phải > 0.
// Trả về true nếu thành công.
func (a *BankAccount) Deposit(amount float64) bool {
	// TODO: Implement this (pointer receiver vì mutate balance)
	return false
}

// Withdraw rút tiền. Amount phải > 0 và <= Balance.
// Trả về true nếu thành công.
func (a *BankAccount) Withdraw(amount float64) bool {
	// TODO: Implement this
	return false
}

// Transfer chuyển tiền từ tài khoản này sang tài khoản khác.
// Trả về true nếu thành công.
func (a *BankAccount) Transfer(to *BankAccount, amount float64) bool {
	// TODO: Implement this
	return false
}

// Summary trả về string "Owner: $Balance" (value receiver vì chỉ đọc).
func (a BankAccount) Summary() string {
	// TODO: Implement this
	return ""
}
