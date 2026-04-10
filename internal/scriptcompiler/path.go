package scriptcompiler

import "fmt"

func Addpath() {
	fmt.Println("export PATH=\"$PWD/.samengine/bin:$PATH")
	fmt.Println("$env:PATH = \"$PWD\\.samengine\\bin;\" + $env:PATH")
	fmt.Println("set PATH=%CD%\\.samengine\\bin;%PATH%")
}
