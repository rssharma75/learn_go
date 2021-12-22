package pointers_test

import (
	"testing"

	"github.com/rssharma75/learn_go/pointers"
)

func assertBalance(got, want pointers.Bitcoin, t *testing.T) {
	t.Helper()
	if want != got {
		t.Errorf("test failed, got:%s, expected:%s", got, want)
	}

}

func assertError(t *testing.T, err error, want string) {
	if err == nil {
		t.Errorf("Test failed, expected error but found none")
	}

	if err.Error() != want {
		t.Errorf("Test Failed, expected error:%s. got error:%s", want, err)
	}

}
func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))

		got := wallet.Balance()

		want := pointers.Bitcoin(10)

		assertBalance(got, want, t)

	})
}

func TestWithdrawl(t *testing.T) {
	t.Run("Withdrawl", func(t *testing.T) {
		wallet := pointers.Wallet{}

		wallet.Deposit(pointers.Bitcoin(20))

		wallet.Withdraw(pointers.Bitcoin(10))

		got := wallet.Balance()

		want := pointers.Bitcoin(10)

		assertBalance(got, want, t)
	})
	t.Run("WIthdrawl Insufficient funds", func(t *testing.T) {
		wallet := pointers.Wallet{}

		wallet.Deposit(pointers.Bitcoin(20))

		err := wallet.Withdraw(pointers.Bitcoin(100))

		// got := wallet.Balance()

		// want := pointers.Bitcoin(10)

		// assertBalance(got, want, t)
		wantErr := "cannot withdraw, insufficient funds"
		assertError(t, err, wantErr)
	})
}
