use std::fs;

fn parse_wheels(wheels_string: &str) -> Vec<Vec<&str>>{
    let mut wheels : Vec<Vec<&str>> = vec![vec!["s"; 4]; 10];

    let mut y = 0;
    for wheel in wheels_string.split("\n").take(10) {
        let mut x = 0;
        for ins in wheel.split(", ") {
            wheels[y][x] = ins;
            x += 1;
        }
        y += 1;
    }
    return wheels
}

fn main() {
    let text = fs::read_to_string("wheels.txt").expect("Failed to read file");

    let wheels = parse_wheels(&text);
    let mut counts = vec![0; wheels.len()];
    let mut coins = 6;

    let mut ins = wheels[coins%10][counts[coins%10]%4];
    while ins != "STOPP" {
        counts[coins%10] += 1;
        match ins {
            "ROTERODDE" => {
                for i in (1..10).step_by(2) {
                    counts[i] += 1;
                }
            },
            "TREKK1FRAODDE" => {
                let mut new_coins = coins;
                let mut mul = 1;
                for (i,c) in ins.to_string().chars().rev().enumerate() {
                    if c == '-' {
                        new_coins *= -1
                    }
                    if c.to_digit(10).unwrap() % 2 == 1 {
                        new_coins


                    }
                }
            },
            "GANGEMSD" => coins *= coins.to_string().replace("-","").chars().take(1).collect::<String>().parse::<usize>().unwrap(),
            "MINUS1" => coins -= 1,
            "MINUS9" => coins -= 9,
            "PLUSS4" => coins += 4,
            _ => {
                println!("{}", ins);
                return
            }
        }

        println!("{}",coins);

        ins = wheels[coins%10][counts[coins%10]%4];

    }
}
