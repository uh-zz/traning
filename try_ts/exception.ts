(async () => {
    try {
        const fetch = require("node-fetch");

        // URLのTypeErrorの可能性あり
        const res = await fetch(`https://api.ipify.org?format=json`);

        // ここでokチェックしないパターンが多いらしい
        if (res.ok) {

            // jsonが返ってきてないSyntaxErrorの可能性あり
            const json = await res.json();
            console.log(json);
        }
    } catch(e) {
        console.log(`error: ${e.message}`);
    } finally {
        console.log("fetch end")
    }
    
    // const e = new Error("hoge")
    // console.log(`e.name: ${e.name}`)
    // console.log(`e.message: ${e.message}`)
    // console.log(`e.stack: ${e.stack}`)
    
})();

// リカバリー処理分岐のためにユーザ定義の例外クラスを作成
//
//
// 立派な継承ツリーを作る必要はない
//
//
//
// まずは共通クラスを作成
class BaseError extends Error {
    constructor(e?: string) {
        super(e);
        this.name = new.target.name;
    }
}

// BaseErrorを継承した独自クラスを作成
class NetWorkAccessError extends BaseError {
    constructor(public statusCode: number, e?: string) {
        super(e)
    }
}

// 追加属性がなければコンストラクタも定義不要
class NoNetWorkError extends BaseError {}


(async() => {
    try {
        await getUser();
    } catch(e) {
        if (e instanceof NoNetWorkError) {
            console.log(`no network`)
        } else if (e instanceof NetWorkAccessError) {
            if (e.statusCode < 500) {
                console.log(`program has bug`)
            } else {
                console.log(`server error`)
            }
        }
    }
})();

async function getUser() {
    throw new NetWorkAccessError(401, "hoge");
}