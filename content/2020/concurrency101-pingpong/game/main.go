package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {

	r := referree{}
	p1 := newPlayer("ping")
	p2 := newPlayer("pong")

	g := newGame()
	g.referree = &r
	g.AddPlayer(p1)
	g.AddPlayer(p2)

	r.Start(g)
}

type Ball struct {
	hits       int
	lastPlayer string
}

func newGame() *game {
	return &game{
		Score:   make(map[string]int),
		Players: make([]*Player, 0),
		table:   make(chan *Ball),
		done:    make(chan struct{}),
	}
}

type game struct {
	sync.Mutex
	Score    map[string]int
	Players  []*Player
	referree *referree
	table    chan *Ball
	done     chan struct{}
}

func (g *game) AddPlayer(p *Player) {
	g.Players = append(g.Players, p)
}

func (g game) Table() chan *Ball {
	return g.table
}

func (g game) Start(done chan *Ball) {

	for _, p := range g.Players {
		go p.Play(g.Table(), done)
	}

	for {
		select {
		case <-g.done:
			for _, p := range g.Players {
				p.Done()
			}
			return
		}
	}
}

func (g *game) Inc(p string) {
	g.Lock()
	defer g.Unlock()

	fmt.Println("increment score for", p)

	if _, ok := g.Score[p]; ok {
		g.Score[p]++
		return
	}
	g.Score[p] = 1
}

func (g game) Finish() bool {
	g.Lock()
	defer g.Unlock()

	for k, v := range g.Score {
		fmt.Println(k, "score:", v)
		if v == 3 {
			return true
		}
	}
	return false
}

func newPlayer(name string) *Player {
	return &Player{
		Name: name,
		done: make(chan struct{}),
	}
}

type Player struct {
	Name string
	done chan struct{}
}

func (p Player) Done() {
	p.done <- struct{}{}
}

func (p Player) Play(table chan *Ball, done chan *Ball) {
	for {
		s := rand.NewSource(time.Now().UnixNano()) // OMIT
		r := rand.New(s)                           // OMIT

		select {
		case ball := <-table:
			v := r.Intn(1001) // OMIT
			if v%11 == 0 {
				log.Println(p.Name, "drop the ball")
				done <- ball
				continue
			}

			ball.hits++
			ball.lastPlayer = p.Name
			time.Sleep(50 * time.Millisecond)

			log.Println(p.Name, "hits ball", ball.hits)

			table <- ball
		case <-p.done:
			fmt.Println(p.Name, "done with the game")
			return
		}
	}
}

type referree struct {
}

func (r *referree) Start(g *game) {

	table := g.Table()

	done := make(chan *Ball)
	go g.Start(done)

	table <- new(Ball)

	for {
		select {
		case b := <-done:
			log.Println("ball is returned to referree. point for ", b.lastPlayer)

			g.Inc(b.lastPlayer)

			if g.Finish() {
				log.Println("game is over")
				g.done <- struct{}{}
				return
			}

			// creating new ball and push it back to the channel
			// to restart the game
			table <- new(Ball)

		}
	}
}
