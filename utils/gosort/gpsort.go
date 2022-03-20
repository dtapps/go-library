package gosort

// QuickSort 快速排序算法
// https://www.topgoer.com/Go%E9%AB%98%E7%BA%A7/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95.html
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitdata) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}
