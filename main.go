
package main
import (
    "encoding/json"
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
    "os"
)
type TeaData struct {
    Idmultiplayergameresult  int
    Iduser   string
    Time    int
    Attack int
    Pieces int
    Place int
    idmultiplayergame int
    datetime string
}
type LAEdata struct {
    lpm float32
    apm float32
    eff float32
}
func calc_data(data []TeaData,i int) LAEdata{
    var laedata LAEdata
    laedata.lpm = 24000*float32(data[i].Pieces)/float32(data[i].Time)
    laedata.apm = 60000*float32(data[i].Attack)/float32(data[i].Time)
    laedata.eff = laedata.apm/laedata.lpm
    return laedata

}
func getjson(name string)  []TeaData{
    var data []TeaData
    params := url.Values{}
    Url, err := url.Parse("http://121.4.147.128:8888/getProfile")
    if err != nil {
        fmt.Println("some error")
    }
    params.Set("id",name)
    Url.RawQuery = params.Encode()
    urlPath := Url.String()
    //fmt.Println(url)
    resp, err := http.Get(urlPath)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("some error")
    }
    return data
}
func main() {
    var data []TeaData
    var laedata LAEdata
    //var now_game int = 0
    var name string = os.Args[1]
    fmt.Printf("LPM\tAPM\tEFF\n")
    //now_game = getjson(name)[0].Idmultiplayergameresult
    data = getjson(name)
    for i:=0;i<15;i++{
    laedata = calc_data(data,i)
    fmt.Printf("%0.3f\t%0.3f\t%0.3f\n",laedata.lpm,laedata.apm,laedata.eff)
    }
}
