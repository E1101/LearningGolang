package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'verifyItems' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING_ARRAY origItems
 *  2. FLOAT_ARRAY origPrices
 *  3. STRING_ARRAY items
 *  4. FLOAT_ARRAY prices
 */

func verifyItems(origItems []string, origPrices []float32, items []string, prices []float32) int32 {

	var r int32

	// create map of given items,
	// so we can easily search against item names
	var is = make(map[string]int)
	for i, name := range items {
		is[name] = i
	}

	for itemIndex, itemName := range origItems {
		if i, ok := is[itemName]; ok {
			// this element has value in list
			if origPrices[itemIndex] != prices[i] {
				r++
			}
		}
	}


	return r
}



func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	origItemsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var origItems []string

	for i := 0; i < int(origItemsCount); i++ {
		origItemsItem := readLine(reader)
		origItems = append(origItems, origItemsItem)
	}

	origPricesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var origPrices []float32

	for i := 0; i < int(origPricesCount); i++ {
		origPricesItemTemp, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
		checkError(err)
		origPricesItem := float32(origPricesItemTemp)
		origPrices = append(origPrices, origPricesItem)
	}

	itemsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var items []string

	for i := 0; i < int(itemsCount); i++ {
		itemsItem := readLine(reader)
		items = append(items, itemsItem)
	}

	pricesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var prices []float32

	for i := 0; i < int(pricesCount); i++ {
		pricesItemTemp, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
		checkError(err)
		pricesItem := float32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	result := verifyItems(origItems, origPrices, items, prices)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
