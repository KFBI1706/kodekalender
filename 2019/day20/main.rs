use primes::is_prime;

fn next(elf: usize, inc: i8) -> usize {
    let new_elf = (elf as i8) + inc;
    return if new_elf == -1 { 4 } else if new_elf == 5 { 0 } else { new_elf as usize}

}

fn main() {
    let mut rule: u8;
    let mut rules: Vec<u8> = vec![0; 1];
    let (mut elves, mut sorted): (Vec<u32>, Vec<u32>) = (vec![0; 5], vec![0;5]);
    let mut elf = 0;
    let mut inc : i8 = 1;

    for step in 1..1000740 {
        let curr_rule = rules.last().unwrap();

        match  curr_rule {
            5 => {
                elf = next(elf,inc);
            },
            4 => {
                elf = 4;
            },
            3 => {
                for _ in 0..2 {
                    elf = next(elf,inc);
                }
            },
            2 => {
                inc = if inc == 1 { -1 } else { 1 }
            },
            1 => {
                elf = elves.iter().position(|&p| p == sorted[0]).unwrap();
            },
            0 => println!("First elf does nothing"),
            _ => {
                println!("FAILED, encountered unknown {}", curr_rule);
                break
            }
        }

        println!("{} {} {} ", elf+1, step, curr_rule);
        elves[elf] += 1;

        sorted = elves.clone();
        sorted.sort();
        //println!("{:?}", elves);
        //println!("{:?}",sorted);
        rule = 5;
        if is_prime(step+1) && elves.iter().filter(|&n| *n == sorted[0]).count() == 1  {
            rule = 1
        } else if (step+1)%28 == 0 {
            rule =2
        } else if (step+1)%2 == 0 && elves.iter().filter(|&n| *n == sorted[4]).count() == 1 && sorted[4] == elves[next(elf, inc)] {
            rule = 3
        } else if (step+1)%7 == 0 {
            rule = 4
        }else{
        }

        rules.push(rule);
        if step == 9 {
            break
        }
    }
    println!("{:?} {:?}",elves, sorted);
}
