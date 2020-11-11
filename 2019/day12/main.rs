
fn main() {
    let mut cnt;
    for i in 1000..10000 {
        if i%1111 != 0 {
            cnt = 1;
            loop {
                let n = i.to_string()
                    .chars()
                    .map(|c| c.to_digit(10).unwrap())
                    .collect::<Vec<u32>>();
                let mut n1;
                let mut n2;
                for i in 1usize..4 {
                    n1 += n[i]*10u32.pow(i as u32);
                    println!("{} {}", n1,n2);
                }
            }
        }
        break
    }
}
