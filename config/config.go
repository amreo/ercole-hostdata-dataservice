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

// Config contains configuration utilities, like readConfig()
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config contains the global configuration of Ercole DataService
var Config Configuration

// Configuration contains Ercole DataService configuration
type Configuration struct {
	// HttpServer contains configuration about internal http server
	HttpServer HttpServer
	// Mongodb contains configuration about database connection, some data logic and migration
	Mongodb Mongodb
}

// HttpServer contains configuration about the internal http servr
type HttpServer struct {
	// Port contains the port of the internal http server
	Port uint16
	// LogHttpRequest enable the logging of the internal http serverl
	LogHttpRequest bool
	// AgentUsername contains the username of the agent
	AgentUsername string
	// AgentPassword contains the password of the agent
	AgentPassword string
}

// Mongodb contains configuration about the database connection, some data logic and migration
type Mongodb struct {
	// URI contains MongoDB connection string/URI like 'mongodb://localhost:27017/ercole'
	URI string
	// DBName contains the name of the database
	DBName string
}

// ReadConfig read, parse and return a Configuration from the configuration file in config.json or /opt/ercole-hostdata-dataservice/config.json
// It also set the global Config with the read value
func ReadConfig() Configuration {
	var err error
	var raw []byte
	var conf Configuration

	//Read and parse config.json or /opt/ercole-hostdata-dataservice/config.json (fallback)
	if raw, err = ioutil.ReadFile("config.json"); err != nil {
		if raw, err = ioutil.ReadFile("/opt/ercole-hostdata-dataservice/config.json"); err != nil {
			log.Fatal("Unable to read configuration file", err)
		}
	}
	if err = json.Unmarshal(raw, &conf); err != nil {
		log.Fatal("Unable to parse configuration file", err)
	}

	//Set the global Config with the read value
	Config = conf

	//Return the read configuration
	return conf
}
