use std::fs;

fn main(){
    let mut map = [[0; 1000] ; 1000];
    let mut px = 0;
    let mut py = 0;
    let mut min = 0;

    let cords = fs::read_to_string("coords.csv")
        .expect("Failed to read file")
        .split("\n").skip(1)
        .map(|s| s.split(",").map(|n| n.parse::<usize>().expect("parse error")).collect::<Vec<usize>>() )
        .collect::<Vec<Vec<usize>>>();

    let mut inc : i32;
    for cord in cords {
        inc = if px < cord[0] { 1 } else { -1 };
        while px != cord[0] {
            px = (px as i32 + inc) as usize;
            map[px][py] += 1;
            min += map[px][py];
        }

        inc = if py < cord[1] { 1 } else { -1 };
        while py != cord[1] {
            py = (py as i32 + inc) as usize;
            map[px][py] += 1;
            min += map[px][py];
        }
    }
    println!("min {}", min);
}