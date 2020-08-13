package main
// More info on Getwd()
// https://golang.org/src/os/getwd.go
import(
  "os" 
  "fmt"
  "log"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}