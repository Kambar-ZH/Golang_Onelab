package acmp

import (
	"bufio"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

const (
	pattern = `Difficulty: (\d+)%\)`
)

func Difficulty(url string) float64 {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"Cookie": []string{"English=1"},
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(resp.Body)
	re := regexp.MustCompile(pattern)
	for sc.Scan() {
		str := sc.Text()
		res := re.FindAllStringSubmatch(str, 1)
		for i := range res {
			diff, err := strconv.Atoi(res[i][1])
			if err != nil {
				return -1
			} else {
				return float64(diff)
			}
		}
	}
	return -1
}