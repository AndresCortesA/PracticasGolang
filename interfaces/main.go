package main

import (
	"fmt"
	"interfaces/users"
)

func main() {
	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee("Robert")
	//	frank.Active
	total := frank.CalcTotal(90, 65, 9.6)

	fmt.Println(total)

	robert.DesactivateEmployee(frank)

	fmt.Println(frank.CalcTotal(90, 65, 92.6))
}
