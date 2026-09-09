package goroutine

import (
	"fmt"
	"sync"
	"time"
)

//minitask 9 DailyRoutine Go Routine dan Waitgroups wg.GO

func DailyWorker() {
	var wg sync.WaitGroup
	fmt.Println("Bangun Pagi")
	wg.Add(1)
	go Mandi(&wg)
	wg.Add(1)
	go BuatKopi(&wg)
	wg.Add(1)
	go Sarapan(&wg)
	wg.Add(1)
	go MerapikanKamar(&wg)

	wg.Wait()
	// time.Sleep(1 * time.Second)
	fmt.Println("Berangkat Kerja")
}

func Mandi(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Mandi Pagi")
}

func MerapikanKamar(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Merapikan Kamar")
}

func Sarapan(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Sarapan Pagi")
}

func BuatKopi(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Buat Kopi & Minum Kopi")
}
