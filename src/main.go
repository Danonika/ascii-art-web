package banner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Get(x, s string) string {
	cur := ""
	ans := [8]string{}
	Database := [1000]string{}
	s = "./files/" + s + ".txt"
	file, _ := os.Open(s)
	defer file.Close()
	reader := bufio.NewReader(file)
	// Converting data from the file to array of strings
	cnt_lines := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		Database[cnt_lines] = line[:len(line)-1]
		cnt_lines++
	}

	// making one string from arguments
	str := x
	newStr := ""
	for _, c := range str {
		if c >= rune(32) && c <= rune(126) {
			newStr += string(c)
		}
	}
	str = newStr
	Arg_slice := strings.Split(str, "\\n")

	for i := 0; i < len(Arg_slice); i++ {

		for _, char := range Arg_slice[i] {
			tmp := int(char - ' ')
			for k := 0; k < 8; k++ {
				ans[k] = ans[k] + Database[tmp*9+k+1]
			}
		}
		// Prints our answer or banner for current set of chars
		for i := range ans {
			cur += fmt.Sprintf("%s\n", ans[i])
		}

		// Formating our answer for new line
		for i := range ans {
			ans[i] = ""
		}
	}
	return cur
}
