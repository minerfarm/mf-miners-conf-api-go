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
	"testing"
)

func TestLoad(t *testing.T) {

	miners := Load("/tmp/pimpminers.conf") // integrate with pimpminers.conf
	var actualResult = miners["240"][0].Repotype
	var expectedResult = "binary"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
