package other

import (
	"encoding/json"
	"fmt"
	"runtime"
	"testing"
)

type executorParams struct {
	GameId string
}

var m runtime.MemStats

func TestMarshal(t *testing.T) {
	jsonStr := "{\"gameId\":\"aaaa\"}"
	params := executorParams{}
	err := json.Unmarshal([]byte(jsonStr), &params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(params)

	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
