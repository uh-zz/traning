fn greet_world() {
    println!("Hello, world!");
    let southern_germany = "Gr ̈ußGott!";
    let japan = "ハローワールド";
    let regions = [southern_germany, japan];
    for region in regions.iter() {
        println!("{}", &region);
    }
}

fn main() {
    greet_world()
}
