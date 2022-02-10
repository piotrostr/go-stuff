package wallet_test

import (
	. "learn_tests/wallet"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wallet", func() {
	It("deposits", func() {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		Expect(got).To(Equal(want))
	})

	It("retrieves balance", func() {
		wallet := Wallet{}
		got := wallet.Balance()
		want := Bitcoin(0)
		Expect(got).To(Equal(want))
	})

	It("withdraws when funds are there", func() {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(0)
		Expect(got).To(Equal(want))
	})

	It("throws err if insufficient balance", func() {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(10))
		Expect(err).NotTo(BeNil())
		Expect(err).To(MatchError("insufficient balance"))
	})
})
