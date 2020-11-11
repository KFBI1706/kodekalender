use std::fs;
use std::collections::HashSet;

fn main() {
    let maze = fs::read_to_string("MAZE.TXT").expect("Failed to open file");
    let mp = json::parse(&maze).unwrap();

    let order = vec!["syd", "aust", "vest, nord"];

    let mut x = 0;
    let mut y = 0;
    let mut c = (y,x);

    let mut stack: Vec<(usize,usize)> = vec![];
    let mut visited = HashSet::new();

    while !(x == 499 && y == 499) {
        visited.insert(c);
        y = c.0;
        x = c.1;
        let mut found = false;
        for dr in order.iter() {
            let (x1, y1) : (usize, usize) = match *dr {
                "syd" => (x,y+1),
                "nord" => (x,y-1),
                "vest" => (x-1,y),
                "aust" => (x+1,y),
                __ => (600,600),
            };

            if mp[y][x][*dr] == false && !visited.contains(&(y1,x1)) {
                //println!("{} {}", c.0,c.1);
                stack.push(c);
                //x = x1;
                //y = y1;
                //
                //println!("pushed {} {}", c.1, c.0);
                c = (y1,x1);
                found = true;
                break
            }
        }
        if !found {
            c = stack.pop().unwrap();
            println!("{:?}", stack);
            println!("poped {} {}", c.1, c.0);
        }

    }

    //while !(x == 499 && y == 499) {
    //    if !visited.contains(&c) {
    //        visited.push(c);
    //    }
    //    x = c.0;
    //    y = c.1;
    //    //println!("{} {}", x, y);
    //    let mut found = false;
    //    for dr in order.iter() {
    //        let (x1, y1) : (usize, usize) = match *dr {
    //            "syd" => (x,y+1),
    //            "nord" => (x,y-1),
    //            "vest" => (x-1,y),
    //            "aust" => (x+1,y),
    //            __ => (600,600),
    //        };
    //        if mp[y][x][*dr] == false && !visited.contains(&(x1,y1)) {
    //            //println!("{} {}", c.0,c.1);
    //            stack.push(c);
    //            //x = x1;
    //            //y = y1;
    //            c = (x1,y1);
    //            found = true;
    //            break
    //        }
    //    }
    //    if !found {
    //        c = stack.pop().unwrap();
    //        println!("{} {}", c.0, c.1);
    //    }
    //}
}
