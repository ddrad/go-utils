package sort


func QuickSort(arr []int) {
	startIndex := 0
	endIndex := len(arr) - 1
	quickSort(arr, startIndex, endIndex)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i, j := start, end;
	cur := i - (i-j)/2;
	for ; i < j; {
		{
			for ; i < cur && (arr[i] <= arr[cur]); {
				i++;
			}
			for ; j > cur && (arr[cur] <= arr[j]); {
				j--;
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i];
				if i == cur {
					cur = j;
				}
				if j == cur {
					cur = i;
				}
			}
		}
		quickSort(arr, start, cur);
		quickSort(arr, cur+1, end);
	}
}