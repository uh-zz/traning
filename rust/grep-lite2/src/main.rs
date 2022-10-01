fn main() {
    let ctx_lines = 2;
    let needle = "oo";
    let haystack = "\
Every face, every shop,
bedroom window, public-house, and
dark square is a picture
feverishly turned--in search of what?
It is the same with books.
What do we seek
through millions of pages?";

    // マッチした行番号を入れる
    let mut tags: Vec<usize> = vec![];

    // コンテクスト行を格納したマッチごとのベクターを入れる
    let mut ctx: Vec<Vec<(usize, String)>> = vec![];

    // マッチした行番号を記録
    for (i, line) in haystack.lines().enumerate() {
        if line.contains(needle) {
            tags.push(i);

            // Vec::with_capacity(n)で、n個の要素のための空間を予約
            // (前後)*指定した行数+対象の行
            let v = Vec::with_capacity(2 * ctx_lines + 1);
            ctx.push(v);
        }
    }

    // マッチ行がない
    if tags.is_empty() {
        return;
    }

    for (i, line) in haystack.lines().enumerate() {
        for (j, tag) in tags.iter().enumerate() {
            // saturating_sub は、減算用のメソッドで、整数のアンダーフローが発生するとき、
            // プログラムをクラッシュする代わりに0を返す
            let lower_bound = tag.saturating_sub(ctx_lines);

            let upper_bound = tag + ctx_lines;

            if (i >= lower_bound) && (i <= upper_bound) {
                let line_as_string = String::from(line);
                let local_ctx = (i, line_as_string);
                ctx[j].push(local_ctx);
            }
        }
    }
    for local_ctx in ctx.iter() {
        for &(i, ref line) in local_ctx.iter() {
            let line_num = i + 1;
            println!("{}: {}", line_num, line);
        }
    }
}
