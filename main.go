package main

import(
	`fmt`
)

func main(){

	purchases:=[...]int{19,13,16,45,76}

	percent := 10
	limit := 30
	const fullPercent=100
	sumCashback:=0
	sumOfPurchase :=0

	for _, purchase := range purchases{
		sumOfPurchase+=purchase
	}

	sumMonth:=sumOfPurchase*percent/fullPercent

	sumCashback+=sumMonth

	if sumCashback>limit{
		fmt.Println(`сумма кэшбека =`, limit)
	} else {
		fmt.Println(`сумма кэшбека =`, sumCashback)
	}
}
