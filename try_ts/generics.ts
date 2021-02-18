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

// 型パラメータに制約をつけたい場合
//
//
// 引数のジェネリクスの型を満たすようなインターフェースを用意
interface Person4 {
    getBirthDay(): Date;
}

function isTodayBirthday<T extends Person4>(person: T): boolean {
    const today = new Date();

    const birthDay = person.getBirthDay();

    return today.getMonth() === birthDay.getMonth() &&
        today.getDate() === birthDay.getDate();
}

// keyofの使い方
//
// TypeScript 2.1で追加された
//
// オブジェクトと key を受け取り、値を返す関数
// 
// このようなインターフェースを用意
// interface ISomeObject {
//     firstKey: string;
//     seocondKey: string;
// }
//
// こうすると暗黙的にany型が返ってくるのでコンパイラでエラーになる
// function getProperty(obj: ISomeObject, key: string) {
//     return obj[key];
// }
//
//
// 解決法1
// インターフェースに keyの型と返り値の型を定義する
// interface ISomeObject {
//     firstKey: string;
//     seocondKey: string;
//     [key: string]: string; // このような定義を追加
// }
//
//
// 解決法2
// keyof を使う
interface ISomeObject {
    firstKey: string;
    seocondKey: string;
    2: string;
}

// keyof type/interfaceという構文にすることで、
// オブジェクトのkey, value情報を取得してくれる
// 
function getProperty(obj: ISomeObject, key: keyof ISomeObject) {
    return obj[key];
}

// ジェネリクスで書くならこんな感じ
function getPropertyGen<S, T extends keyof S>(obj: S, key: T) {
        return obj[key];
}