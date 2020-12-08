package main

import (
	"log"
	"os"

	"github.com/jpradass/PasswordGenerator/controller"
)

func main() {
	logger := log.New(os.Stdout, "[pass-gen]: ", log.LstdFlags)
	ctrller := controller.New(logger)
	ctrller.GetWordlist()
	// ctrller.SaveWordlist()

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			InsecureSkipVerify: true,
	// 		},
	// 	},
	// }
	// res, err := client.Get("https://theworld.com/~reinhold/diceware.wordlist.asc")
	// if err != nil {
	// 	panic(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode == http.StatusOK {
	// 	bb, _ := ioutil.ReadAll(res.Body)
	// 	fmt.Println(string(bb))
	// }

}
