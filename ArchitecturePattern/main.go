package main

import "arch/singleton"
import "arch/factory"

func main() {
	singleton.SingletonPattern()

	payment := factory.GetPaymentMethod("credit")
	if payment != nil {
		println(payment.Pay(100))
	} else {
		println("Invalid payment method")
	}
}
