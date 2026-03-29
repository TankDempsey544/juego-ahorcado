package main
import ("fmt"; "math/rand"; "time")

func main() {
	rand.Seed(time.Now().UnixNano())
	pals := []string{"gato", "perro", "casa", "pala"}
	p := pals[rand.Intn(len(pals))]
	var lets []string
	f := 0

	for {
		gan := true
		fmt.Print("\n")
		for _, c := range p {
			sta := false
			for _, l := range lets {
				if l == string(c) { sta = true }
			}
			if sta { fmt.Print(string(c) + " ") } else { 
				fmt.Print("_ ")
				gan = false 
			}
		}
		fmt.Printf("\nErrores: %d\n", f)

		if gan { fmt.Println("Ganaste!"); break }
		if f >= 6 { fmt.Println("Perdiste! Era:", p); break }

		fmt.Print("Letra: ")
		var n string
		fmt.Scanln(&n)
		
		ok := false
		for _, c := range p {
			if string(c) == n { ok = true }
		}
		lets = append(lets, n)
		if !ok { 
			f++ 
			if f == 1 { fmt.Println(" O ") }
			if f == 2 { fmt.Println(" O\n | ") }
			if f == 3 { fmt.Println(" O\n/| ") }
			if f == 4 { fmt.Println(" O\n/|\\") }
			if f == 5 { fmt.Println(" O\n/|\\\n/  ") }
			if f == 6 { fmt.Println(" O\n/|\\\n/ \\") }
		}
	}
}