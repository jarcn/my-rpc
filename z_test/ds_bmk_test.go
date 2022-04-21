package ztest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func BenchmarkSendData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sendData()
	}
}

func sendData() {
	url := "http://dev-event.kupu.id/e"
	method := "POST"

	payload := strings.NewReader(`{"app_id":"KUPU","device_id":"7C7F3492-BC23-4160-9A12-9C2DD9932D6F","sdk_version":"1.10","sdk_name":"IOS_SDK","event_report_time":1649330328123,"event_time":1649330327868,"event_id":"q_parttimeNearby_page_enter","u_id":"1511979350428880943","content_info":"{\"app_version\":\"1.3.4\",\"role\":2,\"u_id\":\"1511979350428880943\",\"device\":\"iOS\"}"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
