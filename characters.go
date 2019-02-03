/* 
	Created By Andy Panana
*/


package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/hex"
	"strconv"
	"crypto/md5"
	"net/url"
	"time"
	"bufio"
	"strings"
)

type Response struct {
	Code		int			`json:"code"`
	Data		DataStruct	`json:"data"`
}


type DataStruct struct {
	Total		int			`json:"total"`
	Character	[]Character	`json:"results"`
}

type Character struct {
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Released	string		`json:"modified"`
	URI			string		`json:"resourceURI"`
}

func show(url string) {

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData,err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)


	for i := 0; i < len(responseObject.Data.Character); i++ {
		fmt.Println("**************"+ strconv.Itoa(i+1) +"***************")
		fmt.Print("Name : ")
		fmt.Println(responseObject.Data.Character[i].Name)
		fmt.Print("Description : ")
		fmt.Println(responseObject.Data.Character[i].Description)
		fmt.Print("Modified : ")
		fmt.Println(responseObject.Data.Character[i].Released)
		fmt.Print("URI : ")
		fmt.Println(responseObject.Data.Character[i].URI)
	} 
}

func listar() {
	ts := time.Now().UnixNano()
	ts_str := strconv.FormatInt(ts, 10)
	pubKey := "d439c2568cb713ad5363afcb099c2392"
	privKey := "ca35f215624121b951fa3ffe7fc0997355c99b70"
	hash := md5.Sum([]byte(ts_str + privKey + pubKey))

	result := hex.EncodeToString(hash[:])

	url := "http://gateway.marvel.com/v1/public/characters?ts="+ts_str+"&apikey="+pubKey+"&hash="+result

	show(url)
	/*
	response, err := http.Get(url)

	 if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData,err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)


	for i := 0; i < len(responseObject.Data.Character); i++ {
		fmt.Println("**************"+ strconv.Itoa(i+1) +"***************")
		fmt.Println(responseObject.Data.Character[i].Name)
		fmt.Println(responseObject.Data.Character[i].Description)
		fmt.Println(responseObject.Data.Character[i].Released)
		fmt.Println(responseObject.Data.Character[i].URI)
		fmt.Println(" ");
	}  */

}

func buscar() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese el nombre: ")
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	name = strings.TrimSpace(name)

	name_url, _ := url.Parse(name)


	ts := time.Now().UnixNano()
	ts_str := strconv.FormatInt(ts, 10)
	pubKey := "d439c2568cb713ad5363afcb099c2392"
	privKey := "ca35f215624121b951fa3ffe7fc0997355c99b70"
	hash := md5.Sum([]byte(ts_str + privKey + pubKey))

	result := hex.EncodeToString(hash[:])

	url := "http://gateway.marvel.com/v1/public/characters?name="+name_url.EscapedPath()+"&ts="+ts_str+"&apikey="+pubKey+"&hash="+result

	show(url)

	/*
	response, err := http.Get(url)

	 if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData,err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)



	for i := 0; i < len(responseObject.Data.Character); i++ {
		fmt.Println("**************"+ strconv.Itoa(i+1) +"***************")
		fmt.Println(responseObject.Data.Character[i].Name)
		fmt.Println(responseObject.Data.Character[i].Description)
		fmt.Println(responseObject.Data.Character[i].Released)
		fmt.Println(responseObject.Data.Character[i].URI)
		fmt.Println(" ");
	}  */

}

func menu() {
	fmt.Println("\t*************************")
	fmt.Println("\t*   MARVEL API CON GO   *")
	fmt.Println("\t*************************\n")
	fmt.Println("1. Buscar heroe")
	fmt.Println("2. Listar los primeros 20 heroes")
	fmt.Println("3. Salir")
	fmt.Println("Ingrese Opción: ")
}

func main() {
	var opc int
	var input string
	for opc != 3 {
		menu()
		fmt.Scanf("%d\n",&opc)
		if opc == 1 {
			buscar()
		}else {
			if opc == 2 {
				listar()
			} else {
				if opc == 3 {
					fmt.Println("Cerrando el programa")
					break
				}
			}
		}
		
		fmt.Println("Press 'Enter' to continue...")
		fmt.Scanln(&input)
	}

}