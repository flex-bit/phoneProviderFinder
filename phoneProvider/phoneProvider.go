package phoneProvider

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

type phoneUsNumber struct {
	AreaCode string `json:"arecode"`
	Prefix   string `json:"prefix"`
	Local    string `json:"local"`
}

type PhoneInfo struct {
	Provider string `json:"provider"`
	State    string `json:"state"`
	City     string `json:"city"`
	Type     string `json:"type"`
}

func PhoneProviderGet(number string) PhoneInfo {
	//replace any character that is not a number
	patternDeleteCharacterNonDigit := regexp.MustCompile("[^\\d]")
	cleanNumber := patternDeleteCharacterNonDigit.ReplaceAllString(number, "")

	//check if phone number is 10 digits long if not terminate
	if len(cleanNumber) != 10 {
		log.Fatalln("\nYour phone number should be 10 digit long please try again")
	}
	// Extract phone number into 3 peaces, AreaCode,Prefix Local
	patternExtractThreePeacesPhoneNumber := regexp.MustCompile("(\\d{0,3})(\\d{0,3})(\\d{0,4})")
	providePhoneNumberThreePeaces := patternExtractThreePeacesPhoneNumber.FindStringSubmatch(cleanNumber)
	currentPhoneNumber := phoneUsNumber{AreaCode: providePhoneNumberThreePeaces[1], Prefix: providePhoneNumberThreePeaces[2], Local: providePhoneNumberThreePeaces[3]}
	apiUrl := "https://www.fonefinder.net/findome.php?npa=" + currentPhoneNumber.AreaCode + "&nxx=" + currentPhoneNumber.Prefix + "&thoublock=" + currentPhoneNumber.Local
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
	return phoneinfo
}
