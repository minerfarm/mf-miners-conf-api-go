/*==================================================================
       .__
______ |__| _____ ______  Portable Instant Mining Platform
\____ \|  |/     \____  \       By miners, for miners.
|  |_> >  |  Y Y  \  |_> >
|   __/|__|__|_|  /   __/    Support: forum.getpimp.org
|__|            \/|__|

Copyright (c) 2019 getPiMP.org.  All Rights Reserved.
Author: Portable Instant Mining Platform, LLC
License: This code is licensed for use with PiMP only.
Description: PiMP OS pimpminers.conf API wrapper in Golang
Interacts with this file: https://github.com/melt7777/pimpminers-conf/pimpminers.conf
==================================================================
*/
package pimpminersconf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PimpMiner struct {
	Info           string             `json:"info"`
	Platform       string             `json:"platform"`
	Repotype       string             `json:"repotype"`
	Repo           string             `json:"repo"`
	Folder         string             `json:"folder"`
	Binary         string             `json:"binary"`
	Configure      string             `json:"configure"`
	Menu           string             `json:"menu"`
	Postexec       string             `json:"postexec"`
	Profiles       []PimpMinerProfile `json:"profiles"`
	MainVersion    string             
	DevelVersion   string             
	PageURL        string             
	PageURLRegex   string             
	SupportedAlgos string
}

type PimpMinerProfile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cfile string `json:"cfile"`
}

// Init returns a mapstring of PimpMiners populated with data from the pimpminers.conf file
func Init(file string) map[string][]PimpMiner {
	if file == "" {
		file = "/opt/pimp/pimpminers.conf"
	}
	jsonFile, err := os.Open(file) // Open the JSON file
	if err != nil {                // if os.Open returns an error then handle it
		fmt.Println(err)
	}
	defer jsonFile.Close()                   // defer the closing of our jsonFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(jsonFile) // read our opened json file as a byte array.
	var miners map[string][]PimpMiner        // Create stringmap of our slice of structs
	json.Unmarshal(byteValue, &miners)       // Read in the JSON
	return miners
}

// GetMiner returns a PimpMiner object with the specified name from the provided mapstring.
// Note: This is for reading, not setting values.
func GetMiner(name string, miners map[string][]PimpMiner) PimpMiner {
	for _, v := range miners {
		if v[0].Info == name {
			return v[0]
		}
	}
	return PimpMiner{}
}

// JsonPrettyPrint takes a string of json in and returns a prettier string of json out.
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
