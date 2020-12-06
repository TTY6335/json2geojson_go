package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
)

func main() {
/** JSONデコード用に構造体定義 */
type Single_flightdata struct {
		hex                string `json:"hex"`
		flight             string `json:"flight"`
//		alt_baro              int `json:"alt_baro"`
//		alt_geom              int `json:alt_geom`
//		gs                float32 `json:gs`
//		ias                   int `json:ias`
//		tas                   int `json:tas`
//		mach              float32 `json:mach`
//		track             float32 `json:track`
//		track_rate        float32 `json:track_rate`
//		roll              float32 `json:roll`
//		mag_heading       float32 `json:mag_heading`
//		baro_rate             int `json:mag_heading`
//		geom_rate             int `json:geom_rate`
//		squawk                int `json:squawk`
//		emergency          string `json:emergency`
//		category           string `json:category`
//		nav_qnh           float32 `json:nav_qnh`
//		nav_altitude_mcp      int `json:nav_altitude_mcp`
//		nav_altitude_fms      int `json:nav_altitude_fms`
//		nav_heading       float32 `json:nav_heading`
//		lat               float32 `json:lat`
//		lon               float32 `json:lon`
//		nic                   int `json:nic`
//		rc                    int `json:rc`
//		seen_pos          float32 `json:seen_pos`
//		version           float32 `json:version`
//		nic_baro              int `json:nic_baro`
//		nac_p                 int `json:nac_p`
//		nac_v                 int `json:nac_v`
//		sil                   int `json:sil`
//		sil_type           string `json:sil_type`
//		gva	                  int `json:gva`
//		sda	                  int `json:sda`
//		messages              int `json:messages`
//		seen              float32 `json:seen`
//		rssi              float32 `json:rssi`
}

type Full_data struct {
Timestamp float64 `json:"now"`
Messages  uint32 `json:"messages"`
All_flightdata []*Single_flightdata `json:"aircraft"`
}

// JSONファイル読み込み
bytes, err := ioutil.ReadFile("sample.json")
if err != nil {
log.Fatal(err)
}
// JSONデコード
var flight_data Full_data
if err := json.Unmarshal(bytes, &flight_data); err != nil {
log.Fatal(err)
}
fmt.Println(flight_data.Timestamp)
fmt.Println(flight_data.Messages)
fmt.Println(flight_data.All_flightdata)

// デコードしたデータを表示
//for _, single_flight := range flight_data.Single_flight {
//fmt.Println(single_flight.hex)
//fmt.Println(single_flight.flight)
//fmt.Printf("%d : %d\n", single_flight.hex, single_flight.flight)
//}
}
