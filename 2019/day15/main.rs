//https://stackoverflow.com/questions/1406029/how-to-calculate-the-volume-of-a-3d-mesh-object-the-surface-of-which-is-made-up/1568551#1568551
//https://n-e-r-v-o-u-s.com/blog/?p=4415
use std::fs;
use std::str::FromStr;
use std::num::ParseFloatError;

struct Triangle {
    a: Point,
    b: Point,
    c: Point
}

struct Point {
    x: f64,
    y: f64,
    z: f64
}

impl FromStr for Triangle {
    type Err = ParseFloatError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let coords: Vec<f64> = s.split(',')
                                .map(|s| s.parse::<f64>().unwrap())
                                 .collect();
        Ok (Triangle {
                a:Point { x: coords[0], y: coords[1], z: coords[2] },
                b:Point { x: coords[3], y: coords[4], z: coords[5] },
                c:Point { x: coords[6], y: coords[7], z: coords[8] }
            }
        )
    }
}

impl Triangle {
    fn volume(self) -> f64 {
        let v321 = self.c.x*self.b.y*self.a.z;
        let v231 = self.b.x*self.c.y*self.a.z;
        let v312 = self.c.x*self.a.y*self.b.z;
        let v132 = self.a.x*self.c.y*self.b.z;
        let v213 = self.b.x*self.a.y*self.c.z;
        let v123 = self.a.x*self.b.y*self.c.z;
        return (1.0/6.0)*(-v321 + v231 + v312 - v132 - v213 + v123);
    }
}

fn main() {
    let model:f64 = fs::read_to_string("MODEL.CSV").expect("Failed to open model")
        .split("\n").take(25910)
        .map(|s| Triangle::from_str(s).unwrap().volume())
        .sum();
    println!("{:.3}",model /1000.0);
}
