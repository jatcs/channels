package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	
	//goroutine to send a message to the channel
	//after a certain amount of time
	go func() {
		time.Sleep(1 * time.Second)
		//send message "one" into c1
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		//awaits the values simultaneously and prints them 
		//them as they arrive

		//select proceeds with the first receive thats ready
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}


	fmt.Println()

	//timeouts
	//useful for function calls / routines that take a certain amt of time
	//and you wanna see if it fit into a diff (smaller/larger) amount of time
	c3 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "result"
	}()
	//if you are in a (one time / non loop) select
	//the call that takes longer won't get outputted
	select {
	case res := <- c3:
		fmt.Println(res, "got there first")
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
	
	fmt.Println()

	rBoard := make([][]int, 3)
	rBoard = mkBoard()
	//fmt.Println(rBoard)
	//prints [[0 0 0] [0 0 0] [0 0 0]]

	//display(rBoard)
	//prints
	//[0 0 0]
	//[0 0 0]
	//[0 0 0]

	//hope to use channels to keep the display on
	//as it recieves input from the players
	

	var r int
	var c int
	fmt.Println("Enter the row and column you'd like to change")
	fmt.Println("Note: slice indexes start at 0, so row '1' is 0...")
	
	//playerOne(rBoard[:], r, c)
	//display(rBoard)

	p1 := make(chan int, 1)
	p2 := make(chan int, 1)

	go func() {
		fmt.Println("Player One: ")
		fmt.Scanln(&r, &c)
		playerOne(rBoard[:], r, c)
		//prob don't need to bother sending user inputs to the channel
		//p1 <- r
		//p1 <- c
		p1 <- 1
	}()

	rec1 := <- p1

	go func() {
		fmt.Println("Player Two: ")
		fmt.Scanln(&r, &c)
		playerTwo(rBoard[:], r, c)
		//p2 <- r
		//p2 <- c
		p2 <- 2
	}()
	rec2 := <- p2
	fmt.Print(rec1, rec2, "done \n")
	display(rBoard)

	//Current problems:
	//can keep the display up for multiple turns?

	//things to add:
	//maybe a for loop with 9 iterations (fill each 0)
	//alternate p1 and p2
}

func mkBoard() [][]int{
	board := make([][]int, 3)
	for i:=0; i<3; i++ {
		board[i] = make([]int, 3)
		for j:=0; j<3; j++ {
			board[i][j] = 0
		}
	}
	return board
}

func display(b [][]int) {
	//print each row of the board on separate lines
	for i:=0; i<3; i++ {
		fmt.Println(b[i])
	}
}

func playerOne(b [][]int, row int, col int) {
	//player changes a certain 0 to 1 to mark their choice
	//for that turn
	
	//for some reason didn't need to make it a pointer to the slice
	//it still changed the rBoard from main
	b[row][col] = 1
}

func playerTwo(b [][]int, row int, col int) {
	b[row][col] = 2
}
