package main

import "fmt"


func main() {
	test1()
}

type Player interface {
	play(time int, name string) string
}

func playWithPlayer(p Player, time int, name string) string{
	s:= p.play(time, name)
	return s
}

type PlayFunc func(int, string) string

func (f PlayFunc) play(time int, name string) string {
	return f(time, name)
}

func test1() {
	fmt.Println(playWithPlayer(PlayFunc(func(time int, name string) string {
		return fmt.Sprintf("%s %d\n", name, time * 2)
	}), 3, "ali"))
}

