package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intellince`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
		`fleming_ian`:      {`steaks`, `cigars`, `espionage`},
	}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
