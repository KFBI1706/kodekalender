extern crate image;

use image::{ImageBuffer, RgbImage};

use std::fs;

fn main (){
    let img = fs::read_to_string("img.txt")
        .expect("Failed to read file");
    let len = img.chars().count() as u32;
    let mut n :u32 = 100;
    loop {
        println!("{}",n);
        if n == 2000 {
            break
        }
        if len%n == 0  {
            println!("creating image {} {}", n,len/n);
            let mut image : RgbImage = ImageBuffer::new(n,len/n);

            for (x,y,pixel) in image.enumerate_pixels_mut() {
                let i = y*(len/n)+x;
                let c: String = img.chars().skip(i as usize).take(1).collect();
                if c  == "1" {
                    *pixel = image::Rgb([255,255,255])
                }
            }
            image.save("/tmp/output.png").unwrap();
            println!("done");
            break
        }
        n += 1;
    }
}