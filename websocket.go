package ftx

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Dialer struct {
	conn net.Conn
	subs map[interface{}][]chan interface{}
}

func Dial() *Dialer {
	conn, _, _, err := ws.Dial(context.TODO(), "wss://ftx.com/ws/")
	if err != nil {
		panic(err)
	}

	d := &Dialer{
		conn: conn,
	}

	go func() {
		for range time.Tick(10 * time.Second) {
			d.Do("ping", nil)
		}
	}()

	return d
}

func (d *Dialer) Do(op string, args interface{}) error {
	b, err := json.Marshal(map[string]interface{}{
		"args": args,
		"op":   op,
	})
	if err != nil {
		return err
	}

	return wsutil.WriteClientText(d.conn, b)
}

func (d *Dialer) Login(key, secret, acc string) error {
	t := time.Now().UnixNano() / 1000000

	return d.Do("login", map[string]interface{}{
		"key":        key,
		"time":       t,
		"subaccount": nil,
		"sign":       sign([]byte(secret), fmt.Sprintf("%dwebsocket_login", t)),
	})
}

func (d *Dialer) Subscribe(ch string) error {
	b, err := json.Marshal(map[string]interface{}{
		"channel": ch,
		"op":      "subscribe",
	})
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return wsutil.WriteClientText(d.conn, b)
}

func (d *Dialer) Ticker(market string) error {
	b, err := json.Marshal(map[string]interface{}{
		"channel": "ticker",
		"market":  market,
		"op":      "subscribe",
	})
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return wsutil.WriteClientText(d.conn, b)
}

func (d *Dialer) Conn() net.Conn {
	return d.conn
}
