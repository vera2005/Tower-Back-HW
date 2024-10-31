package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	count         bool
	deplicate     bool
	uniqStrings   bool
	numFields     int
	numChars      int
	ignoreRegistr bool
}

// создаем слайс строк. в котором изменяем исходные на
// соответсвующие параметрам пропусков полей, символов, игнора регистра
func Format(lines []string, opts Options) []string {
	var ans []string
	for _, line := range lines {
		if opts.numFields > 0 {
			// смещение полей
			line = SkipFields(line, opts.numFields)
		}
		if opts.numChars > 0 {
			//cмещение символов
			line = SkipChars(line, opts.numChars)
		}
		// игнор регистра
		if opts.ignoreRegistr {
			line = strings.ToLower(line)
		}
		ans = append(ans, line)
	}
	return ans
}

// функция смещения полей
func SkipFields(line string, numFields int) string {
	// разбиение строки по пробелам, получи срез строк
	lines := strings.Fields(line)
	if numFields >= len(lines) {
		return ""
	}
	return strings.Join(lines[numFields:], " ")

}

// функция смещения символов
func SkipChars(line string, numChars int) string {
	if numChars >= len(line) {
		return ""
	}
	return line[numChars:]
}

// формирование списка для вывода в соответвии с параметрами count , deplicate , uniqStrings
func FormResult(lines []string, opts Options) []string {
	var ans []string
	formatLines := Format(lines, opts)

	if opts.count {
		CountOccurrences(formatLines, lines, &ans)
	} else if opts.deplicate {
		FindDuplicates(formatLines, lines, &ans)
	} else if opts.uniqStrings {
		FindUnique(formatLines, lines, &ans)
	} else {
		Find(formatLines, lines, &ans)
		fmt.Println(ans, lines, formatLines)
	}
	return ans
}
func Find(formatLines []string, lines []string, result *[]string) {
	*result = append(*result, lines[0])
	prev := formatLines[0]
	for i := 1; i < len(formatLines); i++ {
		if formatLines[i] != prev {
			*result = append(*result, lines[i])
		}
		prev = formatLines[i]
	}
}

func CountOccurrences(formatLines []string, lines []string, result *[]string) {
	count := 1
	for i := 1; i < len(formatLines); i++ {
		if formatLines[i] == formatLines[i-1] {
			count++
		} else {
			if count > 0 {
				*result = append(*result, fmt.Sprintf("%d %s", count, lines[i-1]))
			}
			count = 1
		}
	}
	if count > 0 { // Обработка последней группы
		*result = append(*result, fmt.Sprintf("%d %s", count, lines[len(lines)-1]))
	}
}

func FindDuplicates(formatLines []string, lines []string, result *[]string) {
	count := 1
	for i := 1; i < len(formatLines); i++ {
		if formatLines[i] == formatLines[i-1] {
			count++
			if count == 2 { // Добавляем только первое вхождение дубликатов
				*result = append(*result, lines[i-1])
			}
		} else {
			count = 1
		}
	}
}

func FindUnique(formatLines []string, lines []string, result *[]string) {
	count := 1
	for i := 1; i < len(lines); i++ {
		if formatLines[i] == formatLines[i-1] {
			count++
		} else {
			if count == 1 { // Добавляем только уникальные вхождения
				*result = append(*result, lines[i-1])
			}
			count = 1
		}
	}
	if count == 1 { // Обработка последней строки если она уникальна
		*result = append(*result, lines[len(lines)-1])
	}
}

// Функция для парсинга флагов командной строки
func ParseFlags() Options {
	// определение флагов - описываем поддерживаемые флаги(тип, имя, значение по умлочанию, описание)
	count := flag.Bool("c", false, "count occurrences of each line")
	deplicate := flag.Bool("d", false, "output only duplicate lines")
	uniqStrings := flag.Bool("u", false, "output only unique lines")
	numFields := flag.Int("f", 0, "skip the first num_fields fields in each line")
	numChars := flag.Int("s", 0, "skip the first num_chars characters in each line")
	ignoreRegistr := flag.Bool("i", false, "ignore case when comparing lines")
	// анализ и присвоение значения переменным, связанным с флагами
	flag.Parse()
	// проверка конфликтов
	if *deplicate && *uniqStrings {
		fmt.Println("uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		os.Exit(1)
	}
	if *count && (*deplicate || *uniqStrings) {
		fmt.Println("uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		os.Exit(1)
	}
	// возвращаем структуру, которая содержит значения всех обработанных параметоров
	return Options{
		count:         *count,
		deplicate:     *deplicate,
		uniqStrings:   *uniqStrings,
		numFields:     *numFields,
		numChars:      *numChars,
		ignoreRegistr: *ignoreRegistr,
	}
}

// Чтение входных данных из файла
func ReadInputFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %s\n", err)
		os.Exit(1)
	}

	return lines
}

// Чтение входных данных из stdin
func ReadInputFromStdin() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Запись результата в файл
func WriteOutputToFile(lines []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, line := range lines {
		fmt.Fprintln(file, line)
	}
}

func main() {
	opts := ParseFlags()
	var lines []string

	if len(flag.Args()) > 0 {
		inputFile := flag.Args()[0]
		lines = ReadInputFromFile(inputFile)
	} else {
		lines = ReadInputFromStdin()
	}

	result := Format(lines, opts)
	finalResult := FormResult(result, opts)

	if len(flag.Args()) > 1 {
		outputFile := flag.Args()[1]
		WriteOutputToFile(finalResult, outputFile)
	} else {
		for _, line := range finalResult {
			fmt.Println(line)
		}
	}
}
