package main

import (
	"fmt"
	"math/rand"
) 

func main() {
	fmt.Println("Welcome to Pig Dice Game")

	total_score:=0
	turn:=1
	turn_score:=0
 
    for(total_score<=20 && turn_score<=20){

    fmt.Println("press h to hold and r to roll")
        var choice string
        fmt.Scan(&choice)
        switch choice {
        case "r":
            roll := rand.Intn(6) + 1
            fmt.Println("your score is ",roll)
            if(roll==1){
                turn_score =0
                fmt.Println(turn_score)
                turn++
            }else{
                turn_score+=roll
                fmt.Println(turn_score)
            }
        case "h":
            turn++
            total_score+=turn_score
            fmt.Println("your total score is ",total_score)
            if(total_score>=20){
                break
            }
        }
    }

    fmt.Println("Congrats you won in",turn, "turns")
	fmt.Println("Thank you for playing the game")

}