package port

import (
	"errors"
	"strconv"
)

type Port struct {
    number int `yaml:"number"`
}

func NewPort(number int) (*Port, error) {
    if number < 0 || number > 65535 {
        return nil, errors.New("invalid port number")
    }
    return &Port{
        number: number,
    }, nil
}

func (p *Port) Number() int {
    return p.number
}

func (p *Port) String() string {
    return strconv.Itoa(p.number)
}
