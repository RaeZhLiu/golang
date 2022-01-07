package main

import (
	"bufio"
	"flag"
	"fmt"
	"goyard/sorter/algorithms/bubblesort"
	"goyard/sorter/algorithms/qsort"
	"io"
	"os"
	"strconv"
	"time"
)

//从命令行指定输入的数据文件和输出的数据文件，并指定对应的排序算法。该程序的用法如下所示:
//USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>
//一个具体的执行过程如下:
//$ ./sorter –I in.dat –o out.dat –a qsort
//The sorting process costs 10us to complete.
//当然，如果输入不合法，应该给出对应的提示

//1. 获取并解析命令行输入;flag包实现了命令行参数的解析
var infile *string = flag.String("i", "infile", "File contains values of sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

//2. 从对应inFile文件中读取输入数据;
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", file)
		return
	}
	defer func() { _ = file.Close() }()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, readErr := br.ReadLine()
		if readErr != nil {
			if readErr != io.EOF {
				err = readErr
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line)
		value, convErr := strconv.Atoi(str)
		if convErr != nil {
			err = convErr
			return
		}
		values = append(values, value)
	}
	return
}

//4. 将排序的结果输出到对应的文件中;
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile) // ignore_security_alert
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer func() { _ = file.Close() }()

	for _, value := range values {
		str := strconv.Itoa(value)
		_, _ = file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort", "qSort":
			qsort.QuickSort(values)
		case "bubblesort", "BubbleSort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

		_ = writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

