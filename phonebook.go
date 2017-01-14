package main

import (
	"fmt"
	"os/exec"
	"os"
)

var phone_book map[int]map[string]string = make(map[int]map[string]string)

func main() {
	giris_ekrani()
}

func giris_ekrani()  {
	clear()

	var islem_numarasi int
	var kayit_sayisi int
	kayit_sayisi = record_count()

	fmt.Print("CenkSoft Telefon Rehberi v1.0\nHoşgeldiniz, Gerçekleştirmek istediğiniz işlemi seçiniz :\n")
	fmt.Printf("[1] Kayıt Ekle\n[2] Kayıt Listele (%d kayıt)\n[3] Kayıt Sil\nIslem Numarası : ", kayit_sayisi)
	fmt.Scan(&islem_numarasi)

	if islem_numarasi==1 {
		add_record()
	}

	if islem_numarasi==2 {
		list_record()
	}

	if islem_numarasi==3 {
		delete_record()
	}

	giris_ekrani()
}

func record_count() int {
	return len(phone_book)
}

func add_record() {
	clear()

	var record map[string]string = make(map[string]string)

	fmt.Print("Kayıt Ekleme Ekranı\n")
	fmt.Print("--------------------------------------------------------\n\n")

	var name string
	fmt.Print("Adı : ")
	fmt.Scan(&name)
	record["name"] = name

	var surname string
	fmt.Print("Soyadı : ")
	fmt.Scan(&surname)
	record["lastname"] = surname

	var phone string
	fmt.Print("Telefon : ")
	fmt.Scan(&phone)
	record["phone"] = phone

	var key int
	key = len(phone_book)

	phone_book[key] = record

	giris_ekrani()
}

func list_record()  {
	var operasyon string
	clear()
	fmt.Print("Kayıt Listeleme Ekranı\n")
	fmt.Print("--------------------------------------------------------\n\n")


	for key, value := range phone_book {
		fmt.Printf("[%d] - %s %s | %s \n", key, value["name"], value["lastname"], value["phone"] )
	}

	fmt.Print("Guncelle [g], Cıkıs [x]\n")
	fmt.Scan(&operasyon)

	if (operasyon == "x") {
		giris_ekrani()
	} else if(operasyon == "g") {
		var index int
		fmt.Print("Kayıt indexi nedir ?\n")
		fmt.Scan(&index)
		update(index)
	} else {
		list_record()
	}
}

func update(index_number int)  {

	clear()

	fmt.Print("Kayıt Guncelleme Ekranı\n")
	fmt.Print("--------------------------------------------------------\n\n")
	fmt.Print(phone_book[index_number]["name"] + " kaydını güncelliyorsunuz\n")

	record := phone_book[index_number]

	var name string
	fmt.Print("Adı : ")
	fmt.Scan(&name)
	record["name"] = name

	var surname string
	fmt.Print("Soyadı : ")
	fmt.Scan(&surname)
	record["lastname"] = surname

	var phone string
	fmt.Print("Telefon : ")
	fmt.Scan(&phone)
	record["phone"] = phone

	phone_book[index_number] = record

	list_record()
}



func delete_record()  {
	clear()
	giris_ekrani()
}


/**********************************   KURCALAMA MERAK ETME ***********************************/

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
