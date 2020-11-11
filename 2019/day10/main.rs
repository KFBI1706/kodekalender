extern crate chrono;

//use std::collections::HashMap;
use chrono::{Duration, Datelike, Weekday,NaiveDate};
use std::str::FromStr;
use std::{fs};
use Product::{Toalettpapir, Sjampo, Tannkrem};
use Unit::{ML,M};

#[derive(Debug)]
enum Unit {
    ML,
    M,
}

#[derive(Debug)]
enum Product {
    Toalettpapir,
    Sjampo,
    Tannkrem,
}

#[derive(Debug)]
struct Line  {
    amount: u32,
    unit: Unit,
    product: Product
}

#[derive(Debug)]
struct LineErr {
    details: String
}

impl LineErr {
    fn new(msg: &str) -> LineErr {
        LineErr{details: msg.to_string()}
    }
}

impl FromStr for Line {
    type Err = LineErr;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let props: Vec<&str> = s.split(" ").collect();

        let product_fromstr = match props[3] {
            "tannkrem" =>  Tannkrem,
            "sjampo" =>  Sjampo,
            "toalettpapir" =>  Toalettpapir,
            _ => return Result::Err(LineErr::new(format!("Failed to parse `{}`", props[3]).as_str())),
        };

        let amount_fromstr = props[1].parse::<u32>().unwrap();

        let unit_fromstr = match props[2] {
            "ml" =>  ML,
            "meter" => M,
            _ => return Result::Err(LineErr::new(format!("Failed to unit `{}`", props[2]).as_str())),
        };

        Ok(Line{amount:amount_fromstr, product: product_fromstr, unit:unit_fromstr})
    }
}

struct Log {
    date: NaiveDate,
    //counters: HashMap<Product, HashMap<Weekday, u32>>,
}


fn main() {
    let lines : Vec<String> = fs::read_to_string("logg.txt")
        .expect("Failed to open file")
        .split("\n")
        .map(|s| s.to_string())
        .collect();

    let mut log = Log{
        date: NaiveDate::from_ymd(2017, 12, 31),
        //counters: HashMap::new(),
    };

    let mut tannkrem = 0;
    let mut sjampo = 0;
    let mut toalettpapir = 0;
    let mut sjampo_sun = 0;
    let mut toalettpapir_sun = 0;

    for (i,ln) in lines.iter().enumerate() {
        if i % 4 == 0 {
            log.date = log.date + Duration::days(1);
            continue
        }
        let line = Line::from_str(ln).unwrap();
        //println!("{:?} {:?} {:?}", line.amount, line.unit, line.product);
        match line.product {
            Tannkrem => tannkrem += line.amount,
            Sjampo => {
                sjampo += line.amount;
                if log.date.weekday() == Weekday::Sun {
                    sjampo_sun += line.amount
                }
            },
            Toalettpapir => {
                toalettpapir += line.amount;
                if log.date.weekday() == Weekday::Wed {
                    toalettpapir_sun += line.amount;
                }
            }
        }
    }
    let prod = (tannkrem/125)*(sjampo/300)*(toalettpapir/25)*toalettpapir_sun*sjampo_sun;
    println!("{}",prod);

}
