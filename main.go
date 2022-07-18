package main

import (
	"fmt"
	"regexp"

	"github.com/dongri/phonenumber"
)

// GetProvider returns the provider of the given phone number.
func GetProvider(mobile string) (*Mobile, error) {
	if mobile == "" {
		return nil, fmt.Errorf("mobile number is empty")
	}

	parsedMobile := phonenumber.Parse(mobile, "KE")
	if parsedMobile == "" {
		return nil, fmt.Errorf("mobile number is invalid")
	}

	for _, m := range getAllNumbers() {
		if m.Pattern != "" && regexp.MustCompile(m.Pattern).MatchString(parsedMobile) {
			m.Number = &parsedMobile
			return &m, nil
		}
	}

	return nil, fmt.Errorf("mobile number is not supported")
}

func main() {
	mobile, err := GetProvider("254702547163")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("{\"code\":", mobile.Code, ",\"pattern\":", mobile.Pattern, ",\"number\":", *mobile.Number, ",\"provider\":", mobile.Provider, "}")
}
