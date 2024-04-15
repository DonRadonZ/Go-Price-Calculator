package prices

import (
	"fmt"
	
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	Prices            []float64
	TaxRate           float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadPrice() {

	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	  }


	prices, err := conversion.StringsToFloats(lines)



	  if err != nil {
		fmt.Println(err)
		return
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
	
	job.TaxIncludedPrices = result
	
	filemanager.WriteJSON(fmt.Sprintf("result_%0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:  []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}