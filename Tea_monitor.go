package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type TeaData struct {
	Idmultiplayergameresult int
	Iduser                  string
	Time                    int
	Attack                  int
	Pieces                  int
	Place                   int
	idmultiplayergame       int
	datetime                string
}
type LAEdata struct {
	lpm float32
	apm float32
	eff float32
}

var data []TeaData

func getjson(name string, urllink string) []TeaData {
	var o_data []TeaData
	params := url.Values{}
	Url, err := url.Parse(urllink + "getProfile")
	if err != nil {
		fmt.Println(err)
	}
	params.Set("id", name)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	//fmt.Println(url)
	resp, err := http.Get(urlPath)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &o_data)
	if err != nil {
		fmt.Println(err)
	}
	return o_data
}
func calc_data(i int) LAEdata {
	var laedata LAEdata
	laedata.lpm = 24000 * float32(data[i].Pieces) / float32(data[i].Time)
	laedata.apm = 60000 * float32(data[i].Attack) / float32(data[i].Time)
	laedata.eff = laedata.apm / laedata.lpm
	if laedata.lpm == 0 {
		laedata.eff = 0
	}
	return laedata
}

func get_ft_data(idx int) {
	var laedata LAEdata
	var sum LAEdata
	var win, los int = 0, 0
	for i := 0; win != idx && los != idx; i++ {
		laedata = calc_data(i)
		sum.lpm += laedata.lpm
		sum.apm += laedata.apm
		sum.eff += laedata.eff
		if data[i].Place == 1 {
			fmt.Printf(" %0.3f\t%0.3f\t%0.3f\n", laedata.lpm, laedata.apm, laedata.eff)
			win += 1
		} else {
			fmt.Printf(" %0.3f\t%0.3f\t%0.3f\n", laedata.lpm, laedata.apm, laedata.eff)
			los += 1
		}
	}
	if win == idx {
		fmt.Printf("Win[%d:%d]\n", win, los)
	} else if los == idx {
		fmt.Printf("Lost[%d:%d]\n", win, los)
	}
	fmt.Printf("[lpm:%0.3f][apm:%0.3f][eff:%0.3f]", sum.lpm/float32(win+los), sum.apm/float32(win+los), sum.eff/float32(win+los))

}
func main() {
	var args_len = len(os.Args)
	var name string = "纯粹之狐"
	var urllink string = "https://teatube.cn:8888/"
	var ft int = 15
	if args_len > 1 {
		name = os.Args[1]
	}
	if args_len > 2 {
		urllink = os.Args[2]
	}
	fmt.Printf(" LPM\tAPM\tEFF\n")
	//now_game = getjson(name)[0].Idmultiplayergameresult
	data = getjson(name, urllink)
	get_ft_data(ft)
}
