use std::fs;

fn main() {
    let generations : Vec<Vec<Vec<u32>>> = fs::read_to_string("generations.txt.liten").expect("Failed to read file")
        .split("\n").take(2)
        .map(|generation| generation.split(";")
            .map(|parents| parents.split(",")
                .map(|parent| parent.parse().unwrap()).collect::<Vec<u32>>()
            ).collect::<Vec<Vec<u32>>>()
        ).collect::<Vec<Vec<Vec<u32>>>>();

    for (i,generation) in generations.iter().enumerate() {
        println!("{} {:?}",i, generation);
        for child in 0..generation.len() {
            println!("{}",child);
            for (j, parent) in generation.iter().enumerate() {

            }
        }
    }
}
