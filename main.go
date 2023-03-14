package main

import (
	"fmt"

	aesdecrypt "github.com/Onboardbase/go-cryptojs-aes-decrypt/decrypt"
)

func main () {
	cipher := "U2FsdGVkX1/cJw+ISoZKbICgFmoJ+Ehx/atOTgPlt4IOCpmgrHJBY9J2gvoupPbNQs3dHsshX2EljZPlnVLATRJRyF8rOMui+uFMl8DV9I0VVA+1IqecvcP+a2gswGSjYEOiEG9kKQQ8cNlZPtGZdD4ksxxKxm/DhDxquzYX9mWwtMLe0EVYqFd0gncHX2wYzOrvKbMLZlGn/rN70XC+kQgEd47syLFOBm/cjzmVFGvymJfvuRNDl66PRzMtoGqZkQ7uufZuFS3rRUpcsiVv+0XbkSTiaZeD2t7cs+pn4G4uAymMJathk2z2pmhhd184rnXpuooaz8pXPmfdL8CwbQ=="

	
    // The passphrase that generated the cipher above
    passphrase := "passcode"
	decrypted, err := aesdecrypt.Run(cipher, passphrase)
	fmt.Println(decrypted, err)
	// var data interface{}
	// err := json.Unmarshal([]byte(decrypted), &data)
	// fmt.Println(data)
	// fmt.Println(err)
	// Outputs -> '{"id":"0a09a9e1-eee2-4a8e-bf77-901be943621c","key":"WHAT","value":"IS_THIS_NAH","comment":"","addedBy":{"name":"Aleem Isiaka","id":"a3748a40-3ab8-4dd1-bac6-e73c1047c95f"},"addedDate":"April 22, 2022 | 20:09 GMT","method":""}'
}