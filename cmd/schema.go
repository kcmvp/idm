package cmd

import "log"
import "entgo.io/contrib/schemast"

func main() {
	_, err := schemast.Load("./ent/schema")
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	//// A no-op since we did not manipulate the Context at all.
	//if err := schemast.Print("./ent/schema"); err != nil {
	//	log.Fatalf("failed: %v", err)
	//}
}
