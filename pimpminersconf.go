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
Interacts with this file: https://github.com/getpimp/pimpminers-conf/pimpminers.conf
==================================================================
*/
package pimpminersconf

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const REMOTE = "https://raw.githubusercontent.com/getpimp/pimpminers-conf/master/pimpminers.conf"
const LOCAL = "/tmp/pimpminers.conf"

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
	SupportedAlgos []string
}

type PimpMinerProfile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cfile string `json:"cfile"`
}

// Load returns a mapstring of PimpMiners populated with data from the pimpminers.conf file
func Load(file string) map[string][]PimpMiner {
	// if no file specified, use default in /opt/pimp.
	if file == "" {
		file = LOCAL
	}
	if FileExists(LOCAL) != "" {
		// download the file
		if err := DownloadFile(file, REMOTE); err != nil {
			panic(err)
		}
	}

	jsonFile, err := os.Open(file) // Open the JSON file
	if err != nil {                // if os.Open returns an error then handle it
		panic(err)
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

// DownloadFile will download a url to a local file.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// FileExists takes a filename string and returns it if it exists, or empty string if it does not.
func FileExists(file string) string {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		// path/to/whatever exists
		return file
	}
	return ""
}

// PrettyPrint takes a string of json in and returns a prettier string of json out.
func PrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
