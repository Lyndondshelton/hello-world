fn main(){
	let mut arr:[u32;10] = [1,45,6,12,3, 20, 16, 11, 4,2];
	let arr_len = arr.len();

	for i in 0..arr_len-1{
		for j in 0..arr_len-i-1{
			if arr[j] > arr[j+1]{
				let temp = arr[j];
				arr[j] = arr[j+1];
				arr[j+1] = temp;
			}
		}
	}

	println!("Array (n={}): ", arr_len);
	for n in 0..arr_len{
		println!("{}", arr[n]);
	}
}
