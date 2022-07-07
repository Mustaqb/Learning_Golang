package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ".,.?.!I love \n Go Programming!!!.,.?.!"

	fmt.Printf("%#v\n",strings.Contains(str,"G"))
	fmt.Printf("%#v\n",strings.ContainsAny(str,"Gaxz"))
	fmt.Printf("%#v\n",strings.ContainsRune(str,'g'))
	fmt.Printf("%#v\n",strings.Count(str,"o"))
	fmt.Printf("%#v\n",strings.Count(str,""))
	fmt.Printf("%#v\n",strings.ToLower(str))
	fmt.Printf("%#v\n",strings.ToUpper(str))

	//compare string without considering case
	fmt.Printf("%#v\n",strings.EqualFold("GOLANG","gOlAnG"))

	fmt.Printf("%#v\n",strings.Repeat("$",20))


	//Last argument mentions how many letters to replace from first
	fmt.Printf("%#v\n",strings.Replace(str," ","_",2))

	//-1 as last argument means replace all
	fmt.Printf("%#v\n",strings.Replace(str," ","_",-1))
	//instead we can use replaceAll
	fmt.Printf("%#v\n",strings.ReplaceAll(str," ","_"))

	newStr:=strings.Split(str," ")
	fmt.Printf("%#v\n",newStr)
	//split and get only the words (split including \t and \n) use Fields
	fmt.Printf("%#v\n",strings.Fields(str))
	fmt.Printf("%#v\n",str)

	fmt.Printf("%#v\n",strings.Join(newStr," "))

	//Trim .,!?
	fmt.Printf("%#v\n",strings.Trim(str,".!?,"))

}
