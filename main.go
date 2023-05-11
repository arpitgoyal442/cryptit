package main

import (

	// These 2 lines below are importing within the same Module i.e. github.com/arpit/module2
	"github.com/arpitgoyal442/cryptit/encrypt"
	"github.com/arpitgoyal442/cryptit/decrypt"


	



	"fmt"

)

func main(){


	fmt.Println(encrypt.Nimbus("AB"))
	fmt.Println(decrypt.Nimbus("E"))
	




}