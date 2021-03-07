package main

import (
	"fmt"
)

func main() {
	var data = "This is Miran"
	message(data)
	fmt.Println(message)
}

func message() {

	var message []string
	for i := 0; i < 5; i++ {
		message[i] = i
		//city[i] = fmt.Sprint(i)
		//fmt.Println(message)
		//fmt.Println(i)
		return 
	}

}

/*func message()
	len_deta := len(message)
	i := 0
	for _, data := range message {
		if data != " " {
			if i%2 == 0 {
				message = message + strings.ToUpper(string(data))
			} else {
				message = message + strings.ToLower(string(data))
				i++
			}
		} else {
			message = message + strings(data)
		}
	}
	fmt.Println(data)
	fmt.Println(message)
}
*/
