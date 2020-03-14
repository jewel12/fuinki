package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var noteSize = 8
var noEvent = -2
var chords = []string{
	"G Em7 Am7 C",
	"GM7 F#m Em7 D",
	"A C7 D7 C7",
	"D A E",
	"A C#m7 Bm7 Bm7onE",
	"FM7 Em7 Am7",
	"F A7 Dm G",
	"A C#7 F#m",
	"F G# C",
	"D E G D",
	"F Eb G C",
	"Am7 Am7 G",
	"C G Am E7 F",
	"C Bb Dm G",
	"Am AmM7 Am7 D7",
	"G Gaug C Cm",
	"D G#dim G D",
	"D F#m7 Am7 B7 Em",
	"Dm G E A",
	"Bm7 E7 Am7 D7 G7",
}

func finki(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "")
	q := r.URL.Query()

	url := q["u"][0]
	fmt.Printf("Playing %s\n", url)
	play(url)
}

func play(url string) {
	melody := genMelody(url)
	chord := genChord(url)
	var ms []string
	for _, m := range melody {
		ms = append(ms, strconv.Itoa(m))
	}
	marg := fmt.Sprintf("MELODY=\"[%s]\"", strings.Join(ms, ","))
	barg := fmt.Sprintf("CHORD=\"%s\"", chord)
	out, err := exec.Command("make", "play", marg, barg).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
	}
}

// Ref.
// - http://www.theoreticallycorrect.com/Helmholtz-Pitch-Numbering/
// - https://github.com/tensorflow/magenta/tree/master/magenta/models/improv_rnn#generate-a-melody-over-chords
func genMelody(url string) []int {
	bs := md5.Sum([]byte(url))
	melody := make([]int, noteSize)

	for i := 0; i < noteSize; i++ {
		s := int(bs[i]) - 128 // 2 回に 1 回は no-event になる
		if s < 0 {
			melody[i] = noEvent
		} else {
			melody[i] = s
		}
	}

	return melody
}

// 2 つ選んで曲を長くしている
func genChord(url string) string {
	bs := md5.Sum([]byte(url))
	c1 := chords[int(bs[0])%len(chords)]
	c2 := chords[int(bs[1])%len(chords)]
	return c1 + " " + c2
}

func main() {
	m := http.NewServeMux()
	s := http.Server{Addr: ":3001", Handler: m}
	m.HandleFunc("/fuinki", funiki)
	s.ListenAndServe()
}
