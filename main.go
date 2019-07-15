package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	// get slice; points to the same memory location
	sl1 := buf[1:4] // [2, 3, 4]
	sl2 := buf[:2]  // [1, 2]
	sl3 := buf[2:]  // [3, 4, 5]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:] // [1, 2, 3, 4, 5]
	// buf = [9, 2, 3, 4, 5], the same memory
	newBuf[0] = 9

	// now newBuf poits to other memory location
	newBuf = append(newBuf, 6)

	// buf    = [9, 2, 3, 4, 5], unchanged
	// newBuf = [1, 2, 3, 4, 5, 6], changed
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	// copy one slice to other slice
	var emptyBuf []int // len=0, cap=0
	// uncorrect - copies only emptyBuf, it is smallest slice in comparison with other

	copied := copy(emptyBuf, buf) // copied = 0
	fmt.Println(copied, emptyBuf)

	// correct
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	// it is possible to copy part of the slice
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) // ints = [1, 5, 6, 4]
	fmt.Println(ints)
}
