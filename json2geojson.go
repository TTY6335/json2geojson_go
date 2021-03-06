package main

import (
	"./src"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
	//"./setting"
)

//このへん参照
//https://qiita.com/mizukmb/items/0441ab2afbeaa29faec9
func EncodingJSON(geojson format.Geojson_data) []byte {
	bdata, err := json.Marshal(geojson)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bdata
}

var timeout = time.Duration(500 * time.Millisecond)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

//json 2 geojson format
func convert_json_format(flight_data format.Full_data) format.Geojson_data {
	var All_geojson_flight_data format.Geojson_data
	var single_geojson_flight_data format.Geojson_Features

	// geojsonに投入
	for _, single_flight := range flight_data.All_flightdata {
		//Geometry
		single_geojson_flight_data.Geometry.Coordinates[0] = single_flight.Lon
		single_geojson_flight_data.Geometry.Coordinates[1] = single_flight.Lat
		single_geojson_flight_data.Geometry.Type = "Point"
		//Type
		single_geojson_flight_data.Type = "Feature"
		//Properties
		single_geojson_flight_data.Properties.Timestamp = flight_data.Timestamp
		single_geojson_flight_data.Properties.Hex = single_flight.Hex
		single_geojson_flight_data.Properties.Flight = single_flight.Flight
		single_geojson_flight_data.Properties.Alt_baro = single_flight.Alt_baro
		single_geojson_flight_data.Properties.Alt_geom = single_flight.Alt_geom
		single_geojson_flight_data.Properties.Gs = single_flight.Gs
		single_geojson_flight_data.Properties.Ias = single_flight.Ias
		single_geojson_flight_data.Properties.Tas = single_flight.Tas
		single_geojson_flight_data.Properties.Mach = single_flight.Mach
		single_geojson_flight_data.Properties.Track = single_flight.Track
		single_geojson_flight_data.Properties.Track_rate = single_flight.Track_rate
		single_geojson_flight_data.Properties.Roll = single_flight.Roll
		single_geojson_flight_data.Properties.Mag_heading = single_flight.Mag_heading
		single_geojson_flight_data.Properties.Baro_rate = single_flight.Baro_rate
		single_geojson_flight_data.Properties.Geom_rate = single_flight.Geom_rate
		single_geojson_flight_data.Properties.Squawk = single_flight.Squawk
		single_geojson_flight_data.Properties.Emergency = single_flight.Emergency
		single_geojson_flight_data.Properties.Category = single_flight.Category
		single_geojson_flight_data.Properties.Nav_qnh = single_flight.Nav_qnh
		single_geojson_flight_data.Properties.Nav_altitude_mcp = single_flight.Nav_altitude_mcp
		single_geojson_flight_data.Properties.Nav_altitude_fms = single_flight.Nav_altitude_fms
		single_geojson_flight_data.Properties.Nav_heading = single_flight.Nav_heading
		single_geojson_flight_data.Properties.Nic = single_flight.Nic
		single_geojson_flight_data.Properties.Rc = single_flight.Rc
		single_geojson_flight_data.Properties.Seen_pos = single_flight.Seen_pos
		single_geojson_flight_data.Properties.Version = single_flight.Version
		single_geojson_flight_data.Properties.Nic_baro = single_flight.Nic_baro
		single_geojson_flight_data.Properties.Nac_p = single_flight.Nac_p
		single_geojson_flight_data.Properties.Nac_v = single_flight.Nac_v
		single_geojson_flight_data.Properties.Sil = single_flight.Sil
		single_geojson_flight_data.Properties.Sil_type = single_flight.Sil_type
		single_geojson_flight_data.Properties.Gva = single_flight.Gva
		single_geojson_flight_data.Properties.Sda = single_flight.Sda
		single_geojson_flight_data.Properties.Messages = single_flight.Messages
		single_geojson_flight_data.Properties.Seen = single_flight.Seen
		single_geojson_flight_data.Properties.Rssi = single_flight.Rssi
		//	fmt.Println(single_geojson_flight_data)
		All_geojson_flight_data.Features = append(All_geojson_flight_data.Features, single_geojson_flight_data)
	}
	All_geojson_flight_data.Type = "FeatureCollection"

	return All_geojson_flight_data
}

func main() {
        http.DefaultTransport.(*http.Transport).MaxIdleConns = 0
        http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 3000
        http.DefaultTransport.(*http.Transport).DisableKeepAlives = true
//      for range time.Tick(1 * time.Second) {
        for range time.Tick(1000 * time.Millisecond) {
//              go do_all()
                go func(){
                        //POST先のURL
                        const URL = url.Post_target
                        //const URL = "http://localhost:8000"
                        // JSONファイル読み込み
                        json_bytes, err := ioutil.ReadFile("/var/run/dump1090-fa/aircraft.json")
                        //json_bytes, err := ioutil.ReadFile("./sample.json")

                        if err != nil {
                                log.Fatal(err)
                        }
                        // JSONデコード
                        var flight_data format.Full_data
                        if err := json.Unmarshal(json_bytes, &flight_data); err != nil {
                                log.Fatal(err)
                        }

                        //GeoJSONに変換
                        var All_geojson_flight_data format.Geojson_data
                        All_geojson_flight_data = convert_json_format(flight_data)
                        //      fmt.Println(All_geojson_flight_data)

                        bdata := EncodingJSON(All_geojson_flight_data)
                        //POST
                        res, err := http.Post(URL, "application/json", bytes.NewBuffer(bdata))
                        if res != nil {
                                //  ここで nilでなければ閉じる
                                defer res.Body.Close()
                        }

                        if err != nil {
                                //IF TIMEOUT save geojson
                                nowUTC := time.Now().UTC()
                                millisecond := nowUTC.UnixNano() / int64(time.Millisecond)
                                geojson_name := "/home/gunma/workplace/adsb/" + strconv.FormatInt(millisecond, 10) + ".geojson"
                                ioutil.WriteFile(geojson_name, []byte(bdata), os.ModePerm)
                                fmt.Println("[!] " + err.Error())
                        } else {
                                fmt.Println("[*] " + res.Status)
                        }

                }()
        }
}
