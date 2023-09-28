package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Sunucu IP adresini alın
	ip := getLocalIP()

	// Yönlendirilecek URL'i belirtin
	url := fmt.Sprintf("http://%s:8080", ip)

	// QR kodu oluşturun
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		fmt.Println("qr code creation error: ", err)
		os.Exit(1)
	}

	// QR kodu görüntüleyin (opsiyonel)
	fmt.Println(qr.ToSmallString(false))

	// QR kodu bir dosyaya kaydedin
	err = qr.WriteFile(256, "./QRcode.png")
	if err != nil {
		fmt.Println("Record error:", err)
		os.Exit(1)
	}

	fmt.Printf("QR code created and saved in file named 'QRcode.png'.\n")

	// 8080 portunda basit bir HTTP sunucusu başlatın
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you are on port 8080")
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("HTTP server initialization error: ", err)
			os.Exit(1)
		}
	}()

	select {}
}

// Sunucunun yerel IP adresini alır
func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("IP error:", err)
		os.Exit(1)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
