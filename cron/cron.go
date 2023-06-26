package main

import (
	"unicode/utf8"
)

func main() {

	//fmt.Println("Please visit http://127.0.0.1:12345/")
	//http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	s := fmt.Sprintf(". -- Time: %s", time.Now().String())
	//	fmt.Fprintf(w, "%v\n", s)
	//	log.Printf("%v\n", s)
	//})
	//if err := http.ListenAndServe(":12345", nil); err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	//c()

	//a := str2runes("fafa")
	//fmt.Println(a)

	cc()
}

func cc() {

}

func str2runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
