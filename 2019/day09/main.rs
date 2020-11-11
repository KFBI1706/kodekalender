use std::fs;

fn main(){
    println!("Sum: {}", fs::read_to_string("krampus.txt")
        .expect("Failed to open file")
        .split("\n")
        .take(1000000)
        .map(|s| s.parse::<usize>().unwrap())
        .filter(|value| {
            let pow = value*value;
            (2..7)
                .map(|i| 10_usize.pow(i))
                .take_while(|v| &pow > v)
                .any(|v| {
                    let last = pow % v;
                    last != 0 && &(pow / v + last) == value
                })
        })
        .sum::<usize>());
}
