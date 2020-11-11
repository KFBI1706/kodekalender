
fn main() {
    println!("{}",(1u64..98765433).filter(|&i| {
        let mut n = i;
        let mut sum = 0;
        while n > 0 {
            sum +=  n%10;
            n /= 10;
        }
        i % sum == 0 && primes::is_prime(sum)
    }).count());
}
