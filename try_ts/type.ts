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

const a: readonly number[] = [1, 2, 3];

// a[0] =1
