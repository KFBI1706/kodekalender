fn main() {
    let mut speed = 10703437.0;
    let mut ice = 0.0;

    for (i,c) in std::fs::read_to_string("terreng.txt").expect("F").trim().chars().enumerate() {
        if speed <= 0.0 {
            println!("{}",i);
            break
        }
        ice = if c != 'I' { 0.0 } else {ice + 1.0};
        match c {
            'G' => speed -= 27.0,
            'A' => speed -= 59.0,
            'S' => speed -= 212.0,
            'F' => speed -= 17.5,
            'I' => speed += 12.0*ice,
            _ => panic!(format!("`{}` char not found", c)),
        }
    }
}
