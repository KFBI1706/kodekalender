extern crate csv;
use std::collections::HashMap;
use std::fs;
use std::fs::File;

const MAN: usize = 0;
const WOMAN: usize = 1;
const LASTNAME1: usize = 2;
const LASTNAME2: usize = 3;
const ALPHABET:&str = "abcdefghijklmnopqrstuvwxyz";

fn main() {
    let names = fs::read_to_string("names.txt").expect("Failed to open names")
        .split("\n---\n")
        .map(|s| s.split("\n").map(|ss| ss.to_string()).collect::<Vec<String>>())
        .collect::<Vec<Vec<String>>>();

    let mut cnt : HashMap<String,u32> = HashMap::new();

    let file = File::open("employees.csv").unwrap();
    let mut rdr = csv::Reader::from_reader(file);
    for result in rdr.records() {
        let record = result.expect("a CSV record");

        let ix = if record[2].to_string() == "M" { MAN } else { WOMAN };
        let name = record[0].to_string();
        let last_name = record[1].to_string().to_lowercase().chars().filter(|c| c.is_alphabetic() ).collect::<String>();

        let (last_name1, last_name2) = (&last_name[..last_name.len()-last_name.len()/2], &last_name[last_name.len()-last_name.len()/2..]);

        let mut sum: usize = name.as_bytes().iter().map(|&c| c as usize).sum();
        let jedi_name = &names[ix][sum%names[ix].len()];

        sum = last_name1.chars()
            .map(|c| ALPHABET.find(c).unwrap() + 1)
            .sum();
        let jedi_last_name1 = &names[LASTNAME1][sum%names[LASTNAME1].len()];

        sum = last_name2.as_bytes().iter()
            .map(|&c| c as usize).product();
        sum *= if ix == MAN { name.len() } else { name.len() + last_name.len() };
        let mut chars = sum.to_string().chars().collect::<Vec<char>>();
        chars.sort();
        sum = chars.into_iter().rev().collect::<String>().parse().unwrap();
        let jedi_last_name2 = &names[LASTNAME2][sum%names[LASTNAME2].len()];

        *cnt.entry(format!("{} {}{}",jedi_name,jedi_last_name1,jedi_last_name2)).or_insert(0) += 1;
    }

    let mut max: u32 = 0;
    let mut name = "";
    for v in cnt.iter() {
        if v.1 > &max {
            max = *v.1;
            name = v.0
        }
    }
    println!("{} {}",name, max);
}
