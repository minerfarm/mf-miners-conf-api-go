## Overview

Copyright (c) 2019 getPiMP.org.  All Rights Reserved.
Author: Portable Instant Mining Platform, LLC
License: This code is licensed for use with PiMP only.
Description: PiMP OS pimpminers.conf API wrapper in Golang
Interacts with this file: https://github.com/melt7777/pimpminers-conf/pimpminers.conf

## Example uses:

Output the RepoType field by Profile #: 

`fmt.Println(miners["240"][0].Repotype)`

Set a field:

```
m := GetMiner("bminer", miners)
m.Repo = "TestRepo"
```

Iterate over the miners:

```
for k, v := range miners {
	fmt.Printf("(Profile %s) %s\n", k, v[0].Info)
}
```