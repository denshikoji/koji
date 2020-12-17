package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {

	// clear the screen
	cls()

	// check which os
	osCheck()

	// say hello
	fmt.Println("hello koji :)")

	// start lifecycle
	go liveLife()

	// wait for user input
	inputHandler()
}

func cls() {

	// windows clears the screen differently
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux clears the screen with special chars
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	}

}

func handle(msg string, err error) {
	if err != nil {
		fmt.Printf(brightred+"\n%s: %s"+white, msg, err)
	}
}

var (

	// baby kojis have 10 food units
	kojiFood = 10

	// baby koji should only eat 10 kibbles
	kojiMaxFood = 10

	// baby koji is simple to please
	kojiHappiness = kojiFood

	// koji is born
	kojiAge = 0
)

func liveLife() {

	// live forever
	for {

		// clear the screen
		cls()

		// are we still alive?
		isKojiAlive()

		// draw the stats
		kojiStat()

		// delay before next cycle
		delayMin(1)

		// consume 1 kibble
		kojiEat()

		// add 1 min to age
		kojiAge++
	}
}

func kojiAchievement(kojiAge int) {

	// unlock name
	if kojiAge == 5 {
		fmt.Println(kojiNeedsName)
	}
}

func isKojiAlive() {

	// did koji starve?
	if kojiFood <= 0 {
		cls()
		kojiAvatar()
		fmt.Println(brightred + kojiHasDiedMsg + nc)
		os.Exit(0)
	}

	// is koji overfed?
	if kojiFood >= kojiMaxFood+5 {
		cls()
		kojiAvatar()
		fmt.Println(brightred + kojiHasDiedMsg + nc)
		os.Exit(0)
	}
}

func delayMin(timeMinutes time.Duration) {
	time.Sleep(timeMinutes * time.Minute)
}

func delaySec(timeSeconds time.Duration) {
	time.Sleep(timeSeconds * time.Second)
}

func kojiEat() {

	// is koji full?
	kojiFood = kojiFood - 1

}

func kojiStat() {

	// clear the screen
	cls()

	// draw koji face
	kojiAvatar()

	// is koji hungry?
	foodStatMsg()

	// is koji ready to achieve?
	kojiAchievement(kojiAge)
}

func kojiAvatar() {

	// draw the name in blue if happy
	if kojiHappiness >= 5 {
		fmt.Printf(brightcyan + kojiName + nc)
	} else if kojiHappiness < 5 && kojiHappiness > 2 {
		fmt.Printf(brightyellow + kojiName + nc)
	} else if kojiHappiness <= 2 {
		fmt.Printf(brightred + kojiName + nc)
	}

	// determine koji sentiment
	if kojiHappiness >= 5 {
		kojiHappyFace()
	} else if kojiHappiness < 5 && kojiHappiness >= 1 {
		kojiSadFace()
	} else if kojiHappiness <= 0 {
		kojiDeadFace()
	}
}

func kojiHappyFace() {

	// pick a happy face
	thisFace := roll(0, len(kojiHappyFaces))
	fmt.Println(kojiHappyFaces[thisFace])

}

func kojiSadFace() {

	// pick a sad face
	thisFace := roll(0, len(kojiSadFaces))
	fmt.Println(kojiSadFaces[thisFace])

}

func kojiDeadFace() {

	// pick a dead face
	thisFace := roll(0, len(kojiDeadFaces))
	fmt.Println(kojiDeadFaces[thisFace])
}

func roll(min, max int) int {
	return rand.Intn(max-min) + min
}

func foodStatMsg() {

	// is koji well fed?
	if kojiFood < 2 {
		kojiHappiness = 2
		fmt.Println(kojiStarvingMsg)
	} else if kojiFood > 2 && kojiFood < 8 {
		kojiHappiness = 6
		fmt.Println(kojiWantsFoodMsg)
	} else if kojiFood == kojiMaxFood {
		kojiHappiness = 10
		fmt.Println(kojiIsHappyMsg)
	} else if kojiFood > kojiMaxFood+1 {
		kojiHappiness = kojiFood / 3
		fmt.Println(kojiIsOverfed)
	} else if kojiFood >= kojiMaxFood+2 {
		kojiHappiness = 2
		fmt.Println(kojiIsSick)
		isKojiAlive()
	}
}

func rawStats() {

	// clear the screen
	cls()

	// draw quick stats
	fmt.Printf("Age: %v mins\n", kojiAge)
	fmt.Printf("Food: %v\n", kojiFood)
	fmt.Printf("Happiness: %v", kojiHappiness)

	// give time to read
	delaySec(1)

	// redraw the screen
	kojiStat()
}

func inputHandler() {

	// wait for user to give commands
	reader := bufio.NewReader(os.Stdin)

	// wait in a loop
	for {

		// format the input
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)

		if strings.Compare("help", text) == 0 {
			menu()
		} else if strings.Compare("exit", text) == 0 {
			os.Exit(0)
		} else if strings.Compare("", text) == 0 {
			kojiStat()
		} else if strings.Compare("s", text) == 0 {
			rawStats()
		} else if strings.Compare("quit", text) == 0 {
			os.Exit(0)
		} else if strings.Compare("close", text) == 0 {
			os.Exit(0)
		} else if strings.Compare("q", text) == 0 {
			os.Exit(0)
		} else if strings.Compare("a", text) == 0 {
			kojiAge++
			kojiStat()
		} else if strings.Compare("n", text) == 0 && kojiAge >= 5 {
			fmt.Printf("name: ")
			var name string
			fmt.Scanf("%s", &name)
			kojiName = name
			kojiStat()
		} else if strings.Compare("feed", text) == 0 {
			feedKoji()
		} else if strings.Compare("f", text) == 0 {
			feedKoji()
		}
	}
}

func feedKoji() {

	// is koji being overfed?
	if kojiFood > kojiMaxFood {
		kojiFood++
		isKojiAlive()
	}

	// is koji full?
	if kojiFood == kojiMaxFood {
		kojiFood++
		cls()
		kojiAvatar()
		fmt.Println(brightgreen + kojiIsFullMsg)
		cls()
		kojiAvatar()
		fmt.Println(nc + kojiIsFullMsg)

		// is koji hungry?
	} else if kojiFood < kojiMaxFood {
		kojiFood++
		cls()
		kojiAvatar()
		fmt.Println(kojiEatingMsg)
	}
}

func menu() {
	fmt.Println("FEED - give 1 food to koji.")
}
