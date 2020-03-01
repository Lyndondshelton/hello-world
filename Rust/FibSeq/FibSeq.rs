fn main(){
	let n = 30;
	let mut x = 0; // n-2
	let mut y = 0; // n-1
	let mut z = 0; // n

	println!("Xn = Xn-1 + Xn-2");
	println!();
	for i in 0..n+1{
		if i == 0{
			x = 0; //Xn-2
		}else if i == 1{
			y = 1; //Xn-1
		}else{
			z = y + x; // Xn = (Xn-1) + (Xn-2)
			x = y; // (Xn-2)
			y = z; // (Xn-1)
		}
	}
	println!("n = {}, Xn = {}", n, z);
}
