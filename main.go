package main


import (
	"fmt"
)

var ciphers = [...]string{
	"Caesar Cipher",
	"Rot13",
}

func getInputString() string {
	var text string
	fmt.Print("\nEnter the string you want to run through GoCrypt: ")
	fmt.Scanf("%s\n", &text)

	return text
}

func getCipherMethod() string {
	var method string
	fmt.Print("\nDo you want to encrypt or decrypt? (E/d): ")
	fmt.Scanf("%1v\n", &method)

	return method
}

func main() {
	fmt.Print(
		"\n\n",
		"     ********             ******                            **    \n",
		"    **//////**           **////**         **   ** ******   /**    \n",
		"   **      //   ******  **    //  ****** //** ** /**///** ******  \n",
		"  /**          **////**/**       //**//*  //***  /**  /**///**/   \n",
		"  /**    *****/**   /**/**        /** /    /**   /******   /**    \n",
		"  //**  ////**/**   /**//**    ** /**      **    /**///    /**    \n",
		"   //******** //******  //****** /***     **     /**       //**   \n",
		"    ////////   //////    //////  ///     //      //         //    \n",
		"\n\n",
		"Welcome to GoCrypt\n",
		"Encrypt and Decrypt anything you want.\n",
	)

	fmt.Println("\nList over what you can use:")

	for i, v := range ciphers {
		fmt.Println("(", i, ") - ", v)
	}


	var cipher int
	for {
		fmt.Print("Enter the id of what you want to use: ")
		fmt.Scanf("%d\n", &cipher)

		if cipher < len(ciphers) && cipher >= 0 {
			break
		}
	}

	textIn := getInputString()
	method := getCipherMethod()

	var textOut string

	switch ciphers[cipher] {
	case "Caesar Cipher":
		textOut = caesar(textIn, method)
	case "Rot13":
		textOut = caesar(textIn, method)
	}

	fmt.Println("\nHere is the output of the cipher:", textOut)
}
