package methods

import "testing"

func TestDeposit(t *testing.T) {
	a := &BankAccount{Owner: "Alice", Balance: 100}
	if !a.Deposit(50) || a.Balance != 150 {
		t.Errorf("After Deposit(50), Balance = %v; want 150", a.Balance)
	}
	if a.Deposit(-10) {
		t.Error("Deposit(-10) should return false")
	}
	if a.Deposit(0) {
		t.Error("Deposit(0) should return false")
	}
}

func TestWithdraw(t *testing.T) {
	a := &BankAccount{Owner: "Bob", Balance: 100}
	if !a.Withdraw(30) || a.Balance != 70 {
		t.Errorf("After Withdraw(30), Balance = %v; want 70", a.Balance)
	}
	if a.Withdraw(200) {
		t.Error("Withdraw(200) should fail - insufficient funds")
	}
	if a.Withdraw(-5) {
		t.Error("Withdraw(-5) should return false")
	}
}

func TestTransfer(t *testing.T) {
	a := &BankAccount{Owner: "Alice", Balance: 100}
	b := &BankAccount{Owner: "Bob", Balance: 50}

	if !a.Transfer(b, 30) {
		t.Fatal("Transfer should succeed")
	}
	if a.Balance != 70 || b.Balance != 80 {
		t.Errorf("After transfer: A=%v B=%v; want A=70 B=80", a.Balance, b.Balance)
	}
	if a.Transfer(b, 200) {
		t.Error("Transfer 200 should fail")
	}
}

func TestSummary(t *testing.T) {
	a := BankAccount{Owner: "Alice", Balance: 100.50}
	want := "Alice: $100.50"
	if got := a.Summary(); got != want {
		t.Errorf("Summary() = %q; want %q", got, want)
	}
}
