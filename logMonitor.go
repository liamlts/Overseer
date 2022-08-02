package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type IPInfo struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Geo struct {
			Host          string  `json:"host"`
			IP            string  `json:"ip"`
			Rdns          string  `json:"rdns"`
			Asn           int     `json:"asn"`
			Isp           string  `json:"isp"`
			CountryName   string  `json:"country_name"`
			CountryCode   string  `json:"country_code"`
			RegionName    string  `json:"region_name"`
			RegionCode    string  `json:"region_code"`
			City          string  `json:"city"`
			PostalCode    string  `json:"postal_code"`
			ContinentName string  `json:"continent_name"`
			ContinentCode string  `json:"continent_code"`
			Latitude      float64 `json:"latitude"`
			Longitude     float64 `json:"longitude"`
			MetroCode     int     `json:"metro_code"`
			Timezone      string  `json:"timezone"`
			Datetime      string  `json:"datetime"`
		} `json:"geo"`
	} `json:"data"`
}

func (info IPInfo) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\n", info.Data.Geo.Host, info.Data.Geo.Rdns, info.Data.Geo.CountryName, info.Data.Geo.City)
}

type String string

func (s String) FilterValue() string { return string(s) }

func apiReq(ip string) IPInfo {
	api := "https://tools.keycdn.com/geo.json?host="
	full := api + ip

	client := &http.Client{}

	req, err := http.NewRequest("GET", full, nil)
	check(err)

	req.Header.Set("User-Agent", "keycdn-tools:https://www.example.com")
	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	var ipinfo IPInfo
	json.Unmarshal(body, &ipinfo)

	return ipinfo
}

func MonitLogs() []String {
	var file []string
	dat, err := os.Open("/var/log/auth.log")
	check(err)
	defer dat.Close()
	lines := bufio.NewScanner(dat)

	for lines.Scan() {
		file = append(file, lines.Text())
	}
	var ipList []string
	for i := range file {
		if strings.Contains(file[i], "Connection closed by") {
			ts := strings.Fields(file[i])
			for ind2 := range ts {
				if strings.HasPrefix(ts[ind2], "by") {
					if ts[ind2+1] != "remote" && ts[ind2+1] != "invalid" && ts[ind2+1] != "/var/log/auth.log" && ts[ind2+1] != "authenticating" {
						ipList = append(ipList, ts[ind2+1])
					}
				}
			}
		}
	}
	if err := lines.Err(); err != nil {
		log.Fatal(err)
	}
	var mIpList []String
	for r := range ipList {
		mIpList = append(mIpList, String(ipList[r]))
	}
	return mIpList
}

func GetMalIps(ips []string) []string {
	ipsHM := make(map[int]string)
	var malIps []string

	for i, s := range ips {
		ipsHM[i] = s
	}

	for key := range ipsHM {
		if ipsHM[key] == ipsHM[key+1] {
			malIps = append(malIps, ipsHM[key])
		}
	}
	return malIps
}

func geoData(malIps []string) map[string]IPInfo {
	geoData := make(map[string]IPInfo)

	for i, ip := range malIps {
		//Rate limited API to 3r/s
		time.Sleep(3001 * time.Millisecond)
		fmt.Println("Getting info on: ", malIps[i])
		geoData[ip] = apiReq(malIps[i])
	}
	return geoData
}

func MalDns() []string {
	ipList := MonitLogs()
	var list []string
	for i := range ipList {
		list = append(list, string(ipList[i]))
	}

	ipHM := geoData(list)
	var dnslist []string
	for _, ipinfo := range ipHM {
		if strings.Contains(ipinfo.Data.Geo.Rdns, "tor") || strings.Contains(ipinfo.Data.Geo.Rdns, "scanner") || strings.Contains(ipinfo.Data.Geo.Rdns, "census") || strings.Contains(ipinfo.Data.Geo.Rdns, "EMERALD-ONION") {
			fmt.Println("Found malicious DNS. Adding to bad actor list", ipinfo.Data.Geo.Rdns)
			dnslist = append(dnslist, ipinfo.Data.Geo.IP)

		}
	}
	return dnslist
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
