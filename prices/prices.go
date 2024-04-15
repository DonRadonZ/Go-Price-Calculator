package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	Prices            []float64
	TaxRate           float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) LoadPrice() {
	file,err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
	  floatPrice, err := strconv.ParseFloat(line, 64)

	  if err != nil {
		fmt.Println("Converting price to float failed.")
		fmt.Println(err)
		file.Close()
		return
	  }

	  prices[lineIndex] = floatPrice
	}

  job.Prices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrice()

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:  []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}