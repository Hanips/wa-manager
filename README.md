# WhatsMeong (Free-Tier Architecture)

**WhatsMeong** is a standalone REST API Gateway and Web Dashboard for WhatsApp, built on top of the open-source [**whatsmeow**](https://github.com/tulir/whatsmeow) library. 

This project has been heavily modified and optimized to run **100% reliably on Free-Tier Cloud Infrastructure** without experiencing memory leaks, crashes, or data corruption.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Render](https://img.shields.io/badge/Render-46E3B7?style=for-the-badge&logo=render&logoColor=white)

---

## 🏗️ 100% Free Architecture

This system is designed with a stateless API approach and crash-proof concurrency to survive the auto-sleep cycles and minimal resources typical of free hosting providers.

The core architecture consists of:

1. ☁️ **Render.com (Web Service - Free Tier)**
   Acts as the primary host for the Golang application. Render automatically spins down (sleeps) the application if there is no traffic for 15 minutes. WhatsMeong is equipped with a self-healing reconnection mechanism to handle these restarts gracefully without losing your WhatsApp session.
   
2. 🐘 **Neon.tech (PostgreSQL - Serverless Free Tier)**
   Because Render's free tier uses an ephemeral filesystem, the default SQLite store used by `whatsmeow` for WhatsApp sessions has been completely migrated to an external PostgreSQL database via Neon.tech. This ensures session persistence across container restarts.
   
3. ⏰ **Cron-job.org (Ping Service - Free)**
   To prevent Render from putting the container to sleep, you can use a free service like Cron-job.org to ping the `/ping` endpoint of this application every 10-14 minutes.

---

## ✨ Features

- **Modern Web Dashboard:** A clean, Cupertino-inspired (Liquid Glass) UI to link devices via QR code, manage Webhooks, and test message delivery. Built natively without heavy frontend frameworks.
- **Comprehensive REST API:** Send text, base64 images, remote documents, vCards, locations, and interactive Polls using standard JSON payloads.
- **Safe Broadcast Mechanism:** A background broadcasting feature with randomized delays between messages to prevent algorithmic bans from Meta.
- **Multi-Webhook Support:** Automatically captures incoming messages and Poll responses, forwarding them to multiple endpoints (e.g., Make.com, n8n) in real-time.
- **Crash-Proof Concurrency:** Protected by Go Mutex locks and frontend request guards. The system will not crash or hang due to UI button spam or sudden network disconnects.
- **Drop-in Replacement Ready:** Easily replace expensive paid API gateways with this self-hosted solution. 

---

## 🚀 Quick Deployment Guide

1. Create a **Neon.tech** account and get your `DATABASE_URL`.
2. Deploy this repository to **Render.com** (Select "Web Service", Language: "Go").
3. Add the following Environment Variables in Render:
   - `DATABASE_URL` = `postgres://...`
   - `API_KEY` = `YOUR_SECRET_PASSWORD` (This will act as your UI password and API Bearer Token)
4. Open your application URL (e.g., `https://wa-manager-xyz.onrender.com`), login with your `API_KEY`, and scan the QR Code using the WhatsApp app on your phone.
5. (Optional) Register your `.../ping` URL to **Cron-job.org** for 24/7 uptime.

---
*Built with Go. The core WhatsApp Web engine rights belong to the whatsmeow developers.*
