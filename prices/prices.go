package prices

import (
	"fmt"
	
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager 		  filemanager.FileManager `json:"-"`
	Prices            []float64 `json:"prices"`
	TaxRate           float64 `json:"tax_rate"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrice() {

	lines, err := job.IOManager.ReadLines("")

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
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		Prices:  []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}