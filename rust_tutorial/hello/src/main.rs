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
    // rustの変数はデフォルトで非可変→read-only
    let a = 10;
    println!("{}", a);

    greet_world()
}
