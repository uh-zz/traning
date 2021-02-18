// 関数宣言の場合は入出力の引数や関数本体の定義時に
// Tがなんらかの型であるかのように利用できる
//
// 次の関数は第一引数の値を、第二引数の数だけ含む配列を返す
function multiply<T>(value: T, n: number): Array<T> {
    const result: Array<T> = [];
    result.length = n;
    result.fill(value);
    return result;
}

const num2 = 2
const str2 = "hoge";

console.log(multiply(num2, 5));
console.log(multiply(str2, 3));

// [ 2, 2, 2, 2, 2 ]
// [ 'hoge', 'hoge', 'hoge' ]

// 型を明示的に指定したり
console.log(multiply<number>(-1, 5));

// 型推論も可能！！！
console.log(multiply("mazika", 10));
