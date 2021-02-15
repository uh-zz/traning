console.log("タイマー呼び出し前");
setTimeout(() => {
  console.log("時間が来た");
}, 1000);
console.log("タイマー呼び出し後");

// JavaScriptでは時間のかかる処理を実行する場合
// 完了した後に呼び出す処理を処理系に渡すことはあっても
// そこで処理を止めることはしない
// 
// sleepしないってこと

// 昔JavaScriptでは時間のかかる関数には
// 引数にコールバック関数を設定し、
// 時間のかかる処理の後に実行する関数を表現した

// Promiseがない時代のコールバック地獄
//
//
//　data1を読み込んだら、data2...と階層的に読んでいく
// const fs = require('fs');
// fs.readFile('data1.txt', function(data1) {
//      fs.readFile('data2.txt', function(data2) {
//          fs.readFile('data3.txt', function(data3) {
//              fs.readFile('data4.txt', function(data4) {
//                  fs.readFile('data5.txt', function(data5) {
//                      console.log(data1 + data2 + data3 + data4 + data5);
//                  })
//              })
//          })
//      })
//  })

// awaitはPromiseのシュガーシンタックス
// then()を呼び出した引数で渡される値が関数の返り値になる
//
// 新: 非同期処理をawaitで待つ（ただし、awaitはasync関数の中でのみ有効）
// const resp = await fetch(url);
// const json = await resp.json();
// console.log(json);

// asyncはPromiseを返す
// setTimeoutは最初がコールバックという変態仕様なので仕方なくnew Promise
const sleep = async (time: number): Promise<number> => {
    return new Promise<number>(resolve => {
      setTimeout(()=> {
        resolve(time);
      }, time);
    });
};

async function Asynchoge(): Promise<number> {
    await sleep(100);    
    return 10;
}

// for of, if, while, switchはawaitとの相性も良い
// ただし、ループ内のawaitには注意
//
//　一つずつ処理の完了を待つのでボトルネックになりがち
// for (const value of iterable) {
//     await doSomething(value);
// }
// console.log("この行は全部のループが終わったら実行される");

// このawaitでは待たずにループが終わってしまう
//
// このコード書いてたや、、直すか
// iterable.forEach(async value => {
//     await doSomething(value);
// });
// console.log("この行はループ内の各処理が回る前に即座に実行される");

// [TODO] 時間ある時に実装しよう
// async function 味噌汁温め(): Promise<味噌汁> {
//     await ガスレンジ();
//     return new 味噌汁();
// }
  
// async function ご飯温め(): Promise<ご飯> {
//     await 電子レンジ();
//     return new ご飯();
// }
  
// 両方の操作を非同期で行い、どちらの完了も待つ
// const [a味噌汁, aご飯] = await Promise.all([味噌汁温め(), ご飯温め()]);
// いただきます(a味噌汁, aご飯);