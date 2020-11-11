fn rev(mut n:u64) -> u64 {
    let mut rev = 0;
    while  n > 0  {
        rev = rev * 10 + n % 10;
        n /= 10
    }
    return rev
}

fn main() {
    let mut sum: u64 = 0;
    for i in 1u64..123454321 {
        let r = rev(i);
        if i != r && i+r == rev(i+r) {
            sum += i;
        }
    }
    println!("{}", sum);
}
