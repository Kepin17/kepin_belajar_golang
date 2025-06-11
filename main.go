package main

import (
	"fmt"
)


func main() {

	option := 0

	// uang yang dimiliki player

	// staterpack player
	const money = 100000
	remainingMoney := money

	// daftar item di toko
	// setiap item adalah map dengan nama, harga, dan stok

		shopItems := []map[string]interface{} {
			{
				"nama": "Barbarian",
				"harga": 10000,
				"stok": 10,
			},
			{
				"nama": "Assassin",
				"harga": 15000,
				"stok": 5,
			},
			{
				"nama": "Tank",
				"harga": 12000,
				"stok": 5,
			},
		}


	


	for  {	
		fmt.Print("=======================================\n")
		fmt.Print("              TOKO KEPIN               \n")
		fmt.Print("=======================================\n")
		fmt.Print("1. Beli\n")
		fmt.Print("2. Jual\n")
		fmt.Print("3. Keluar\n")
		fmt.Print("=======================================\n")

		fmt.Print("Masukkan pilihan (1-3): ")
		fmt.Scan(&option)

		if option == 1 {

	

		for i := range shopItems {
			fmt.Printf("%d. %s - Rp%d (Stok: %d)\n", i+1, shopItems[i]["nama"], shopItems[i]["harga"], shopItems[i]["stok"])
		}
		fmt.Printf("Pilih item yang ingin dibeli (1-%d): ", len(shopItems))
		var itemchoice int
		fmt.Scan(&itemchoice)
		if itemchoice > len(shopItems) {
			fmt.Println("Pilihan tidak valid.")
			continue
		}else {
			item := shopItems[itemchoice-1]
			price := item["harga"].(int)
			stok := item["stok"].(int)	

			if stok <= 0 {
				fmt.Println("Stok tidak tersedia.")
				continue
			}

			var quantity int
			fmt.Printf("Masukkan jumlah %s yang ingin dibeli: ", item["nama"])
			fmt.Scan(&quantity)
			if quantity <= 0 {
				fmt.Println("Jumlah tidak valid.")
				continue
			}

			if quantity > stok {
				fmt.Printf("Maaf, stok %s tidak mencukupi. Stok saat ini: %d\n", item["nama"], stok)
				continue
			}
			totalPrice := price * quantity
			if totalPrice > remainingMoney {
				fmt.Printf("Maaf, uang Anda tidak mencukupi. Total harga: Rp%d\n", totalPrice)
				continue
			}

			// update stok 
			

			item["stok"] = stok - quantity
			fmt.Printf("Anda telah membeli %d %s dengan total harga Rp%d. Sisa uang Anda: Rp%d\n", quantity, item["nama"], totalPrice, remainingMoney - totalPrice)
			remainingMoney -= totalPrice
			fmt.Printf("Stok %s sekarang: %d\n", item["nama"], item["stok"].(int) - quantity)
			continue
		}

		

		
	}
		
}
}