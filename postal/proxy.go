package postal

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ProxyRecord struct {
	Address      string
	Port         int
	ShortCountry string
	LongCountry  string
	ProxyType    string
	IsGoogle     bool
	IsHttps      bool
}

type ProxyList struct {
	proxies []interface{}
}

func NewProxyRecord(address string, port int, shortCountry string, longCountry string,
	proxyType string, isGoogle bool, isHttps bool) *ProxyRecord {
	return &ProxyRecord{Address: address, Port: port, ShortCountry: shortCountry, LongCountry: longCountry,
		ProxyType: proxyType, IsGoogle: isGoogle, IsHttps: isHttps}
}

func (pr *ProxyRecord) String() string {
	return fmt.Sprintf("%s:%d %s\n", pr.Address, pr.Port, pr.ShortCountry)
}

func (pl *ProxyList) String() string {
	result := ""
	if pl != nil {
		for _, item := range pl.proxies {
			result = result + fmt.Sprintf("%s", item)
		}
	} else {
		return result
	}
	return result
}

func (pl *ProxyList) AddProxy(proxy interface{}) {
	pl.proxies = append(pl.proxies, proxy)
}

func (pl *ProxyList) RandomProxy() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	proxyIdx := r.Intn(len(pl.proxies))
	return pl.proxies[proxyIdx]
}

func NewProxyList() *ProxyList {
	var isgoogle bool
	var ishttps bool
	plist := &ProxyList{}
	body, err := MakeDirectRequest("https://free-proxy-list.net/")
	if err == nil {
		re, _ := regexp.Compile(`\d*\.\d*\.\d*\.\d*\</td><td>\d*\</td><td>\D*\</td><td>\D*\</td><td`)
		res := re.FindAllString(body, -1)
		for _, item := range res {
			pass1 := strings.Replace(item, "</td><td>", ":", -1)
			pass2 := strings.Replace(pass1, "</td><td class=\"hm\">", ":", -1)
			pass3 := strings.Replace(pass2, "</td><td class=\"hx\">", ":", -1)
			pass4 := strings.Replace(pass3, "</td><td", "", -1)
			//fmt.Printf("string: %s\n", pass4)
			parts := strings.Split(pass4, ":")
			port, _ := strconv.Atoi(parts[1])
			isgoogle = false
			if parts[5] == "yes" {
				isgoogle = true
			}
			ishttps = false
			if parts[6] == "yes" {
				ishttps = true
			}
			plist.AddProxy(NewProxyRecord(parts[0], port, parts[2], parts[3], parts[4], isgoogle, ishttps))
		}
	}
	return plist
}
