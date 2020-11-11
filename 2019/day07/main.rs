
fn main() {
    for x in 2..8 {
        for y_inv in 2..27644436 {
            if (y_inv * x) % 27644437 == 1 {
                let z : i64 = 5897 * y_inv;
                println!("{}: Code: {}", x, z % 27644437);
                break
            }
        }
    }
}
