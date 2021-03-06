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

class SmallAnimal {

    // プロパティ定義はクラス宣言の中で行う
    animalType: string

    // コンストラクタ（省略可能）
    constructor() {
        this.animalType = "ポメラニアン";
    }

    say() {
        console.log(`animalType: ${this.animalType}`)
    }
}

const smallAnimal = new SmallAnimal();
smallAnimal.say();

class SmallDog {

    // このクラスの中からのみ参照可
    // private secretPlace: string

    // typescript固有の書き方
    // コンストラクタにアクセス修飾子をつけると
    // そのままプロパティになる（便利）
    //  
    // ※つけないとエラー
    constructor(private secretPlace: string) {

    }

    // ほる
    dig(): string {
        return this.secretPlace;
    }

    // 埋める
    bury(treasure: string) {
        this.secretPlace = treasure
    }    
}

const miniatureDachshund = new SmallDog("骨とか");

// miniatureDachshund.bury("何か");

// クラスの外からは参照不可
// miniatureDachshund.secretPlace;

console.log(miniatureDachshund.dig());

// クラスの仕組みができる前のJavaScriptでは
// privateを表現するために、メンバー名の前に「_」をつけて
// 明示的にアクセスしないように促した


// staticメンバ
//
//
// Javaと同じ
// インスタンスではなくクラスで１つだけの要素につける
class StaticSample {

    // 静的なプロパティ
    static StaticVariable: number;
    
    variable: number;

    // 静的なメソッド
    static classMethod() {
        
        console.log(this.StaticVariable);

        console.log(StaticSample.StaticVariable);

        // 通常のプロパティは参照不可
        // console.log(this.variable);     
    }

    // 通常のメソッド
    method() {
        
        // thisでの静的プロパティは参照不可
        // console.log(this.StaticVariable);

        // クラス名での静的プロパティは参照可
        console.log(StaticSample.StaticVariable);

        // 通常のプロパティは参照可
        console.log(this.variable);     
    }
}

// 静的メソッドが便利そうな唯一の使い道
//
// 多用することはない
//
//
// ファクトリーメソッドの実装
// 
// [TODO] メリット調べる
// インスタンスを生成する専用のメソッドだから？
// 使うときにnewを明示的にしなくていいから？
class Point {

    constructor(public x: number, public y: number) {

    }

    // 極座標を作成するファクトリーメソッド
    static polar(length: number, angle: number) {
        return new Point(
            length * Math.cos(angle),
            length * Math.sin(angle)
        );
    } 
}

console.log(new Point(10, 20));
console.log(Point.polar(10, Math.PI * 0.25));


// thisについて調べたこと
// 
const whitney = {

    sing() {
        console.log("and i will always love you");
    },

    // setTimeout経由でfunctionに定義したthis.sing()は呼べない
    //
    //
    // thisの参照先が、whitneyなのかwindowなのかぶれる
    // start() {
    //     setTimeout(function() {
    //         this.sing()
    //     }, 3000)
    // }

    // 実行コンテキストをコントロールする方法
    //
    // 変数selfを用意してthisを外出しにする
    //
    // thisは実行時のコンテキストに依存してしまう
    //
    //
    // setTimeoutの中ではグローバルにコンテキストが移るので、
    // whitneyを指すthisではなくなる
    //
    // 一時的な変数selfをアッパースコープで定義して、
    // self経由でwhitneyを呼び出す
    //
    //
    // この実装方法は、JavaScriptの仕様である
    // 内側から外側のスコープへと順番に変数が定義されているか探す仕組み
    // →　スコープチェーンを利用したもの
    //
    // アッパースコープを用意しない場合、
    // thisは最上部のwindowでそれ以上ないのでwindowがthisに入る
    // start() {
    //     const self = this;
    //     setTimeout(function() {
    //         self.sing();
    //     }, 3000);
    // } 

    // bind()はES5で追加された
    // 関数が呼ばれる時に依存するコンテキストを指定するための関数
    // start() {
    //     setTimeout(function() {
    //         this.sing();        
    //     }.bind(this), 3000);
    // }

    // アロー関数はES6で追加された
    // アロー関数ではthis宣言時のスコープ(レキシカルスコープ)のthisを参照する
    // 
    // 宣言時とはつまり、whitneyの宣言時
    start() {
        setTimeout(() => {
            this.sing()
        }, 30)
    }
}
whitney.start()

//　typescriptでの書き方
//
//
// class SmallAnimal {
//
    // プロパティを作成(このプロパティをインスタンスクラスフィールドとよぶ)
    //
    // インスタンスクラスフィールドを使うことで、クラス宣言時にプロパティ宣言ができ
    // インスタンス化される時に設定される
//     fav = "小田原";
//
//
    // イベントハンドラなどにメソッドを渡すとき、
    // 以前はthisごとbindして渡していたが、
    // アロー関数を使うことで、インスタンスクラスフィールド含め
    // bindなしで問題なく渡せる    
//     say = () => {
//       console.log(`私は${this.fav}が好きです`);
//     };
// }
//
// 以前はconstructorでbindしていた
// class SmallAnimal {
//     constructor() {
//       this._fav = "小春日";
//       this.say = this.say.bind(this);
//     }
  
//     say() {
//       console.log(`私は${this._fav}が好きです`);
//     };
//   }

// 継承
//
//
class SmallAnimal2 {
    eat() {
        console.log("hogehoge");
    }
}

class Pomeranian extends SmallAnimal2 {
    eat() {
        console.log("i'm pomeranian");
    }
}

const pome = new Pomeranian();
pome.eat()

// インターフェース
//
//
interface Animal {
    eat(): void;
}

class Pomeranian2 implements Animal {
    eat() {}
}

// 未実装だとエラー
// class Pomeranian3 implements Animal {
// }

// デコレータ
//
//
// コンパイル通らず、、
// とりあえず使う時になったら改めて調査する
// function StrongZero(target: any) {
//     target.prototype.drink = function () {
//         console.log('i drank strongzero');
//     };
//     return target;
// }

// @StrongZero
// class Dranker {

// }

// const sa = new Dranker();
// sa.drink();