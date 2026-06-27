# 📚 WhatsMeong API Documentation (Super Lengkap)

Berikut adalah dokumentasi lengkap penggunaan (*cURL*) untuk seluruh fitur yang ada di sistem, dari pengiriman teks perorangan biasa, gambar base64, dokumen, hingga fitur premium seperti *Broadcast* dan *Polling*.

---

## 1. 💬 Chat Biasa (Perorangan)
Mengirim pesan teks sederhana ke satu nomor tujuan.

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "6285779336660",
    "message": "Pesan otomatis dari Render Cloud!"
  }'
```

---

## 2. 👥 Kirim ke Grup (Group Message)
Sama seperti kirim pesan biasa (Anda bisa melampirkan teks, gambar, lokasi, dll), namun isi `phone` dengan format **ID Grup WhatsApp** yang berakhiran `@g.us` atau nomor grup dengan tanda strip `-`.

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "120363022123456789@g.us",
    "message": "Halo Anggota Grup! Ini adalah pesan dari bot."
  }'
```

---

## 3. 📄 Kirim Dokumen (PDF, Excel, dll)
Mendownload dokumen dari URL dan langsung mengirimkannya via WhatsApp.

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  --max-time 180 \
  -d '{
    "phone": "6285779336660",
    "message": "Halo Bos, ini Laporan Tahunan perusahaan kita ya.",
    "document_url": "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
    "file_name": "Laporan_Tahunan_2026.pdf"
  }'
```

---

## 4. 🖼️ Kirim Gambar (URL atau Base64)

**A. Melalui URL Gambar:**
```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  --max-time 180 \
  -d '{
    "phone": "6285779336660",
    "message": "Laporan Saham: Tes Gambar berhasil masuk!",
    "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Cow_female_black_white.jpg/1280px-Cow_female_black_white.jpg"
  }'
```

**B. Melalui Format Base64:**
Berguna jika gambar dibuat secara dinamis oleh sistem Anda.
```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  --max-time 180 \
  -d '{
    "phone": "6285779336660",
    "message": "Laporan Saham: Tes Gambar berhasil masuk!",
    "image_base64": "kode_base64_kamu_di_sini"
  }'
```

---

## 5. 📢 Fitur Broadcast (Kirim Masal dengan Delay Acak)
Kirim pesan sekaligus ke puluhan nomor (Array) dengan jeda (*delay*) **acak** agar 100% terhindar dari pemblokiran (banned). Aman dari *timeout* karena diproses di *background*!

> 💡 **Trik Anti-Banned:** Anda bisa menetapkan jeda yang bervariasi dengan menambahkan `delay_ms_max`. Sistem akan otomatis mengocok waktu pengiriman secara random di antara `delay_ms` sampai `delay_ms_max`.

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/broadcast" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "phones": [
      "6285779336660",
      "6281234567890",
      "6289876543210"
    ],
    "delay_ms": 1300,
    "delay_ms_max": 4300,
    "payload": {
      "message": "Pesan Broadcast! Waktu kirim di-random antara 1.3 detik sampai 4.3 detik untuk tiap nomor."
    }
  }'
```
*Catatan: Jika `delay_ms_max` tidak diisi (atau lebih kecil dari `delay_ms`), maka sistem akan memakai jeda yang fixed/sama rata sesuai `delay_ms`.*

---

## 6. 📍 Kirim Lokasi (Maps)
Mengirim pin titik lokasi akurat yang bisa ditekan dan membuka navigasi Google Maps/Waze.

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "6285779336660",
    "location_lat": -6.2088,
    "location_lng": 106.8456,
    "location_name": "Kantor Pusat Monas",
    "message": "Jakarta Pusat, DKI Jakarta"
  }'
```

---

## 7. 📇 Kirim Kartu Kontak (VCard)
Mengirim nomor WhatsApp Anda dalam bentuk Kartu Nama (Contact Card) agar penerima bisa dengan mudah menekan tombol "Simpan Kontak".

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "6285779336660",
    "contact_name": "Hanief CS Neuralens",
    "contact_vcard": "BEGIN:VCARD\nVERSION:3.0\nN:;Hanief CS Neuralens;;;\nFN:Hanief CS Neuralens\nTEL;type=CELL;waid=6285779336660:+62 857-7933-6660\nEND:VCARD"
  }'
```

---

## 8. 📊 Kirim Polling (Pemungutan Suara)
Kirim daftar pilihan (*Polls*) yang interaktif. Jika *Webhook* Anda diaktifkan, jawaban yang dipilih *user* akan dilemparkan kembali ke server (Make/n8n) Anda secara *real-time*!

