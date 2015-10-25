package single_parse_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkParsingA5MBJSONDocument(b *testing.B) {
	raw, err := ioutil.ReadFile("AllCards.json")
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(raw)
		dec := json.NewDecoder(reader)
		var decoded map[string]interface{}
		err = dec.Decode(&decoded)
		if err != nil {
			log.Fatalln("error decoding:", err.Error())
		}
	}
}
