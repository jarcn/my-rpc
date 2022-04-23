package ztest

import (
	"net/url"
	"testing"
)

var addrs = []string{}

func init() {
	addrs = append(addrs, "Jalan M.H Thamrin 1 & 2, Gondangdia, Kecamatan Menteng, Daerah Khusus ibukota Jakarta")
}

func TestNB(t *testing.T) {
	for i := range addrs {
		uri := "https://api.nextbillion.io/h/geocode?q=" + url.PathEscape(addrs[i]) + "&key=626737e151e6467abb117363d0b67d04"
		// uri := "https://kupu.nextbillion.io/search/json?input=" + url.PathEscape(addrs[i]) + "&key=626737e151e6467abb117363d0b67d04"
		t.Log(uri)
		// res, err := http.Get(uri)
		// if err != nil {
		// 	t.Fatal(err)
		// 	return
		// }
		// defer res.Body.Close()
		// body, err := ioutil.ReadAll(res.Body)
		// if err != nil {
		// 	t.Fatal(err)
		// 	return
		// }
		// t.Log(string(body))
	}
}
