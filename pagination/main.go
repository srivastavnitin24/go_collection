package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	pagination()

}

func pagination() {
	for i := 1; ; i++ {
		res := make(map[string]interface{})
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://jsonmock.hackerrank.com/api/articles", nil)
		req.Header.Add("Accept", "application/json")
		q := req.URL.Query()
		//q.Add("author", author)
		q.Add("page", strconv.Itoa(i))
		req.URL.RawQuery = q.Encode()
		fmt.Println("query : ", req.URL)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Errored when sending request to the server")
		}

		defer resp.Body.Close()

		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		err = json.Unmarshal([]byte(respBody), &res)
		if err != nil {
			fmt.Println("Errored when sending request to the server")
		}
		fmt.Println("--------- : ", res["page"])
		fmt.Println("--------- : notttttttttttt ")
		if res["data"] == nil {
			fmt.Println("--------- : hereeeeee ")
			break
		}

	}

}
