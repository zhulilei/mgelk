package libs

import (
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"log"
)

var uri map[string][]string

func Geturi(index, domain, fromtime, totime string, topX int) map[string][]string {
	var uriList []string
	uri = make(map[string][]string)
	client, err := elastic.NewClient(
		elastic.SetURL("http://10.170.164.76:9200"))
	if err != nil {
		panic(err)
	}

	from := DealTime(fromtime)
	to := DealTime(totime)

	//dateRangeAgg := elastic.NewRangeQuery("@timestamp").From("2016/09/28-06:39").To("2016/09/28-06:40").Format("yyyy/MM/dd-HH:mm")
	dateRangeAgg := elastic.NewRangeQuery("@timestamp").From(from).To(to).Format("yyyy/MM/dd-HH:mm")
	aggline := elastic.NewTermsAggregation().Field("uri.raw").Size(topX).OrderByCountDesc()
	termQuery := elastic.NewTermQuery("host.raw", domain)

	bq := elastic.NewBoolQuery()
	bq = bq.Must(dateRangeAgg)
	bq = bq.Must(termQuery)

	searchResult, err := client.Search().
		Index(index).
		Query(bq).                       // search in index "twitter"               // return all results, but ...
		SearchType("count").             // ... do not return hits, just the count
		Aggregation("aggline", aggline). // add our aggregation to the query
		Pretty(true).                    // pretty print request and response JSON
		Do()                             // execute
	if err != nil {
		panic(err)
	}

	agg, found := searchResult.Aggregations.Terms("aggline")
	fmt.Printf("Found a total of %d datas\n", searchResult.TotalHits())
	if !found {
		log.Fatalf("we sould have a terms aggregation called %q", "aggline")
	}
	for _, userBucket := range agg.Buckets {
		//fmt.Println(userBucket.Key)
		uriList = append(uriList, fmt.Sprintf("%s", userBucket.Key))
	}
	uri[domain] = uriList
	//fmt.Println(uri)
	return uri

}
