package port

import "errors"

type Port struct {
    number int
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
