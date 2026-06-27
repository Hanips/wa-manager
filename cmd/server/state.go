package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var client *whatsmeow.Client
var qrCode string
var qrCodeMu sync.RWMutex
var qrFlowMu sync.Mutex // Mencegah dua startQRFlow() berjalan bersamaan
var appDB *sql.DB
var waContainer *sqlstore.Container

func startQRFlow() {
	// Guard: pastikan hanya 1 QR flow berjalan sekaligus
	if !qrFlowMu.TryLock() {
		fmt.Println("startQRFlow: already running, skipping duplicate call.")
		return
	}

	if client != nil {
		client.Disconnect()
	}

	// Buat ulang Client dan DeviceStore untuk memastikan state benar-benar bersih setelah logout
	deviceStore, err := waContainer.GetFirstDevice(context.Background())
	if err != nil {
		fmt.Printf("Failed to get device: %v\n", err)
		qrFlowMu.Unlock()
		return
	}
	clientLog := waLog.Stdout("Client", "INFO", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	qrChan, _ := client.GetQRChannel(context.Background())
	err = client.Connect()
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		qrFlowMu.Unlock() // Release lock jika gagal
		
		// Self-healing: Retry setelah 5 detik agar tidak mati selamanya jika ada network hiccup
		go func() {
			time.Sleep(5 * time.Second)
			startQRFlow()
		}()
		return
	}
	
	// Jalankan di background
	go func() {
		// Pastikan lock selalu dilepas saat goroutine ini selesai
		defer qrFlowMu.Unlock() 
		
		for evt := range qrChan {
			if evt.Event == "code" {
				qrCodeMu.Lock()
				qrCode = evt.Code
				qrCodeMu.Unlock()
				fmt.Println("QR code available. Visit /qr to see it.")
			} else if evt.Event == "timeout" {
				fmt.Println("QR code expired (timeout). Getting a new one...")
				qrCodeMu.Lock()
				qrCode = ""
				qrCodeMu.Unlock()
				
				// Mulai ulang flow baru setelah jeda
				go func() {
					time.Sleep(2 * time.Second)
					startQRFlow()
				}()
				return // Keluar dari loop, trigger defer Unlock()
			} else if evt.Event == "success" {
				fmt.Println("Login successful!")
				qrCodeMu.Lock()
				qrCode = ""
				qrCodeMu.Unlock()
				return // Keluar dari loop, trigger defer Unlock()
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	}()
}
