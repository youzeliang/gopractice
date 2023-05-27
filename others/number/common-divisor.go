package number

// greatest common divisor

//  欧几里得辗转相除法：
//
//gcd(x,y)表示x和y的最大公约数
//
//进入运算时:x!=0,y!=0，本质上就是不断转换成求等价更小数的最大公约数。如果x%y=0，返回y，即最大公约数。
//gcd(x,y)=gcd(y,x%y)
//
//证明：
//设k=x/y，b=x%y  则：x=ky+b
//如果n能够同时整除x和y，则(y%n)=0,(ky+b)%n=0，则b%n=0，即n也同时能够整除y和b。
//由上得出：同时能够整除y和(b=x%y)的数，也必然能够同时整除x和y。故而gcd(x,y)=gcd(y,x%y)。当(b=x%y)=0，即y可以整除x，这时的y也就是所求的最大公约数了。
//
//附上两条重要性质：gcd(a,b)=gcd(b,a)，gcd(-a,b)=gcd(a,b)

/*
*   穷举法：最大公约数
 */
func gcdNormal(x, y int) int {
	var n int
	if x > y {
		n = y
	} else {
		n = x
	}
	for i := n; i >= 1; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1
}

/*
*辗转相除法：最大公约数
*递归写法，进入运算是x和y都不为0
 */
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

/*
*辗转相除法：最大公约数
*非递归写法
 */
func gcdx(x, y int) int {
	var tmp int
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

/*
*穷举写法：最小公倍数
 */
func lcmNormal(x, y int) int {
	var top int = x * y
	var i = x
	if x < y {
		i = y
	}
	for ; i <= top; i++ {
		if i%x == 0 && i%y == 0 {
			return i
		}
	}
	return top
}

/*
*公式解法：最小公倍数=两数之积/最大公约数
 */
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
