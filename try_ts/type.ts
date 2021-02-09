let birthYear: number | '平成' | '昭和';

birthYear = 1989

birthYear = '平成'

// birthYear = '令和'
// birthYear = null

const address = `hoge
fuga`

console.log(address)

const smalls = [
    "小動物",
    "小型車",
    "小論文"
]

// カンマを入れることで要素を省略している
// カンマを１ついれると、"小動物"がスキップされて、otherには残り２つが入る
// 残余構文(Rest)
const [, ...other] = smalls;

console.log(other);

// 存在チェック
if (smalls.includes("小動物")) {
    console.log("べんべんべんり")
}

const numbers = [30, 1, 200]
console.log(numbers.sort((a,b) => a - b))

const stations = [
    {name: "池袋", users: 558623},
    {name: "新宿", users: 775386},
    {name: "渋谷", users: 366128},
    {name: "東京", users: 462589}
];

// 駅の利用者数でソート
console.log(
    stations.sort((a, b) => b.users - a.users) 
);

const stations2 = [
    {name: "大手町", lines: 5, yomi: "おおてまち"},
    {name: "飯田橋", lines: 7, yomi: "いいだばし"},
    {name: "永田町", lines: 5, yomi: "ながたちょう"},
];


function cmpNum(a: number, b: number) {
    return (a < b) ? -1 : (a === b) ? 0 : 1;
}

function cmpStr(a: string, b: string) {
    return (a < b) ? -1 : (a === b) ? 0 : 1;
}

console.log(
    stations2.sort((a,b) => {
        const lineScore = cmpNum(a.lines, b.lines);
        const yomiScore = cmpStr(a.yomi, b.yomi);

        return lineScore * 10 + yomiScore
    })
)

// 元のオブジェクトを破壊したくない場合
const nodestroy = [...stations2].sort((a,b) => b.lines - a.lines);

nodestroy.forEach((station, index) => {
    console.log(index, station);
})

// 読み込み専用のreadonlyは型名の前につける
const a: readonly number[] = [1, 2, 3];

// a[0] =1

// anyは型チェックを放棄
function someFunction(opt: any) {
    console.log(opt.debug);
}


type BirthYear = number | string;

const birthday: BirthYear = "平成";
console.log(birthday);

type FoodMenu = "北極" | "冷やし味噌";
const myOrder: FoodMenu = "北極";
orderFood(myOrder);

function orderFood(food: FoodMenu) {
    console.log(food);    
}

// 構造体やん
type Person = {
    name: string;
    favoriteBank: string;
    favoriteGyudon: string;
}

// インターフェースを指定　←わかりよい
const person: Person = {
    name: "hoge",
    favoriteBank: "mizuho",
    favoriteGyudon: "matsuya"
}

console.log(person)

// readonlyを使うと読み込み専用
// 変数名の後ろに?をつけるとオプションにできる
type Person2 = {
    name: string;
    readonly favoriteBank: string;
    favoriteGyudon?: string;
}

let person2: Person2 = {
    name: "fuga",
    favoriteBank: "smbc",
}

person2.name = "hoge";

// これがエラーになる
// person2.favoriteBank = "mizuho";

// Partialをつけると全ての要素の設定が必要なくなる
const wzz: Partial<Person> = { name: "hoge"};

// mapを作る
let postalCodes: { [key: string]: string } = {
    "602-0000": "京都市上京区",
    "602-0827": "京都市上京区相生町",
    "602-0828": "京都市上京区愛染寺町",
    "602-0054": "京都市上京区飛鳥井町",
};

postalCodes["601-0054"] = "京都市上京区飛鳥井町";

console.log(postalCodes);

for (const key of Object.keys(postalCodes)) {
    
}

for (const value of Object.values(postalCodes)) {
    
}

// key と値両方とりだす
for (const [key, value] of Object.entries(postalCodes)) {
    console.log(key, value)
}


type Twitter = {
    twitterId: string;
}

type Instagram = {
    instagramId: string;
}
  
// 型合成
const sns: Twitter & Instagram = {
    twitterId: "hoge",
    instagramId: "fuga"
}

// 型合成（インターフェース）
interface PartyPeople extends Twitter, Instagram {
}

const sns2: PartyPeople = {
    twitterId: "hoge",
    instagramId: "fuga"
}

// userNameOrIdは文字列か数値
let userNameOrId: string | number = getUser();

// 型ガードにはtypeofを使う
if (typeof userNameOrId === "string") {
    // このif文の中では、userNameOrIdは文字列型として扱われる
    this.setState({
        userName: userNameOrId.toUpperCase()
    });
} else {
    // このif文の中では、userNameOrIdは数値型として扱われる
    const user = this.repository.findUserByID(userNameOrId);
    this.setState({
        userName: user.getName()
    });
}

