package main

import (
	"fmt"
	"time"
)

type User struct {
	userName string
	age      int
	email    string
}

func main() {

	start := time.Now()

	userData := getDetails()

	fmt.Println(userData)

	fmt.Println(time.Now().Sub(start))

}

func getDetails() User {

	name := make(chan string)
	age := make(chan int)
	email := make(chan string)

	go getName(name)
	go getAge(age)
	go getEmail(email)

	return User{
		userName: <-name,
		age:      <-age,
		email:    <-email,
	}
}

func getName(ch chan string) {
	name := "Athun Lal"
	time.Sleep(time.Second * 1)

	ch <- name
	close(ch)
}

func getAge(ch chan int) {
	age := 23
	time.Sleep(2 * time.Second)
	ch <- age
	close(ch)

}

func getEmail(ch chan string) {
	email := "athunlalp@gamil.com"
	time.Sleep(3 * time.Second)
	ch <- email
	close(ch)

}
