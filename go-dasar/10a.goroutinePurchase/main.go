package main

import (
	"fmt"
	"sync"
	"time"
)

// Product merepresentasikan produk dengan stok dan kunci Mutex
type Product struct {
	mu    sync.Mutex
	stock int
}

// Purchase mencoba mengurangi stok jika tersedia
func (p *Product) Purchase(userID int, limiter chan struct{}, wg *sync.WaitGroup) {
	// 4. DONE: Pastikan WaitGroup selesai setelah goroutine ini exit
	defer wg.Done()

	// 1. Ambil Token dari Limiter (BLOKIR jika Channel penuh)
	// Ini memastikan hanya sejumlah goroutine yang diizinkan masuk ke logika pembelian.
	limiter <- struct{}{}

	// Pastikan token dikembalikan ke channel saat goroutine keluar dari fungsi
	defer func() { <-limiter }()

	// 2. LOCK: Kunci Mutex untuk Critical Section (mengakses dan mengubah 'stock')
	p.mu.Lock()
	defer p.mu.Unlock()

	// Simulasi delay Database check
	time.Sleep(time.Millisecond * 2)

	// Logic Pembelian
	if p.stock > 0 {
		p.stock--
		fmt.Printf("[SUKSES] User %d membeli. Sisa stok: %d\n", userID, p.stock)
	} else {
		fmt.Printf("[GAGAL ] User %d, Produk habis.\n", userID)
	}
}

func main() {
	const initialStock = 5
	const totalBuyers = 500
	const maxConcurrentLimit = 10 // Maksimal 10 goroutine memproses pembelian secara bersamaan

	product := Product{stock: initialStock}
	var wg sync.WaitGroup

	// Channel sebagai Request Limiter (Buffered Channel)
	// Kapasitas 10 berarti hanya 10 goroutine yang dapat menempati channel ini.
	limiter := make(chan struct{}, maxConcurrentLimit)

	startTime := time.Now()
	fmt.Printf("Flash Sale Dimulai! Stok Awal: %d, Total Pembeli: %d\n", initialStock, totalBuyers)
	fmt.Printf("Maksimal Konkurensi: %d\n", maxConcurrentLimit)
	fmt.Println("----------------------------------------------------------------")

	// Luncurkan Goroutine
	for i := 1; i <= totalBuyers; i++ {
		wg.Add(1)
		go product.Purchase(i, limiter, &wg)
	}

	// Wait: Tunggu semua goroutine pembeli selesai
	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("Flash Sale Selesai dalam %s\n", elapsed)
	fmt.Printf("Stok Akhir Produk: %d\n", product.stock)
	fmt.Printf("Total Pembeli Sukses: %d\n", initialStock-product.stock)
}
