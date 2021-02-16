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

package mfminersconf

import (
	"testing"
)

func TestLoad(t *testing.T) {

	miners := Load("/tmp/mf-miners.conf")
	var actualResult = miners["240"][0].Repotype
	var expectedResult = "binary"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
