function doSomething(){
    {
        // {}をスコープとして扱うこともできる！
        const hoge = 'fuga'
        console.log(hoge)
    }
    // スコープの外からは参照できない
    console.log(hoge)
}