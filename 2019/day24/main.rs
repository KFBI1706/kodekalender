use std::fs;
use std::usize;

const SIZE: usize = 350;

fn main() {
    for (i,bokstav) in fs::read_to_string("turer.txt").expect("Failed to open file").split("---\n").enumerate() {
        let mut mp = vec![vec!['.';SIZE]; SIZE];
        let mut maxx = 0;
        let mut maxy = 0;
        let mut minx = usize::MAX;
        let mut miny = usize::MAX;
        for cord in bokstav.split("\n").filter(|&s| s != "") {
            let v = cord.split(",").map(|n| n.parse().unwrap()).collect::<Vec<usize>>();
            mp[SIZE-1-v[1]][v[0]] = 'x';
            minx = if v[0] < minx { v[0] } else { minx };
            maxx = if v[0] > maxx { v[0] } else { maxx };
            miny = if v[1] < miny { v[1] } else { miny };
            maxy = if v[1] > maxy { v[1] } else { maxy };
        }

        if i == 0 || i == 11 || i == 17 || i == 20 || i == 31 || i == 33 {
            for line in mp.into_iter().skip(SIZE-maxy-1).take(maxy-miny+1) {
                println!("{}", line.into_iter().skip(minx).take(maxx-minx+1).collect::<String>());
            }
        }
    }
}
