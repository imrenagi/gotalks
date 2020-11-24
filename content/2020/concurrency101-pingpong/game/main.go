package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

//STARTMAIN, OMIT
func main() {

	r := referree{}
	p1 := newPlayer("ping")
	p2 := newPlayer("pong")

	g := newGame(p1, p2) // HL3
	r.Start(g)           //start on main goroutine // HL3
}

//STOPMAIN, OMIT

type ball struct {
	hits       int
	lastPlayer string
}

//STARTNEWGAME, OMIT
func newGame(player ...*player) *game {
	g := &game{ //...
		Score:   make(map[string]int), //OMIT
		Players: player,               //OMIT
		table:   make(chan *ball),     //OMIT
		done:    make(chan struct{}),  //OMIT
	}

	go g.loop() // HL3
	return g
}

//STOPNEWGAME, OMIT

//STARTGAMESTRUCT, OMIT
type game struct {
	sync.Mutex // HL3
	Score      map[string]int
	Players    []*player
	table      chan *ball
	done       chan struct{}
}

// Table returns the channel used to
// distribute the ball
func (g *game) Table() chan *ball { // HL3
	return g.table
}

//STOPGAMESTRUCT, OMIT

//STARTGAMELOOP, OMIT
func (g *game) loop() {
	for {
		select {
		case <-g.done: // HL3
			for _, p := range g.Players {
				p.Done()
			}
			return
		}
	}
}

//STOPGAMELOOP, OMIT

//STARTGAMESTART, OMIT
// Start tells all player to start playing
// internally it spawn goroutine for each players
func (g *game) Start(done chan *ball) {
	for _, p := range g.Players {
		go p.Play(g.Table(), done) // HL3
	}
}

//STOPGAMESTART, OMIT

//STARTGAMEINCSCORE, OMIT
// IncrementScore increment the score of a player.
// Once the score equal to max score, then the player wins
func (g *game) IncrementScore(p string) bool {
	g.Lock()         // HL3
	defer g.Unlock() // HL3

	log.Println("increment score for", p) // OMIT
	if _, ok := g.Score[p]; ok {          // <- potential race // HL3
		g.Score[p]++
	} else {
		g.Score[p] = 1
	}

	if g.Score[p] == 3 {
		return true
	}

	return false
}

//STOPGAMEINCSCORE, OMIT

//STARTPLAYER, OMIT
func newPlayer(name string) *player {
	return &player{
		Name: name,
		done: make(chan struct{}),
	}
}

type player struct {
	Name string
	done chan struct{} // HL3
}

//STOPPLAYER, OMIT

//STARTPLAYERDONE, OMIT
func (p player) Done() {
	p.done <- struct{}{} // HL3
}

//STOPPLAYERDONE, OMIT

//STARTPLAYERPLAY, OMIT
func (p player) Play(table chan *ball, done chan *ball) { // HL3
	for {
		s := rand.NewSource(time.Now().UnixNano()) // OMIT
		r := rand.New(s)                           // OMIT
		//.....
		select {
		case ball := <-table:
			v := r.Intn(1001) // OMIT
			if v%11 == 0 {
				log.Println(p.Name, "drop the ball") // OMIT
				done <- ball
				continue //continue instead of return // HL3
			}

			ball.hits++
			ball.lastPlayer = p.Name
			time.Sleep(50 * time.Millisecond)
			log.Println(p.Name, "hits ball", ball.hits) // OMIT
			table <- ball
		case <-p.done: // receive from done channel // HL3
			log.Println(p.Name, "done with the game") // OMIT
			return                                    // HL3
		}
	}
}

//STOPPLAYERPLAY, OMIT

//STARTREFERREESTART, OMIT
type referree struct{} // HL3

func (r *referree) Start(g *game) {
	done := make(chan *ball) // HL3
	g.Start(done)

	table := g.Table()
	table <- new(ball) // HL3

	for {
		select {
		case b := <-done:
			log.Println("ball is returned to referree. point for ", b.lastPlayer) // OMIT
			if g.IncrementScore(b.lastPlayer) {
				log.Println("game is over. winner is", b.lastPlayer) // OMIT
				g.done <- struct{}{}                                 // HL3
				return                                               // HL3
			}
			table <- new(ball) // start new round // HL3
		}
	}
}

//STOPREFERREESTART, OMIT
