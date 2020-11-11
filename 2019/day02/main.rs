use std::fs;

fn main() {
    let world = fs::read_to_string("world.txt")
        .expect("Failed to read file");
    let mut water = 0;
    for layer in world.split("\n") {
        water += layer.trim().matches(" ").count();
    }
    println!("{}",water)
}
