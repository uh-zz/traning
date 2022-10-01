use std::ops::Add;
use std::time::Duration;

// FYI: https://github.com/rust-in-action/code
fn main() {
    // rustの変数はデフォルトで非可変→read-only
    let a = 10;
    println!("{}", a);

    // rustの参照
    let b = 42;
    let c = &b;

    let d = b + *c;
    println!("{}", d);

    let needle = 42;
    let haystack = [1, 1, 2, 5, 15, 42, 203, 877, 4140, 21147];

    for item in &haystack {
        if *item == needle {
            println!("{}", *item);
        }
    }

    greet_world();

    println!("{}", add(10, 20));
    println!("{}", add(1.2, 3.4));

    println!("{:?}", add(Duration::new(5, 0), Duration::new(10, 0)));
}

fn greet_world() {
    println!("Hello, world!");
    let southern_germany = "Gr ̈ußGott!";
    let japan = "ハローワールド";
    let regions = [southern_germany, japan];
    for region in regions.iter() {
        println!("{}", &region);
    }
}

// rustでのジェネリクス
// ジェネリクスを使用する場合、トレイト境界を入れる必要がある
// この例では、「Tは、std::ops::Add を実装していなければならない」という意味
// トレイトという言語機能は、インターフェースに例えられる
// つまり、複数の型に同じトレイトを持たせることで、それらの型が共通の振る舞いをもつことができる
fn add<T: Add<Output = T>>(i: T, j: T) -> T {
    i + j
}
