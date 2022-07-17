package metrics

import (
	"fmt"
	"log"
	"sync"

	models "github.com/events/models"
	parse "github.com/events/parse"
)

// Worker function to consume data and caliculate metrics
// TODO possible int overflow
// Lets use sync.map https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
func ConsumeData(wg *sync.WaitGroup, msg chan string, sm *sync.Map) {
	wg.Add(1)
	defer wg.Done()
	for {
		data := <-msg

		if data == "END" {
			break
		}
		element, err := parse.ParseData(data)
		if err != nil {
			log.Println("Cannot parse input string ")
			continue
		}
		//Update stats
		output, status := sm.Load(element.Market)
		mk := models.Market{}
		//Calculate values and store
		if !status {
			mk.VolWeighted = element.Price
		} else {
			mk = output.(models.Market)
			mk.VolWeighted = ((mk.VolWeighted * mk.TotalVol) + (element.Volume * element.Price)) / (mk.TotalVol + element.Volume)
		}
		mk.NumOfElements += 1
		if element.IsBuy {
			mk.NumOfTrue += 1
		}
		mk.SumOfVolume += element.Volume
		mk.SumOfPrice += element.Price
		mk.TotalVol += element.Volume

		mk.SetPercent()
		mk.SetMeanPrice()
		mk.SetMeanVol()
		sm.Store(element.Market, mk)
	}
	log.Println("Done consuming ")
}

//Print metrics to stdout
func PrintMetrics(sm *sync.Map) {
	sm.Range(func(k, v interface{}) bool {
		o := v.(models.Market)
		log.Printf(fmt.Sprintf("{\"market\":%d,\"total_volume\":%f,\"mean_price\":%f,\"mean_volume\":%f,\"volume_weighted_average_price\":%f,\"percentage_buy\":%f}",
			k, o.TotalVol, o.MeanPrice, o.MeanVol, o.VolWeighted, o.PercentBuy))
		return true
	})
}
