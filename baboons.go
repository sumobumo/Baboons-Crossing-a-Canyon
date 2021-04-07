package main 

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
// declaring total number of baboons which will be crossing the canyon
const (
	MAX_BABOONS=30
)

var (
	Mutex=new(sync.Mutex)
	eastWG=new(sync.WaitGroup)
	westWG = new(sync.WaitGroup)
	eastdoor = true
	westdoor = true
	listOfbaboons = make([]int,0,5)		//slice to hold the ids of baboons on the rope
	
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())	//to randomly generate numbers
	destiny :=""	//baboon's destiny
	wg:=new(sync.WaitGroup)	
	for id:=0;id<MAX_BABOONS;id++ {
		//sleeping so that each baboons arrives on the canyon randomly between 1 to 8 seconds
		time.Sleep(time.Duration(rand.Intn(8)+1)*time.Second)
		direction := rand.Intn(2)	//0 for eastward moving and 1 for westward moving
		if (direction == 0){
			destiny="east"
		}else {
			destiny ="west"
		}
		wg.Add(1)	// add 1 whenever a boboon gets on the rope
		go baboons(id,destiny,wg)
	}
	wg.Wait()	// waiting till all the baboons has finished crossing the canyon
	fmt.Println("All Baboons has finished crossing the canyon")
}

func baboons(id int, destiny string,wg *sync.WaitGroup){
	fmt.Println("Baboon",id,"arrived at the canyon and wants to go ",destiny)
	
	switch destiny {
		case "east" : 
			//checking if west door is open for west baboon to go east
			//also checks the east door is locked so that no baboons from east gets on the rope
			if eastdoor!=false && westdoor!=true {
				Mutex.Lock()
				eastdoor = false	//closes east door
				Mutex.Unlock()
				fmt.Println("Closing east door. Baboons from east cannot get on rope")
			}	
			fmt.Println("Baboons ", listOfbaboons, "are on the rope")
			westWG.Wait()	//waits till baboons on rope coming from east gets off the rope
			Mutex.Lock()
			eastdoor=false
			westdoor =true		//wants to go east so door on west must be open to get on rope
	
			Mutex.Unlock()
			eastWG.Add(1)	//record that a baboon is going east
			moveBaboon(id,destiny)
			eastWG.Done()	
			
		case "west" : 
			//checking if east door is open for east baboon to go west
			//also checks the west door is locked so that no baboons from west gets on the rope
			if (eastdoor!=true && westdoor!=false){
				Mutex.Lock()
				westdoor=false	//closes west door
				Mutex.Unlock()
				fmt.Println("Closing west door. Baboons from west cannot get on rope")
			} 	
			fmt.Println("Baboons ", listOfbaboons, "are on the rope")
			eastWG.Wait()	// waits till baboons on rope coming from west gets off the rope
			Mutex.Lock()
			eastdoor=true	//wants to go west so door on east must be open to get on rope
			westdoor =false
			
			Mutex.Unlock()
			westWG.Add(1)	// record that a baboon is going west 
			moveBaboon(id,destiny)
			westWG.Done()
			
		default:
	}
	wg.Done()	// remove 1 from waitgroup whenever a baboon gets off the rope
}

func moveBaboon(id int, destiny string){
	
	Mutex.Lock()
	listOfbaboons = append(listOfbaboons,id)	// baboons gets on the rope add baboon id to list
	Mutex.Unlock()
	
	//inter-baboon spacing is 1 second
	time.Sleep(time.Duration(1)*time.Second)
	fmt.Println("Baboon",id,"is on the rope and is moving ",destiny) 
	//Each traversal takes exactly 4 seconds, after the baboon is on the rope.
	time.Sleep(time.Duration(4)*time.Second)
	fmt.Println("Baboon",id,"has finished crossing canyon. It is on west. It got off the rope")
	
	Mutex.Lock()
	
	listOfbaboons = listOfbaboons[1:] //baboon gets off the rope so delete the id of the baboon which is first on the list
	Mutex.Unlock()
}

