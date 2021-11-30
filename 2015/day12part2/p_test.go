package p_test

import (
	"aoc/ax"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	input := <-ax.MustReadFileLines("input")
	for i, tc := range []struct {
		jsonInput string
		want      int
	}{
		{"[1,2,3]", 6},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{input, 10},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			require.Equal(t, tc.want, run(tc.jsonInput))
		})
	}
}

func run(line string) int {
	var obj interface{}
	lineBytes := []byte(line)
	ax.Check(json.Unmarshal(lineBytes, &obj), "unmarshal err")
	sum := visit(obj)
	return sum
}

func visit(node interface{}) int {
	var sum int
	switch v := node.(type) {
	case map[string]interface{}:
		for _, val := range v {
			if s, ok := val.(string); ok && s == "red" {
				return 0
			}
			sum += visit(val)
		}
	case []interface{}:
		for _, val := range v {
			sum += visit(val)
		}
	case float64:
		sum += int(v)
	}
	return sum
}
