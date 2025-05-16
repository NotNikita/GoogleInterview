package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    RUtoUS = iota
    RUtoEU
    UStoRU
    UStoEU
    EUtoRU
    EUtoUS

    RU = 0
    US = 1
    EU = 2
)

func main() {
    in, out := getBuffers()
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)
    
    for testCase := 0; testCase < t; testCase++ {
        // Считываем данные для трех банков
        banks := make([][6]float64, 3)
        
        for bank := 0; bank < 3; bank++ {
            for j := 0; j < 6; j++ {
                var n, m int
                fmt.Fscan(in, &n, &m)
                banks[bank][j] = float64(m) / float64(n)
            }
			// copy(banks[1][:], banks[0][:])
			// copy(banks[2][:], banks[0][:])

        }
        
        // Находим максимальное количество долларов
        maxDollars := findMaxDollars(banks)
        fmt.Fprintln(out, maxDollars)
    }
}

func findMaxDollars(banks [][6]float64) float64 {
    // Начинаем с 1 рубля
    maxUSD := 0.00
    
    // Перебираем все возможные последовательности банков
    // (всего 3! = 6 возможных последовательностей)
    bankSequences := [][]int{
        {0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0},
    }
    
    for _, sequence := range bankSequences {
        // Для каждой последовательности банков перебираем все возможные пути конвертации
        // Рассматриваем все варианты начальных конвертаций:
        
        // Вариант 1: Сразу конвертируем рубль в доллар в первом банке
        rubles := 1.0
        dollars := rubles * banks[sequence[0]][RUtoUS]
        maxUSD = max(maxUSD, tryPaths(sequence, 1, US, 0, dollars, 0, banks))
        
        // Вариант 2: Сначала конвертируем рубль в евро в первом банке
        euros := rubles * banks[sequence[0]][RUtoEU]
        maxUSD = max(maxUSD, tryPaths(sequence, 1, EU, 0, 0, euros, banks))
        
        // Вариант 3: Не конвертируем в первом банке (пропускаем этот вариант, т.к. нет смысла)
        // maxUSD = max(maxUSD, tryPaths(sequence, 1, RU, rubles, 0, 0, banks))
    }
    
    return maxUSD
}

// Рекурсивная функция для перебора всех возможных путей конвертации
func tryPaths(sequence []int, bankIndex int, currency int, rubles, dollars, euros float64, banks [][6]float64) float64 {
    // Если прошли все банки, возвращаем количество долларов
    if bankIndex >= len(sequence) {
        return dollars
    }
    
    bank := sequence[bankIndex]
    maxUSD := dollars
    
    // В зависимости от текущей валюты пробуем разные варианты конвертации
    switch currency {
    case RU:
        // Можем конвертировать рубли в доллары
        newDollars := dollars + rubles * banks[bank][RUtoUS]
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, US, 0, newDollars, euros, banks))
        
        // Можем конвертировать рубли в евро
        newEuros := euros + rubles * banks[bank][RUtoEU]
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, EU, 0, dollars, newEuros, banks))
        
        // Можем не конвертировать
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, RU, rubles, dollars, euros, banks))
    
    case US:
        // Можем конвертировать доллары в рубли
        newRubles := rubles + dollars * banks[bank][UStoRU]
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, RU, newRubles, 0, euros, banks))
        
        // Можем конвертировать доллары в евро
        newEuros := euros + dollars * banks[bank][UStoEU]
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, EU, rubles, 0, newEuros, banks))
        
        // Можем не конвертировать
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, US, rubles, dollars, euros, banks))
    
    case EU:
        // Можем конвертировать евро в рубли
        newRubles := rubles + euros * banks[bank][EUtoRU]
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, RU, newRubles, dollars, 0, banks))
        
        // Можем конвертировать евро в доллары
        newDollars := dollars + euros * banks[bank][EUtoUS]
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, US, rubles, newDollars, 0, banks))
        
        // Можем не конвертировать
        maxUSD = max(maxUSD, tryPaths(sequence, bankIndex+1, EU, rubles, dollars, euros, banks))
    }
    
    return maxUSD
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}