```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/send" \
  -H "Authorization: Bearer hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "6285779336660",
    "poll_name": "Pilih Layanan Kami:",
    "poll_options": ["Predict Ticker", "Analisa Laporan", "Bantuan CS"],
    "poll_selectable": 1
  }'
```

---

## 9. 🔗 Pengelolaan Multi-Webhook via API
Endpoint untuk mengonfigurasi banyak URL Webhook dari luar (misal Make.com atau N8N).

**Melihat daftar Webhook aktif:**
```bash
curl -X GET "https://wa-manager-mkvg.onrender.com/api/webhook?key=hanief123321"
```

**Mendaftarkan/Update Webhook:**
```bash
curl -X POST "https://wa-manager-mkvg.onrender.com/api/webhook?key=hanief123321" \
  -H "Content-Type: application/json" \
  -d '{
    "webhook_urls": [
      "https://make.com/webhooks/...",
      "https://n8n.neuralens.com/webhook/..."
    ]
  }'
```

---

## 10. 💻 Implementasi di Node.js (Axios)
Contoh penggunaan API menggunakan pustaka `axios` di Node.js (misal untuk backend Express, NestJS, atau skrip cron).

```javascript
const axios = require('axios');

const API_URL = "https://wa-manager-mkvg.onrender.com/send";
const API_KEY = "hanief123321";

async function sendWhatsApp() {
  try {
    const response = await axios.post(API_URL, {
      phone: "6285779336660",
      message: "Pesan ini dikirim otomatis dari Node.js!"
    }, {
      headers: {
        'Authorization': `Bearer ${API_KEY}`,
        'Content-Type': 'application/json'
      }
    });

    console.log("Berhasil:", response.data);
  } catch (error) {
    console.error("Gagal mengirim:", error.response ? error.response.data : error.message);
  }
}

sendWhatsApp();
```

---

## 11. 🐍 Implementasi di Python (Requests)
Contoh penggunaan API menggunakan pustaka `requests` di Python (misal untuk Bot, AI Automation, atau Web Scraper).

```python
import requests

API_URL = "https://wa-manager-mkvg.onrender.com/send"
API_KEY = "hanief123321"

headers = {
    "Authorization": f"Bearer {API_KEY}",
    "Content-Type": "application/json"
}

payload = {
    "phone": "6285779336660",
    "message": "Halo dari skrip Python! 🚀"
}

try:
    response = requests.post(API_URL, json=payload, headers=headers)
    response.raise_for_status() # Raise exception jika error
    print("Berhasil:", response.json())
except requests.exceptions.RequestException as e:
    print("Error:", e)
```

---

## 12. 🐘 Implementasi di PHP (cURL)
Contoh penggunaan API menggunakan cURL murni di PHP (misal untuk web Laravel, CodeIgniter, atau integrasi form kontak).

```php
<?php

$apiUrl = 'https://wa-manager-mkvg.onrender.com/send';
$apiKey = 'hanief123321';

$data = [
    'phone' => '6285779336660',
    'message' => 'Notifikasi dari website PHP Anda!'
];

$ch = curl_init($apiUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    'Authorization: Bearer ' . $apiKey,
    'Content-Type: application/json'
]);

$response = curl_exec($ch);
$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($httpCode >= 200 && $httpCode < 300) {
    echo "Berhasil: " . $response;
} else {
    echo "Gagal HTTP Code $httpCode : " . $response;
}
?>
```

---

## 13. 📋 Struktur Webhook Payload (Data Balasan)
Jika Anda menggunakan webhook (seperti Make.com atau n8n), ini adalah bentuk format JSON yang akan dikirim server WhatsMeong ke URL webhook Anda setiap kali ada pesan masuk atau jawaban *Polling*.

**Contoh Payload Pesan Teks/Media:**
```json
{
  "sender": "6285779336660",
  "name": "Hanief Neuralens",
  "chat": "6285779336660@s.whatsapp.net",
  "message": "Halo, saya tertarik dengan layanan Anda",
  "timestamp": "2026-06-27T10:00:00Z"
}
```

**Contoh Payload Jawaban Polling (`type: poll_vote`):**
```json
{
  "sender": "6285779336660",
  "name": "Hanief Neuralens",
  "type": "poll_vote",
  "message": "Predict Ticker", 
  "timestamp": "2026-06-27T10:05:00Z"
}
```
*(Nilai `message` akan persis mengikuti teks opsi polling yang diklik oleh user).*
