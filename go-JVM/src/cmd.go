package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Cmd struct {
	helpFlag    bool // help信息
	versionFlag bool
	cpOption    string   // classpath
	XjreOption  string   // jre目录
	class       string   // 类名称
	args        []string // 参数数组
	XDebug      bool     // 是否输出调试信息
	verbose     bool
}

func parseCmd() *Cmd {

	cmd := &Cmd{}
	flag.Usage = printUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.XDebug, "XDebug", false, "print debug msg")
	flag.BoolVar(&cmd.verbose, "verbose", false, "verbose class flag")

	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}

func (recv *Cmd) ToString() string {

	msg := fmt.Sprintf("  helpFlag:%s \n", strconv.FormatBool(recv.helpFlag))
	msg += fmt.Sprintf("  versionFlag:%s \n", strconv.FormatBool(recv.versionFlag))
	msg += fmt.Sprintf("  cpOption:%s \n", recv.cpOption)
	msg += fmt.Sprintf("  XjreOption:%s \n", recv.XjreOption)
	msg += fmt.Sprintf("  class:%s \n", recv.class)
	msg += "  args: "
	msg += fmt.Sprint(recv.args)
	msg += "\n"
	msg += fmt.Sprintf("  class:%s \n", recv.class)
	msg += fmt.Sprintf("  XDebug:%s \n", strconv.FormatBool(recv.XDebug))

	return msg
}
