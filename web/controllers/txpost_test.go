package controllers

import (
	"testing"
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"

	"unichain-go/core"
)

func TestTXController_Post(t *testing.T) {
	url := "http://127.0.0.1:19984/tx"

	tx := core.CreateTransactionForTest()
	var jsonStr = []byte(tx)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}