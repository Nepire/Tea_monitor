package main
import (
    "encoding/json"
    "net/http"
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
    lpm int
    apm int
    eff float32
}
func calc_data(data []TeaData,i int) LAEdata{
    var laedata LAEdata
    laedata.lpm = 24000*data[i].Pieces/data[i].Time
    laedata.apm = 60000*data[i].Attack/data[i].Time
    laedata.eff = float32(laedata.apm)/float32(laedata.lpm)
    return laedata

}
func getjson(name string)  []TeaData{
    var data []TeaData
    var url string = "http://139.199.75.237:8888/getProfile?id="
    url += name//os.Args[1]
    resp, err := http.Get(url)
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
    var now_game int = 0
    var name string = os.Args[1]
    fmt.Printf("LPM\tAPM\tEFF\n")
    now_game = getjson(name)[0].Idmultiplayergameresult
    for true {
        data = getjson(name)
        if data[0].Idmultiplayergameresult!=now_game{
            now_game = data[0].Idmultiplayergameresult
            laedata = calc_data(data,0)
            fmt.Printf("%d\t%d\t%0.4f\n",laedata.lpm,laedata.apm,laedata.eff)
        }
    }
}
