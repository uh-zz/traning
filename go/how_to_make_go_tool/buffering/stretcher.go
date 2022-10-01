// +build ignore

package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func parseEvents() (string, error) {

	// 標準入力のバッファリング
	reader := bufio.NewReader(os.Stdin)

	// 読み進めずに先頭の１バイトを覗く
	b, _ := reader.Peek(1)
	if string(b) == "[" {
		// JSONとしてデコード
		var evs ConsulEvents
		dec := json.NewDecoder(reader)
		if err := dec.Decode(&evs); err != nil {
			return "", err
		}

		ev := &evs[len(evs)-1]
		return string(ev.Payload), nil
	} else {
		// JSONでなければ改行終端なので1行読み取る
		line, err := reader.ReadString('\n')
		return line, err
	}
}
