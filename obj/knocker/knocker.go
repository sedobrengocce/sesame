package knocker

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/sedobrengocce/sesame/obj/port"
)

type Knocker struct {
    host string
    sequence []port.Port
    delay int
    udp bool
}

func NewKnocker(host string) *Knocker {
    return &Knocker{
        host: host,
        sequence: []port.Port{},  
        delay: 10,
        udp: false,
    }
}

func (k *Knocker) WithSequence(sequence []int) error {
    tmp := []port.Port{}
    for _, s := range sequence {
        p, err := port.NewPort(s)
        if err != nil {
            return err
        }
        tmp = append(tmp, *p)
    }
    if len(tmp) < 2 {
        return errors.New("sequence must have at least 2 ports")
    }
    k.sequence = tmp
    return nil
}

func (k *Knocker) WithDelay(delay int) {
    k.delay = delay
}

func (k *Knocker) WithUdp(udp bool) {
    k.udp = udp
}

func (k *Knocker) Knock() error {
    packetType := "tcp"
    if k.udp {
        packetType = "udp"
    }
    for _, p := range k.sequence {
        fmt.Printf("Knocking on %s:%d with %s\n", k.host, p.Number(), packetType)
        target := fmt.Sprintf("%s:%d", k.host, p.Number())
        conn, err := net.DialTimeout(packetType, target, 500 * time.Millisecond)
        if err != nil {
            e := err.(net.Error)
            if !e.Timeout() {
                return err
            }
        } else {
            defer conn.Close()
        }
        time.Sleep(time.Duration(k.delay) * time.Millisecond)
    }
    return nil
}

