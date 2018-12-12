package main
import (
	"fmt"
	"os"
)
func main() {
	userFile := "hanru.txt"
	fl, err := os.Open(userFile)
	defer fl.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}