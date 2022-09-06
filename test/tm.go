package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

type Message struct {
	Elapsed Duration `json:"elapsed"`
}

func main() {
	msgEnc, err := json.Marshal(&Message{
		Elapsed: Duration{time.Second * 5},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", msgEnc)

	var msg Message
	if err := json.Unmarshal([]byte(`{"elapsed": "1h"}`), &msg); err != nil {
		panic(err)
	}
	fmt.Println(msg.Elapsed)
	fmt.Printf("%#v\n", msg)
}
