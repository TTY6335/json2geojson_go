package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
)

func main() {
/** JSONデコード用に構造体定義 */
//}

type Full_data struct {
	Timestamp float64 `json:"now"`
	Messages  uint32 `json:"messages"`
	All_flightdata []struct {
		Hex                string `json:"hex"`
		Flight             string `json:"flight,omitempty"`
		Alt_baro              int `json:"alt_baro"`
		Alt_geom              int `json:"alt_geom,omitempty"`
		Gs                float32 `json:"gs,omitempty"`
		Ias                   int `json:"ias,omitempty"`
		Tas                   int `json:"tas,omitempty"`
		Mach              float32 `json:"mach,omitempty"`
		Track             float32 `json:"track,omitempty"`
		Track_rate        float32 `json:"track_rate,omitempty"`
		Roll              float32 `json:"roll,omitempty"`
		Mag_heading       float32 `json:"mag_heading,omitempty"`
		Baro_rate             int `json:"mag_heading,omitempty"`
		Geom_rate             int `json:"geom_rate,omitempty"`
		Squawk             string `json:"squawk,omitempty"`
		Emergency          string `json:"emergency,omitempty"`
		Category           string `json:"category,omitempty"`
		Nav_qnh           float32 `json:"nav_qnh,omitempty"`
		Nav_altitude_mcp      int `json:"nav_altitude_mcp,omitempty"`
		Nav_altitude_fms      int `json:"nav_altitude_fms,omitempty"`
		Nav_heading       float32 `json:"nav_heading,omitempty"`
		Lat               float32 `json:"lat,omitempty"`
		Lon               float32 `json:"lon,omitempty"`
		Nic                   int `json:"nic,omitempty"`
		Rc                    int `json:"rc,omitempty"`
		Seen_pos          float32 `json:"seen_pos,omitempty"`
		Version           float32 `json:"version,omitempty"`
		Nic_baro              int `json:"nic_baro,omitempty"`
		Nac_p                 int `json:"nac_p,omitempty"`
		Nac_v                 int `json:"nac_v,omitempty"`
		Sil                   int `json:"sil,omitempty"`
		Sil_type           string `json:"sil_type,omitempty"`
		Gva	                  int `json:"gva,omitempty"`
		Sda	                  int `json:"sda,omitempty"`
		Messages              int `json:"messages,omitempty"`
		Seen              float32 `json:"seen,omitempty"`
		Rssi              float32 `json:"rssi,omitempty"`
	}	`json:"aircraft"`
}

// JSONファイル読み込み
bytes, err := ioutil.ReadFile("./aircraft.json")
if err != nil {
log.Fatal(err)
}
// JSONデコード
var flight_data Full_data
if err := json.Unmarshal(bytes, &flight_data); err != nil {
log.Fatal(err)
}


type Properties struct {
Hex                string `json:"hex"`
Flight             string `json:"flight,omitempty"`
Alt_baro              int `json:"alt_baro"`
Alt_geom              int `json:"alt_geom,omitempty"`
Gs                float32 `json:"gs,omitempty"`
Ias                   int `json:"ias,omitempty"`
Tas                   int `json:"tas,omitempty"`
Mach              float32 `json:"mach,omitempty"`
Track             float32 `json:"track,omitempty"`
Track_rate        float32 `json:"track_rate,omitempty"`
Roll              float32 `json:"roll,omitempty"`
Mag_heading       float32 `json:"mag_heading,omitempty"`
Baro_rate             int `json:"mag_heading,omitempty"`
Geom_rate             int `json:"geom_rate,omitempty"`
Squawk             string `json:"squawk,omitempty"`
Emergency          string `json:"emergency,omitempty"`
Category           string `json:"category,omitempty"`
Nav_qnh           float32 `json:"nav_qnh,omitempty"`
Nav_altitude_mcp      int `json:"nav_altitude_mcp,omitempty"`
Nav_altitude_fms      int `json:"nav_altitude_fms,omitempty"`
Nav_heading       float32 `json:"nav_heading,omitempty"`
Nic                   int `json:"nic,omitempty"`
Rc                    int `json:"rc,omitempty"`
Seen_pos          float32 `json:"seen_pos,omitempty"`
Version           float32 `json:"version,omitempty"`
Nic_baro              int `json:"nic_baro,omitempty"`
Nac_p                 int `json:"nac_p,omitempty"`
Nac_v                 int `json:"nac_v,omitempty"`
Sil                   int `json:"sil,omitempty"`
Sil_type           string `json:"sil_type,omitempty"`
Gva	                  int `json:"gva,omitempty"`
Sda	                  int `json:"sda,omitempty"`
Messages              int `json:"messages,omitempty"`
Seen              float32 `json:"seen,omitempty"`
Rssi              float32 `json:"rssi,omitempty"`
}
type Geojson_data struct {
Geometry struct {
Coordinates []float64 `json:"coordinates"`
Type        string    `json:"type"`
} `json:"geometry"`
Properties `json:"properties"`
Type string `json:"type"`
}


var geojson_flight_data Geojson_data
var geojson_properties Properties

// デコードしたデータを表示
for _, single_flight := range flight_data.All_flightdata {
	geojson_properties.Hex              =single_flight.Hex
	geojson_properties.Flight           =single_flight.Flight
	geojson_properties.Alt_baro         =single_flight.Alt_baro
	geojson_properties.Alt_geom         =single_flight.Alt_geom
	geojson_properties.Gs               =single_flight.Gs
	geojson_properties.Ias              =single_flight.Ias
	geojson_properties.Tas              =single_flight.Tas
	geojson_properties.Mach             =single_flight.Mach
	geojson_properties.Track            =single_flight.Track
	geojson_properties.Track_rate       =single_flight.Track_rate
	geojson_properties.Roll             =single_flight.Roll
	geojson_properties.Mag_heading      =single_flight.Mag_heading
	geojson_properties.Baro_rate        =single_flight.Baro_rate
	geojson_properties.Geom_rate        =single_flight.Geom_rate
	geojson_properties.Squawk           =single_flight.Squawk
	geojson_properties.Emergency        =single_flight.Emergency
	geojson_properties.Category         =single_flight.Category
	geojson_properties.Nav_qnh          =single_flight.Nav_qnh
	geojson_properties.Nav_altitude_mcp =single_flight.Nav_altitude_mcp
	geojson_properties.Nav_altitude_fms =single_flight.Nav_altitude_fms
	geojson_properties.Nav_heading      =single_flight.Nav_heading
	geojson_properties.Nic              =single_flight.Nic
	geojson_properties.Rc               =single_flight.Rc
	geojson_properties.Seen_pos         =single_flight.Seen_pos
	geojson_properties.Version          =single_flight.Version
	geojson_properties.Nic_baro         =single_flight.Nic_baro
	geojson_properties.Nac_p            =single_flight.Nac_p
	geojson_properties.Nac_v            =single_flight.Nac_v
	geojson_properties.Sil              =single_flight.Sil
	geojson_properties.Sil_type         =single_flight.Sil_type
	geojson_properties.Gva	            =single_flight.Gva
	geojson_properties.Sda	            =single_flight.Sda
	geojson_properties.Messages         =single_flight.Messages
	geojson_properties.Seen             =single_flight.Seen
	geojson_properties.Rssi             =single_flight.Rssi
	fmt.Printf("%s : %s\n", geojson_properties.Hex, geojson_properties.Flight)
	geojson_flight_data.Properties=geojson_properties
}
}
