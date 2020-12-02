package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
)

/** JSONデコード用に構造体定義 */
type Full_aircraft struct {
timestamp float32 `json:"now"`
messages      int `json:"messages"`
aircraft aircraft `json:"aircraft"`
}

type Aircraft struct {
	hex                string `json:"hex"`
	flight             string `json:flight`
	alt_baro              int `json:"alt_baro"`
	alt_geom              int `json:"alt_geom"`
	gs                float32 `json:"gs"`
	ias                   int `json:"ias"`
	tas                   int `json:"tas"`
	mach              float32  `json:"mach"`
	track             float32 `json:"track"`
	track_rate        float32 `json:"track_rate"`
	roll                  int `json:"roll"`
	mag_heading       float32 `json:"mag_heading"`
	baro_rate         float32 `json:"mag_heading"`
	geom_rate         float32 `json:"geom_rate"`
	squawk                int `json:"squawk"`
	emergency          string `json:"emergency"`
	category           string `json:"category"`
	nav_qnh               int `json:"nav_qnh"`
	nav_altitude_mcp      int `json:"nav_altitude_mcp"`
	nav_altitude_fms      int `json:"nav_altitude_fms"`
	nav_heading           int `json:"nav_heading"`
	lat               float32 `json:"lat"`
	lon               float32 `json:"lon"`
	nic                   int `json:"nic"`
	rc                    int `json:"rc"`
	seen_pos          float32 `json:"seen_pos"`
	version           float32 `json:"version"`
	nic_baro              int `json:"nic_baro"`
	nac_p                 int `json:"nac_p"`
	nac_v                 int `json:"nac_v"`
	sil                   int `json:"sil"`
	sil_type           string `json:"sil_type"`
	gva	                  int `json:"gva"`
	sda	                  int `json:"sda"`
	messages              int `json:"messages"`
	seen              float32 `json:"seen"`
	rssi              float32 `json:"rssi"`
}

func main() {
// JSONファイル読み込み
bytes, err := ioutil.ReadFile("vro.json")
if err != nil {
log.Fatal(err)
}
// JSONデコード
var persons []Person
if err := json.Unmarshal(bytes, &persons); err != nil {
log.Fatal(err)
}
// デコードしたデータを表示
for _, p := range persons {
fmt.Printf("%d : %s\n", p.Id, p.Name)
}
}
