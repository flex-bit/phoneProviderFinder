package phoneProvider

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

type PhoneInfo struct {
	Provider string `json:"provider"`
	State    string `json:"state"`
	City     string `json:"city"`
	Type     string `json:"type"`
}

func PhoneProviderGet() {
	apiUrl := "https://www.fonefinder.net/findome.php?npa=614&nxx=364&thoublock=1215"
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}
	// Getting City, State, Typ
	pattern := regexp.MustCompile("findcity\\.php\\?.*?cityname=(\\w+)&state=(\\w+)'>\\w+.*?(\\w+)<\\/A.*?(\\w+).php'>(.*?)<\\/A><TD>(.*?)<")
	matchPhoneInfo := pattern.FindStringSubmatch(string(body))
	phoneinfo := PhoneInfo{Provider: matchPhoneInfo[4], State: matchPhoneInfo[3], City: matchPhoneInfo[1], Type: matchPhoneInfo[6]}
	fmt.Printf("State: %s, City: %s, Provider: %s", phoneinfo.State, phoneinfo.City, phoneinfo.Provider)
}
