use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    let f = File::open("readme.md").unwrap();
    let mut reader = BufReader::new(f);

    let mut line = String::new();

    loop {
        let len = reader.read_line(&mut line).unwrap();

        if len == 0 {
            break;
        }

        println!("{} ({} bytes long)", line, len);

        // Stringの長さを0に縮めて、前行の内容が残らないようにする
        line.truncate(0);
    }
}
