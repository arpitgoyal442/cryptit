package main

import (

	// These 2 lines below are importing within the same Module i.e. github.com/arpit/module2
	"github.com/arpit/module2/encrypt"
	"github.com/arpit/module2/decrypt"


	



	"fmt"

)

func main(){


	fmt.Println(encrypt.Nimbus("AB"))
	fmt.Println(decrypt.Nimbus("E"))
	




}