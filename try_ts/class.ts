// クラスについて
//
//
//
// Date型
//
//
// ミリ秒単位のエポック秒を取得
// UNIX時刻 = UNIX秒 = エポック時刻 = エポック秒
//
// UTC時刻が返される
const v = new Date;
console.log(v)

// 時刻をシンプルに扱うための指針
//
// クライアントで時刻を取得してサーバに送信するときは、Date.now()などで
// エポック時刻に変換してから送信する（タイムゾーンなし）
//
// サーバでは常にエポック時刻で扱う
// ミリ秒単位かマイクロ秒単位かはルールで決める
//
// サーバからフロントに送った段階で、Date.now()を使ってローカル時刻化する

// 明示的にUTC時刻のDateインスタンスを生成
const d = new Date(Date.UTC(2020, 8, 21, 11, 10, 5))
console.log(d)

// 同一日時の判定方法
//
// 0時0分0秒のエポック秒
const today = (new Date()).setHours(0,0,0,0);

const someDate: Date = new Date()

const isSameDay = (new Date(someDate)).setHours(0,0,0,0) === today;
console.log(isSameDay)

// RegExp型
// 
//
const input = "03-1234-5678";

if (input.match(/\d{2,3}-\d{3,4}-\d{4}/)) {
    console.log("電話番号です");
}

// (パターン)を(?<名前>パターン)と書くことができる
const match = "01-2345-6789".match(/^(?<AC>0\d{1})-(?<MA>\d{4})-(?<SA>\d{4})$/);

// anyにしてるけど正しくなさそう
const { AC, MA, SA }: any = match != null ? match.groups : null;
console.log(AC, MA, SA);

// JSON
//
//
//
type Person3 = { name: string }
let person3: Person3;
// JSONをパースするときはtryで囲うとよい
try{
    person3 = JSON.parse(input)
} catch(e: unknown) {
    // fallback
    // SyntaxErrorの場合は仮の値を入れる
    person3 = { name: "hoge" }
}

// JSONは作成されたのではなく、発見された


// URL
//
//
//
const u = new URL("https://developer.mozilla.org/en-US/docs/Web/API/URL#Methods");


// searchParams: URLSearchParams { 'search' => 'url search params' }
// クエリーパラメーターを解析してくれる
const c = new URL("https://caniuse.com/?search=url%20search%20params")
console.log(c)