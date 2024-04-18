package prices

import (
	
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager 		  iomanager.IOManager `json:"-"`
	Prices            []float64 `json:"prices"`
	TaxRate           float64 `json:"tax_rate"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrice() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	  }


	prices, err := conversion.StringsToFloats(lines)



	  if err != nil {
		return err
	  }

  job.Prices = prices
  return nil
  
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error)  {
	err := job.LoadPrice()


	// errorChan <- errors.New("An error!")

	if err != nil {
	  // return err
	  errorChan <- err
	  return
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		Prices:  []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}