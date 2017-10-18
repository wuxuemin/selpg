package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type selpgArgs struct {
	start      int    //start page
	end        int    //end page
	pagelen    int    //number of lines in a page,default 72
	pagetype   bool   //define the end of the page
	infilename string //file name
	printdest  string //destination
}

func usageTips(err string) {
	fmt.Fprintf(os.Stderr, err+"\n"+
		`usageTips:-s startpage e endpage -l length_of_page [-f type of file] [-d dest] [filename]`)
}

func main() {
	saAddr := selpgArgs{
		start:      -1,
		end:        -1,
		pagelen:    72,
		pagetype:   false,
		infilename: "",
		printdest:  "",
	} //initial default value
	//split the arguments
	flag.IntVar(&saAddr.start, "s", -1, "Set the start page that should be set to be greater than 0.")
	flag.IntVar(&saAddr.end, "e", -1, "Set the end page that should be set to be greater than or equal to the start page.")
	flag.IntVar(&saAddr.pagelen, "l", 72, "Set the length of one page that is greater than 0.")
	flag.BoolVar(&saAddr.pagetype, "f", false, "-f means the end of the page")
	flag.StringVar(&saAddr.printdest, "d", "", "Set the destination of print.")
	flag.Parse()

	//judge if the input is valid and give tips
	if saAddr.start == -1 || saAddr.end == -1 || saAddr.start < 1 || saAddr.end < 1 {
		usageTips("the start page or the end page is  set incorrectly!")
		return
	}

	if saAddr.start > saAddr.end {
		usageTips("the start page or the end page invalid set!")
		return
	}

	if saAddr.pagelen != 72 && saAddr.pagetype == true {
		usageTips("")
		return
	}

	if len(flag.Args()) > 1 {
		usageTips("can't read more that one file!")
		return
	}

	if len(flag.Args()) == 0 {
		saAddr.infilename = ""
	} else {
		saAddr.infilename = flag.Args()[0]
	}

	if saAddr.pagetype == false {
		type1op(saAddr) //file 1 to deal with
	} else {
		type2op(saAddr) //file 2 to deal with
	}
}

// type 1 data and operate it
func type1op(saAddr selpgArgs) {
	var printer *exec.Cmd
	var printErr error
	tmpStr := fmt.Sprintf("%s", saAddr.printdest)
	printer = exec.Command("lp", "-d", tmpStr)
	stdin, printErr := printer.StdinPipe()
	if printErr != nil {
		usageTips("could not open pipe to \"" + tmpStr + "\"!")
		os.Exit(1)
	}

	// dealing with the page type
	var pageCnt int
	var lineCnt int
	var lineFlag1 *bufio.Scanner
	fout := os.Stdout
	if saAddr.infilename != "" {
		inputStream, printErr := os.Open(saAddr.infilename)
		if printErr != nil {
			usageTips("could not open pipe to \"" + tmpStr + "\"!")
			os.Exit(1)
		}
		lineFlag1 = bufio.NewScanner(inputStream)
	} else {
		lineFlag1 = bufio.NewScanner(os.Stdin)
	}
	pageCnt = 1
	lineCnt = 0
	for lineFlag1.Scan() {
		if pageCnt >= saAddr.start && lineCnt <= saAddr.end {
			fout.Write([]byte(lineFlag1.Text() + "\n"))
			stdin.Write([]byte(lineFlag1.Text() + "\n"))
		}
		lineCnt++
		if lineCnt = lineCnt % saAddr.pagelen; lineCnt == 0 {
			pageCnt++
		}
	}

	if pageCnt < saAddr.start {
		fmt.Fprintf(os.Stderr, "WARNING:start_page (%d) greater than total pages (%d)\n", saAddr.start, pageCnt)
	} else {
		if pageCnt < saAddr.end {
			fmt.Fprintf(os.Stderr, "WARING:end_page (%d) greater than total pages (%d)\n", saAddr.end, pageCnt)
		}
	}

	if saAddr.printdest != "" {
		stdin.Close()
		printer.Stdout = os.Stdout
		printer.Start()
	}

}

func type2op(saAddr selpgArgs) {
	var printer *exec.Cmd
	var printErr error
	tmpStr := fmt.Sprintf("%s", saAddr.printdest)
	printer = exec.Command("lp", "-d", tmpStr)
	stdin, printErr := printer.StdinPipe()
	if printErr != nil {
		usageTips("could not open pipe to \"" + tmpStr + "\"!")
		os.Exit(1)
	}

	// dealing with the page type
	var pageCnt int
	var lineFlag1 *bufio.Scanner
	fout := os.Stdout

	if saAddr.infilename != "" {
		inputStream, printErr := os.Open(saAddr.infilename)
		if printErr != nil {
			usageTips("could not open pipe to \"" + tmpStr + "\"!")
			os.Exit(1)
		}
		lineFlag1 = bufio.NewScanner(inputStream)
	} else {
		lineFlag1 = bufio.NewScanner(os.Stdin)
	}
	var flag bool
	for lineFlag1.Scan() {
		flag = false
		for _, a := range lineFlag1.Text() {
			if a == '\f' {
				if pageCnt >= saAddr.start && pageCnt <= saAddr.end {
					flag = true
					fout.Write([]byte("\n"))
					stdin.Write([]byte("\n"))
				}
				pageCnt++
			} else {
				if pageCnt >= saAddr.start && pageCnt <= saAddr.end {
					fout.Write([]byte(string(a)))
					stdin.Write([]byte(string(a)))
				}
			}
		}
		if pageCnt >= saAddr.start && pageCnt <= saAddr.end && flag != true {
			fout.Write([]byte("\n"))
			stdin.Write([]byte("\n"))
		}
		flag = false
	}

	if pageCnt < saAddr.start {
		fmt.Fprintf(os.Stderr, "WARNING:start_page (%d) greater than total pages (%d)\n", saAddr.start, pageCnt)
	} else {
		if pageCnt < saAddr.end {
			fmt.Fprintf(os.Stderr, "WARING:end_page (%d) greater than total pages (%d)\n", saAddr.end, pageCnt)
		}
	}
	if saAddr.printdest != "" {
		stdin.Close()
		printer.Stdout = os.Stdout
		printer.Start()
	}
}
