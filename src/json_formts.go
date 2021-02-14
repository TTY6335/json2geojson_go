package format
/** JSONデコード用に構造体定義 */
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

//geojson用の構造体
type Geojson_Geometry struct {
	Coordinates [2]float32 `json:"coordinates"`
	Type            string    `json:"type"`
}

type Geojson_Features struct{
	Type         string            `json:"type"`
	Geometry     Geojson_Geometry  `json:"geometry"`
	Properties   Geojson_Properties`json:"properties"`
}

type Geojson_Properties struct {
Timestamp         float64 `json:"timestamp"`
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
	Type string `json:"type"`
	Features []Geojson_Features`json:"features"`
}

