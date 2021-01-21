package main
import (
	"os"
	"fmt"
	"time"
	"strconv"
	"strings"
	"github.com/briandowns/spinner"
)
var (
	h int
	m int
	s int
	flag = 0
	terminator = 0
)
// timer hms
func main(){
	var argv = len(os.Args)
	if argv != 2 {
		Usage()
		return
	}
	var args = os.Args[1]
	for index :=0;index <len(args);index++{
		//fmt.Printf("str[%d]=%c\n",index,args[index])
		if strings.EqualFold(string(args[index]), "h") {
			h,_ = strconv.Atoi(string(args[flag : index]))
			flag = index+1
		} else if strings.EqualFold(string(args[index]), "m") {
			m,_ = strconv.Atoi(string(args[flag : index]))
			flag = index+1
		} else if strings.EqualFold(string(args[index]), "s") {
			s,_ = strconv.Atoi(string(args[flag : index]))
		}
	}
	terminator = h * 60 * 60 + m * 60 + s
	fmt.Printf("Counting on <<<%d:%d:%d\n",h,m,s)
	var tickerStopped = false
	ticker := time.NewTicker(1 * time.Second)

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond, spinner.WithWriter(os.Stderr)) // Build our new spinner
	s.FinalMSG = fmt.Sprintf("time‘s up\n")
	s.Start() // Start the spinner

	go func() {
		//退出条件
		time.Sleep(time.Second * time.Duration(terminator))
		ticker.Stop()
		tickerStopped = true
	}()

	start := time.Now()
	for {
		if tickerStopped {
			break
		}
		<- ticker.C
		elapsed := time.Since(start)
		s.Suffix = fmt.Sprintf(" >>> %v",elapsed)
	}
	s.Stop()
}

// Usage ... app user guide
func Usage() {
	fmt.Println("Version: 20201210")
	fmt.Println("Usage: e.g: timer 10h20m12s")
	fmt.Println(" 'h' is Hour --- 'm' is Minute --- 's' is Seconds")
}
