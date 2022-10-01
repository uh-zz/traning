// !(P || Q) == !P && !Q

// 以下のようなロジック
// hogeはfugaでもないかつ、nyaoでもない
let hoge = 'hoge';
if (hoge != 'fuga' && hoge != 'nyao') {
    console.log('hoge is hoge?')
}

// ド・モルガンの法則より、次のように言い換える
// （hogeはfugaあるいはnyaoのどちらか）ではない
if (!(hoge == 'fuga' || hoge == 'nyao')) {
    console.log('hoge is hoge?')
}

// !(P && Q) == !P || !Q// 以下のようなロジック

// hogeはfugaではなく、あるいはnyaoでもない
// tscが優秀なので、コンパイルエラーになる
if (hoge != 'fuga' || hoge != 'nyao') {
    console.log('hoge is not hoge')
} else {
    console.log('hoge is hoge?')
}

// ド・モルガンの法則より、次のように言い換える
// （hogeはfugaかつnyao）ではない
//
// tscが優秀なので、コンパイルエラーになる
// 使い道としてはオブジェクトの比較とかかなあと思ったり...
if (!(hoge == 'fuga' && hoge == 'nyao')) {
    console.log('hoge is hoge?')
}