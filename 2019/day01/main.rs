use std::fs;

fn main() {
    let mut dragon: i32 = 50;
    let mut overflow: i32 = 0;
    let mut berserk_counter: i32 = 0;

    let sheep_cnt: Vec<i32> = fs::read_to_string("sau.txt")
        .expect("Failed to read file").split(", ")
        .map(|s| s.parse().unwrap())
        .collect();

    for (i, sheep) in sheep_cnt.iter().enumerate() {
        if *sheep+overflow >= dragon {
            overflow += sheep-dragon;
            berserk_counter = 0;
            dragon+=1;
        }else{
            overflow = 0;
            berserk_counter += 1;
            dragon-=1;
            if berserk_counter == 5 {
                println!("Found {}",i);
                break
            }
        }
    }
}