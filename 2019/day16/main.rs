use std::fs;

fn main() {
    let lines: Vec<String> = fs::read_to_string("fjord.txt").expect("Failed to open file")
        .split("\n")
        .map(|s| s.to_string())
        .collect();

    let cols = (0..lines[0].len()).map(|i| {
        lines.iter().map(|s| s.chars().skip(i).next().unwrap()).collect::<String>()
    }).collect::<Vec<String>>();

    let mut y: i64 = cols[1].find("B").unwrap() as i64;
    let mut dy: i64 = -1;
    let mut cnt:i64 = 0;

    for x in 1..cols.len() {
        if cols[x].chars().skip((y + dy * 3) as usize).next().unwrap() == '#' {
            cnt += 1;
            dy *= -1;
        }else {
            y += dy
        }
    }
    println!("{}",cnt)
}
