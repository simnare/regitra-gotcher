package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
	"regexp"
)

func main() {
	const City = "VL"
	res, err := http.Get(fmt.Sprintf("https://www.eregitra.lt/viesa/interv/INT_GRAFIKAS_VIEW_T.php?Padalinys=%v&Action=", City))

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	z := html.NewTokenizer(res.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.TextToken:
			t := z.Token()

			r, _ := regexp.Compile("[0-9]{4}.[0-9]{2}.[0-9]{2}")

			if (!r.MatchString(t.Data)) {
				continue
			}

			date := r.FindString(t.Data)
			fmt.Println("Got a date", date)
		}
	}

}
