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

package service

import (
	"context"
	"log"
	"time"

	"github.com/amreo/ercole-hostdata-dataservice/config"
	"github.com/amreo/ercole-hostdata-dataservice/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceInterface interface {
	UpdateHostInfo(hostdata model.HostData) error
}

type Service struct {
}

var Client *mongo.Client

func SetupDatabase() {
	Client = ConnectMongodb()
	log.Println("Connected to MongoDB!", config.Config.Mongodb.URI)

}

func ConnectMongodb() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Config.Mongodb.URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func SaveHostData(hostdata model.HostData) error {
	hostdata.ServerVersion = "latest"
	hostdata.Archived = false
	hostdata.CreatedAt = time.Now()

	collection := Client.Database(config.Config.Mongodb.DBName).Collection("hosts")
	res, err := collection.UpdateOne(context.TODO(), bson.D{
		{"hostname", hostdata.Hostname},
		{"archived", false},
	}, bson.D{
		{"$set", bson.D{
			{"archived", true},
		}},
	})
	log.Println(res)
	if err != nil {
		return err
	}

	// var res model.HostData
	// collection.FindOne(context.TODO(), bson.D{
	// 	{"Hostname", hostdata.Hostname},
	// 	{"Archived", false},
	// }).Decode(res)
	// log.Println("---")
	// log.Println(utils.ToJson(hostdata))
	_, err = collection.InsertOne(context.TODO(), hostdata)
	return err
}
