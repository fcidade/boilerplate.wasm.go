package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	keepRunning()
}

func keepRunning() {
	<-make(chan bool)
}

func jsonWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed!"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJSON(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to JSON %s\n", err.Error())
			return err.Error()
		}
		return pretty
	})
}

func prettyJSON(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "\t")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
