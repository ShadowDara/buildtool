package cli

import "fmt"

func banner() {
	text := 
` ________  _________   
|\   __  \|\___   ___\ 
\ \  \|\ /\|___ \  \_| 
 \ \   __  \   \ \  \  
  \ \  \|\  \   \ \  \ 
   \ \_______\   \ \__\
    \|_______|    \|__|
`

	fmt.Println(text)
}
