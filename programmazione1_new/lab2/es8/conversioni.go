package main

import "fmt"

func main() {
	fmt.Println("Scegli la conversione: \n 1] secondi>ore \n 2] secondi>minuti \n 3] minuti>ore \n 4] minuti>seconi \n 5] ore>secondi \n 6] ore>minuti \n 7] minuti>giorni e ore \n 8]minuti>anni e giorni")
	var scelta, valore int
	var conversione float64
	fmt.Scan(&scelta)
	if (scelta < 1) || (scelta >= 9) {
		fmt.Println("Scelta errata")
	}

	fmt.Print("Inserisci il valore da convertire: ")
	fmt.Scan(&valore)
	switch scelta {
	case 1: //secondi>ore
		conversione = float64(valore) / 3600
		fmt.Println(valore, "secondi corrispondono a ", conversione, " ore.")
	case 2: //secondi>minuti
		conversione = float64(valore) / 60
		fmt.Println(valore, "secondi corrispondono a ", conversione, " minuti.")
	case 3: //minuti>ore
		conversione = float64(valore) / 60
		fmt.Println(valore, "minuti corrispondono a ", conversione, " ore.")
	case 4: //minuti>secondi
		conversione = float64(valore) * 60
		fmt.Println(valore, "minut corrispondono a ", conversione, " secondi.")
	case 5: //ore>secondi
		conversione = float64(valore) * 3600
		fmt.Println(valore, "ore corrispondono a ", conversione, " secondi.")
	case 6: //ore>minuti
		conversione = float64(valore) * 60
		fmt.Println(valore, "ore corrispondono a ", conversione, " minuti.")
	case 7: //minuti>giorni e ore
		conversionegiorni := float64(valore) / 1440
		conversioneore := float64(valore%1440) / 60
		fmt.Println(valore, "minut corrispondono a ", conversionegiorni, " giorni e ", conversioneore, "ore.")
	case 8: //minut>anni e giorni
		conversioneanni := float64(valore) / 525600
		if conversioneanni < 1 {
			conversioneanni = 0
		}
		conversionegiorni := float64(valore%525600) / 1440
		fmt.Println(valore, "secondi corrispondono a ", conversioneanni, " anni e", conversionegiorni, " giorni.")
	}
}
