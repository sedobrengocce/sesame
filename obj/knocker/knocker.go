package knocker

import (
	"math/rand"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/sedobrengocce/sesame/obj/port"
)

type Knocker struct {
    host string `yaml:"host"`
    sequence []port.Port `yaml:"sequence"`
    delay int `yaml:"delay"`
    udp bool `yaml:"udp"`
    verbose bool
    fake bool
}

func NewKnocker(host string) *Knocker {
    return &Knocker{
        host: host,
        sequence: []port.Port{},  
        delay: 10,
        udp: false,
        verbose: false,
        fake: false,
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

func (k *Knocker) WithVerbose(verbose bool) {
    k.verbose = verbose
}

func (k *Knocker) WithFake(fake bool) {
    k.fake = fake
}

func (k *Knocker) Knock() error {
    packetType := "tcp"
    if k.udp {
        packetType = "udp"
    }
    fmt.Printf("Sesame ...\n")
    host := k.host
    sequence := []int{}
    for _, p := range k.sequence {
        sequence = append(sequence, p.Number())
    }
    verbPacketType := packetType
    if k.fake {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        host = fmt.Sprintf("%d.%d.%d.%d", r.Intn(254), r.Intn(254), r.Intn(254), r.Intn(254))
        l := len(k.sequence)
        sequence = []int{}
        for i := 0; i < l; i++ {
            sequence = append(sequence, r.Intn(65535))
        }
        if rand.Intn(1) == 0 {
            verbPacketType = "tcp"
        } else {
            verbPacketType = "udp"
        }
    }
    for i, p := range k.sequence {
        if k.verbose || k.fake {
            fmt.Printf("Knocking on %s:%d with %s\n", host, sequence[i], verbPacketType)
        }
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
    fmt.Println("... open your self!")

    return nil
}

