package main

import (
	"fmt"
)

const N = 9

// Sudoku tahtasını ekrana yazdırma fonksiyonu
func SudokuYaz(sudoku [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if sudoku[i][j] == 0 {
				fmt.Printf(". ") // Boş hücre için nokta kullan
			} else {
				fmt.Printf("%d ", sudoku[i][j])
			}
		}
		fmt.Println()
	}
}

// Belirli bir sayının belirli bir pozisyona yerleştirilebilir olup olmadığını kontrol eden fonksiyon
func GecerliSudoku(sudoku [N][N]int, row, col, deger int) bool {
	// Satır ve sütun kontrolü
	for i := 0; i < N; i++ {
		if sudoku[i][col] == deger || sudoku[row][i] == deger {
			return false
		}
	}

	// 3x3 kontrolü
	ilkSatır := row - row%3
	ilkSutun := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[ilkSatır+i][ilkSutun+j] == deger {
				return false
			}
		}
	}
	return true
}

// Sudoku çözme fonksiyonu
func SudokuÇöz(sudoku *[N][N]int) bool {
	row := -1
	col := -1
	empty := false

	// Boş hücreyi bulmak icin
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if sudoku[i][j] == 0 {
				row = i
				col = j
				empty = true
				break
			}
		}
		if empty {
			break
		}
	}

	// Eğer boş hücre yoksa Sudoku çözülmüştür
	if !empty {
		return true
	}

	// 1'den 9'a kadar sayıları deneyerek çözüm ara
	for deger := 1; deger <= 9; deger++ {
		if GecerliSudoku(*sudoku, row, col, deger) {
			sudoku[row][col] = deger // Geçici olarak yerleştir

			if SudokuÇöz(sudoku) { // Geri izleme
				return true
			}

			sudoku[row][col] = 0 // Geri al
		}
	}
	return false
}

func main() {
	var board [N][N]int
	// Girdi alınırken nokta kullanılarak boş hücreler sıfırla doldurulacak
	input := [N]string{
		"34.91..2.",
		".96.8..41",
		"..8.2..7.",
		".6..57.39",
		"1.2.6.7..",
		"97..3..64",
		"45.2.8..6",
		".8..9..5.",
		"6.3..189.",
	}

	// Nokta (.) karakterlerini sıfır (0) ile değiştirir
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if input[i][j] == '.' {
				board[i][j] = 0
			} else {
				var num int
				fmt.Sscanf(string(input[i][j]), "%d", &num)
				board[i][j] = num
			}
		}
	}

	
	if SudokuÇöz(&board) {
		fmt.Println("Sudoku çözüldü:")
		SudokuYaz(board)
	} else {
		fmt.Println("Error: Çözüm bulunamadı.")
	}
}
