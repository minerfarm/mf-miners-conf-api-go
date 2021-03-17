# mf-minersconf-go

Package mfminersconf is the API wrapper for interacting with the mf-miners.conf
json library. Copyright (c) 2021 miner.farm. All Rights Reserved. Author:
Portable Instant Mining Platform, LLC License: This code is licensed for use
with miner.farm only.

    ._____.___ ._______._______  .________
    :         |:_ ____/: .___  \ |    ___/		   Miner.Farm OS
    |   \  /  ||   _/  | :   |  ||___    \
    |   |\/   ||   |   |     :  ||       /	  Support: forum.getpimp.org
    |___| |   ||_. |    \_. ___/ |__:___/ 	Copyright (c) 2021 getPiMP.org
          |___|  :/       :/        :
                 :        :

Interacts with this Git repo: <https://github.com/minerfarm/mf-minersconf>

## Usage

### Example uses

Output the Repotype field by Profile #:

```go
fmt.Println(miners["240"][0].Repotype)
```

Set a field:

```go
m := GetMiner("bminer", miners)
m.Repo = "TestRepo"
```

Iterate over the miners:

```go
for k, v := range miners {
	fmt.Printf("(Profile %s) %s\n", k, v[0].Info)
}
```

### Functions

#### func  Clone

```go
func Clone() *git.Repository
```
Clone will clone the mf-miners-conf repo to /tmp/mf-miners-conf.

#### func  Commit

```go
func Commit(r *git.Repository, msg string) bool
```
Commit will commit the mf-miners-conf repo to git. (Maintainers only.) Returns
true if success.

#### func  DownloadFile

```go
func DownloadFile(filepath string, url string) error
```
DownloadFile will download a url to a stagingFile file.

#### func  FileExists

```go
func FileExists(file string) string
```
FileExists takes a filename string and returns it if it exists, or empty string
if it does not.

#### func  Load

```go
func Load(file string) map[string][]Miner
```
Load returns a mapstring of Miners populated with data from the mf-miners.conf
file

#### func  PrettyPrint

```go
func PrettyPrint(in string) string
```
PrettyPrint takes a string of json in and returns a prettier string of json out.

#### func  RunCommand

```go
func RunCommand(cmd string) string
```
RunCommand runs the specified string as a command in a bash shell, and returns
its output as a string.

#### func  Save

```go
func Save(jsonfile string, m map[string][]Miner)
```
Save takes a map of strings to Miners object, marshals it into json, and then
saves it as a file.

#### func  SetRepo

```go
func SetRepo(name string, repo string, miners map[string][]Miner) string
```
SetRepo updates a Miner object's repo property with the specified filename.
Note: This is for pimpup 2.x.

#### type Miner

```go
type Miner struct {
	Info           string              `json:"info"`
	Platform       string              `json:"platform"`
	Repotype       string              `json:"repotype"`
	Repo           string              `json:"repo"`
	Folder         string              `json:"folder"`
	Binary         string              `json:"binary"`
	Configure      string              `json:"configure"`
	Menu           string              `json:"menu"`
	Postexec       string              `json:"postexec"`
	Profiles       []MinerProfile      `json:"profiles"`
	MainVersion    string              `json:"MainVersion"`
	DevelVersion   string              `json:"DevelVersion"`
	PageURL        string              `json:"PageURL"`
	PageURLRegex   string              `json:"PageURLRegex"`
	SupportedAlgos []map[string]string `json:"SupportedAlgos"`
	BtcTalk        string              `json:"BtcTalk"`
}
```

Miner describes all of the attributes of each miner which is found in the
mf-miners.conf json library.

#### func  GetMiner

```go
func GetMiner(name string, miners map[string][]Miner) Miner
```
GetMiner returns a Miner object with the specified name from the provided
mapstring. Note: This is for reading, not setting values.

#### type MinerProfile

```go
type MinerProfile struct {
	ID     string               `json:"id"`
	Name   string               `json:"name"`
	Cfile  string               `json:"cfile"`
	Config []MinerProfileConfig `json:"Config"`
}
```

MinerProfile is a pcfg file that configures a Miner.

#### type MinerProfileConfig

```go
type MinerProfileConfig struct {
	Flags      string `json:"FLAGS"`
	CONF       string `json:"CONF"`
	API        string `json:"API"`
	POOL_TITLE string `json:"POOL_TITLE"`
	TYPE       string `json:"TYPE"`
	Extra      string `json:"EXTRA"`
	Notes      string `json:"NOTES"`
	Template   string `json:"TEMPLATE"`
}
```

MinerProfileConfig is a series of template and control metadata which allows
confgen to build the pcfg files.
