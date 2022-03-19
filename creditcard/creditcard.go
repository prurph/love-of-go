package creditcard

import "errors"

type creditcard struct {
	number string
}

func New(number string) (creditcard, error) {
	if number == "" {
		return creditcard{}, errors.New("Empty number")
	}
	return creditcard{number: number}, nil
}

func (cc creditcard) Number() string {
	return cc.number
}
