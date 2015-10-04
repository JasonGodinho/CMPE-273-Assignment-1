package main

import (
	"fmt"
	"net/rpc"
	"strconv"
	"strings"
)

var Sr1 StockRequestObject
var Sresp StockResponseObject

type StockRequestObject struct {
	Name       [5]string
	Percentage [5]int
	Budget     float32
	TradeId    int
}

type StockResponseObject struct {
	TradeId            int
	Name               [5]string
	NumberOfStocks     [5]int
	StockValue         [5]float64
	UnvestedAmount     float64
	CurrentMarketValue float64
	ProfitLoss         [5]string
}

func client() {
	var i3 string
	fmt.Scanln(&i3)
}
func main() {
	GetInput()
	var input string
	fmt.Scanln(&input)
}

func GetInput() {
	var Sr1 StockRequestObject
	var Sresp StockResponseObject
	var InputFromUser int
	InputFromUser = 0

	fmt.Println("Enter 1 for buying stocks or 2 for getting portfolio")
	fmt.Scanln(&InputFromUser)

	if InputFromUser == 2 {
		fmt.Println("Enter the trade ID")
		Trid := ""
		fmt.Scanln(&Trid)
		NewInt, err := strconv.Atoi(Trid)
		if err != nil {
			fmt.Println("Error occurred. Please enter proper trade id")
		}
		Sr1.TradeId = NewInt
	} else if InputFromUser == 1 {
		var BudgetData float32
		fmt.Println("Enter the budget")
		fmt.Scanln(&BudgetData)

		Sr1.Budget = BudgetData

		fmt.Println("Enter the stock values in foll format : ")
		fmt.Println("StockName1,Percentage1,StockName2,Percentage2,StockName3,Percentage3")
		fmt.Println("Example:")
		fmt.Println("Goog,50,YHOO,50")

		UserString := ""
		fmt.Scanln(&UserString)

		var s []string
		s = strings.Split(UserString, ",")
		Latestindex := 0
		for index := 0; Latestindex < len(s); index++ {
			Sr1.Name[index] = s[Latestindex]
			s3, _ := strconv.Atoi(s[Latestindex+1])
			Sr1.Percentage[index] = int(s3)
			Latestindex = Latestindex + 2
		}
	} else {
		fmt.Println("Nothing found")
	}
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	if InputFromUser == 1 {

		err = c.Call("Server.Receive", Sr1, &Sresp)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println()
			fmt.Println()
			fmt.Println("Trade ID: ", Sresp.TradeId, " and Remaining amount: ", Sresp.UnvestedAmount)
			for index := 0; index < len(Sresp.NumberOfStocks); index++ {
				if Sresp.Name[index] != "" {
					fmt.Println("Name of Stock: ", Sresp.Name[index], " Number of Stocks: ", Sresp.NumberOfStocks[index], " Value of Stocks bought: ", Sresp.ProfitLoss[index], Sresp.StockValue[index])
				}
			}
		}
	} else if InputFromUser == 2 {
		err = c.Call("Server.GetPortfolio", Sr1, &Sresp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("Trade ID: ", Sresp.TradeId, " and Remaining amount: ", Sresp.UnvestedAmount)

		for index := 0; index < len(Sresp.NumberOfStocks); index++ {
			if Sresp.Name[index] != "" {
				fmt.Println("Name of Stock:", Sresp.Name[index], " Number of Stocks:", Sresp.NumberOfStocks[index], " Value of Stock:", Sresp.ProfitLoss[index], Sresp.StockValue[index])
			}
		}
		fmt.Println("Current total value ", Sresp.CurrentMarketValue)
	}
}
