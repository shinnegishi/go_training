
package main

import (
    "fmt"
    "os/exec"
	"time"
	"log"
)

func main() {
	out, err := exec.Command("cmd", "/C", "echo hello > test.txt").Output()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println(out)
	}
    fmt.Println("処理開始: ", time.Now().Format("15:04:05"))
    cmd := exec.Command("sleep", "5s")
    cmd.Start()
    fmt.Println("sleep中: ", time.Now().Format("15:04:05"))
    cmd.Wait()
    fmt.Println("sleep終了: ", time.Now().Format("15:04:05"))
}