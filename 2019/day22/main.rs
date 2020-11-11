extern crate regex;

use regex::Regex;
use std::fs;

fn main(){
    let tree = Regex::new(r"#( |\n)").unwrap();
    let cnt = tree.find_iter(fs::read_to_string("forest.txt.liten").unwrap().as_str()).count();
    println!("{}", cnt*40);
}
