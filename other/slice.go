package other

import (
	"bytes"
	"fmt"
)

// 重新切出新的slice时，新slice会引用原slice的底层数组，如果是大容量对象可能会分配打内存
func Reslice() {
	data := get1()
	fmt.Println(len(data), cap(data), &data[0])
}
func get1() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

// 针对大容量slice切出小slice时，用copy解决
func CopySlice() {
	data := get2()
	fmt.Println(len(data), cap(data), &data[0])
}
func get2() (res []byte) {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res = make([]byte, 3)
	copy(res, raw[:3])
	return
}

func JoinSlice() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	println(sepIndex)
	//dir1的容量为path的总容量，未重新分配一个buffer来保存 X
	//dir1 := path[:sepIndex]
	dir1 := path[:sepIndex:sepIndex]
	println("dir1 cap:", len(dir1), cap(dir1))
	dir2 := path[sepIndex+1:]
	println("dir1:", string(dir1))
	println("dir2:", string(dir2))
	println("dir1 point:", &dir1[0])

	dir1 = append(dir1, "suffix"...)
	println("current path:", string(path))
	println("dir1 point:", &dir1[0])
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1:", string(dir1))
	println("dir2:", string(dir2))
	println("new path:", string(path))
}
