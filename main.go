package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8000"
	klasor := ""
	fmt.Printf("Oluşturulacak port numarasını giriniz (8000 için boş bırakınız) : ")
	fmt.Scanf("%s", &port)
	if port == "" {
		port = "8000"
	} else if len(port) < 4 {
		port = "8000"
		fmt.Println("Port değerini 4 karakterden az girdiğiniz için port değeri 8000 olarak ayarlandı.")
	} else if len(port) > 4 {
		port = "8000"
		fmt.Println("Port değerini 4 karakterden fazla girdiğiniz için port değeri 8000 olarak ayarlandı.")
	}
	fmt.Printf("Klasor adı giriniz : ")
	fmt.Scanf("%s", &klasor)
	http.Handle("/", http.FileServer(http.Dir(klasor)))
	fmt.Printf("http:localhost: %s oluşturuldu \n", port)
	http.ListenAndServe(":"+port, nil)

}
