extern crate image;

use image::{GenericImageView, ImageBuffer};

fn main() {
    let img = image::open("mush.png").unwrap();
    let (width, height) = img.dimensions();

    let mut px = img.get_pixel(0,0);
    ImageBuffer::from_fn(width,height, |x, y| {
        let mut pixel = img.get_pixel(x,y);
        if !(x == 0 && y == 0) {
            for i in 0..3 {
                pixel[i] ^= px[i];
            }
        }
        px = img.get_pixel(x,y);
        pixel
    }).save("hammer.png").unwrap();
}
