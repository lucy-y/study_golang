package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"os"
)

type extractedJob struct {
	id			string
	title		string
	location 	string
	salary		string
	summary		string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	
	for i := 0; i<totalPages; i++ {
		extractedJob := getPage(i)
		
		jobs = append(jobs, extractedJob...) // [ x x x ]
		// jobs = append(jobs, extractedJob)  // [ [x] [x] [x] ]
	}
	writeJobs(jobs)
	fmt.Println("End... totalJobCount: ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link","Title","Location","Salary","Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/jobs?jk="+job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

}

func getPage(page int) []extractedJob{
	var jobs []extractedJob

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50);
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc,err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection){
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs

}

func extractJob(card *goquery.Selection) extractedJob{
	
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())

	return extractedJob{
		id: id, 
		title: title, 
		location: location, 
		salary: salary, 
		summary: summary,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("failed : ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

