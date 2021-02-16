# mfminersconf-go

```go
/*Package mfminersconf is the API wrapper for interacting with the mf-miners.conf json library.
Copyright (c) 2021 miner.farm.  All Rights Reserved.
Author: Portable Instant Mining Platform, LLC
License: This code is licensed for use with miner.farm only.
  ._____.___ ._______._______  .________
  :         |:_ ____/: .___  \ |    ___/		   Miner.Farm OS
  |   \  /  ||   _/  | :   |  ||___    \
  |   |\/   ||   |   |     :  ||       /	  Support: forum.getpimp.org
  |___| |   ||_. |    \_. ___/ |__:___/ 	Copyright (c) 2021 getPiMP.org
        |___|  :/       :/        :
               :        :
*/
```

Interacts with this Git repo: <https://github.com/minerfarm/mf-miners-conf>

## Example uses

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
