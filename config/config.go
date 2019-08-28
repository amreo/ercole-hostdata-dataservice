// Copyright (c) 2019 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Config Configuration

type Configuration struct {
	HttpServer HttpServer
	Mongodb    Mongodb
}

type HttpServer struct {
	Listen         string
	LogHttpRequest bool
}

type Mongodb struct {
	URI    string
	DBName string
	Debug  bool
}

func ReadConfig() Configuration {
	var err error
	var raw []byte
	var conf Configuration

	//Read and parse config.json or /app/config.json
	if raw, err = ioutil.ReadFile("config.json"); err != nil {
		if raw, err = ioutil.ReadFile("/app/config.json"); err != nil {
			log.Fatal("Unable to read configuration file", err)
		}
	}
	if err = json.Unmarshal(raw, &conf); err != nil {
		log.Fatal("Unable to parse configuration file", err)
	}
	Config = conf

	return conf
}
