package main

import (
	"fmt"
	"math/rand"
	"vinLink/web/Link"
	"vinLink/web/Storage"
)

func main() {
	linkList := Storage.GetData()

	link := Link.New(fmt.Sprintf("test.com/%v", rand.Intn(5000)))

	linkList[link.Id()] = link.Link()


	Storage.SaveData(linkList)

	fmt.Println("done")
}
