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
	printdest := flag.String("d", "", "Set the destination of print.")
	flag.Parse()

	//judge if the input is valid and give tips
	if saAddr.start == -1 || saAddr.end == -1 || saAddr.start > saAddr.end || saAddr.start < 1 || saAddr.end < 1 {
		usageTips("the start page or the end page is not set correctly!")
		return
	}

	if saAddr.pagelen != 72 && saAddr.pagetype == true {
		usageTips("")
		return
	}

	if len(flag.Args()) > 1 {
		usageTips("")
		return
	}

	if len(flag.Args()) == 1 {
		saAddr.infilename = flag.Args()[0]
	}

	if *printdest != "" {
		saAddr.printdest = *printdest
	}

	if saAddr.pagetype == false {
		type1(saAddr, saAddr.infilename != "", saAddr.printdest != "")
	} else {
		type2(saAddr, saAddr.infilename != "", saAddr.printdest != "")
	}
}

//type 1 data and operate it
func type1(saAddr selpgArgs, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	curpage := 1
	curlines := 0
	if file {
		filein, err := os.OpenFile(saAddr.infilename, os.O_RDWR, os.ModeType)
		defer filein.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(filein)
		for line.Scan() {
			if curpage >= saAddr.start && curpage <= saAddr.end {
				os.Stdout.Write([]byte(line.Text() + "\n"))
				stdin.Write([]byte(line.Text() + "\n"))
			}
			curlines++
			if curlines %= saAddr.pagelen; curlines == 0 {
				curpage++
			}
		}
	} else {
		tmps := bufio.NewScanner(os.Stdin)
		for tmps.Scan() {
			if curpage >= saAddr.start && curpage <= saAddr.end {
				os.Stdout.Write([]byte(tmps.Text() + "\n"))
				stdin.Write([]byte(tmps.Text() + "\n"))
			}
			curlines++
			if curlines %= saAddr.pagelen; curlines == 0 {
				curpage++
			}
		}
	}
	if curpage < saAddr.end {
		fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
	}
	if pipe {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}

}

func type2(saAddr selpgArgs, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	curpage := 1
	if file {
		filein, err := os.OpenFile(saAddr.infilename, os.O_RDWR, os.ModeType)
		defer filein.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(filein)
		for line.Scan() {
			flag := false
			for _, c := range line.Text() {
				if c == '\f' {
					if curpage >= saAddr.start && curpage <= saAddr.end {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					curpage++
				} else {
					if curpage >= saAddr.start && curpage <= saAddr.end {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && curpage >= saAddr.start && curpage <= saAddr.end {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	} else {
		tmps := bufio.NewScanner(os.Stdin)
		for tmps.Scan() {
			flag := false
			for _, c := range tmps.Text() {
				if c == '\f' {
					if curpage >= saAddr.start && curpage <= saAddr.end {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					curpage++
				} else {
					if curpage >= saAddr.start && curpage <= saAddr.end {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && curpage >= saAddr.start && curpage <= saAddr.end {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	}
	if curpage < saAddr.end {
		fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
	}
	if pipe {

		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}

func usageTips(err string) {
	fmt.Fprintf(os.Stderr, err+"\n"+
		`usage: [-s start page(>=1)] [e end page(>=s)] [-l length of page(default 72)] [-f type of file(default 1)] [-d dest] [filename specify input file]
`)
}
