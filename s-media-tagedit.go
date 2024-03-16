/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

func main() {
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).
}
