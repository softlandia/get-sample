package main

// http://supply-plan.purchase.svc.k8s.dataline/api/v1/SupplyPlan/periodReport?dateFrom=02%2F07%2F2020&dateTo=02%2F07%2F2020
// пример запроса к сервису ПЛП (план поставок) (см. скрин "ПЛП.JPG")
// период [02/07/2020 - 02/07/2020]

//TODO выяснить формат даты при запросе к сервису ПЛП, (!) при вызове на slagger вариант с датами в виде 27/06/2020 не прошел

// руками можно попробовать здесь http://supply-plan.purchase.svc.k8s.dataline/swagger/#/ раздел /api/v1/SupplyPlan/periodReport

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("hi\n")
	res, err := http.Get("http://supply-plan.purchase.svc.k8s.dataline/api/v1/limits?dateFrom=01%2F07%2F2020&dateTo=01%2F07%2F2020")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
