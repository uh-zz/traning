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
    
    const e = new Error("hoge")
    console.log(`e.name: ${e.name}`)
    console.log(`e.message: ${e.message}`)
    console.log(`e.stack: ${e.stack}`)
    
})();
