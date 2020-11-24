package main

import (
	"log"
	"math/rand"
	"time"
)

//STARTPLAYER, OMIT
func player(name string, table chan *Ball, done chan *Ball) { //accept done channel // HL1
	for {
		s := rand.NewSource(time.Now().UnixNano()) // OMIT
		r := rand.New(s)                           // OMIT

		select {
		case ball := <-table:
			v := r.Intn(1001) // OMIT
			if v%11 == 0 {    //random condition to drop ball // HL1
				log.Println(name, "drop the ball")
				done <- ball
				return
			}

			ball.hits++
			ball.lastPlayer = name //assign the player name // HL1
			log.Println(name, "hit the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return //terminate goroutine to avoid leak // HL1
		}
	}
}

//STOPPLAYER, OMIT

//STARTBALL, OMIT
type Ball struct {
	hits       int
	lastPlayer string
}

//STOPBALL, OMIT

//STARTREFERREE, OMIT
func referree(table chan *Ball, done chan *Ball) { // HL2
	table <- new(Ball)

	for {
		select {
		case ball := <-done:
			log.Println("game over. winner:", ball.lastPlayer)
			return
		}
	}
}

//STOPREFERREE, OMIT

//STARTMAIN, OMIT
func main() {

	table := make(chan *Ball)
	done := make(chan *Ball)

	go player("ping", table, done)
	go player("pong", table, done)

	referree(table, done)
}

//STOPMAIN, OMIT
