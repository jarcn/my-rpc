package ztest

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

var addrs = []string{}

func init() {
	// addrs = append(addrs, "Jalan M.H Thamrin 1 & 2, Gondangdia, Kecamatan Menteng, Daerah Khusus ibukota Jakarta")
	// addrs = append(addrs, "Sinarmas Land Plaza, Tower II, Jl. M.H. Thamrin No.51, RT.9/RW.4, Gondangdia, Kec. Menteng, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10350, Indonesia")
	// addrs = append(addrs, "Kuningan, Karet Kuningan, Kecamatan Setiabudi, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta, Indonesia")
	// addrs = append(addrs, "Cibodas, Tangerang City, Banten, Indonesia")
}

func TestNB(t *testing.T) {
	for i := range addrs {
		uri := "https://api.nextbillion.io/h/geocode?q=" + url.PathEscape(addrs[i]) + "&key=626737e151e6467abb117363d0b67d04"
		// uri := "https://kupu.nextbillion.io/search/json?input=" + url.PathEscape(addrs[i]) + "&key=626737e151e6467abb117363d0b67d04"
		t.Log(uri)
		res, err := http.Get(uri)
		if err != nil {
			t.Fatal(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log(string(body))
	}
}
