function checkFlag(flag: boolean): string {
    console.log(flag);
    return "check done";
}

const normalize = (input: string): string => {
    return input.toLowerCase();
}

console.log(normalize("Hoge"));

function hello(): void {
    console.log("hello");
}

interface Greeter {
    hello(): void;
}

// arg1 -> string
// arg2 -> 無名関数(arg3: string) => number
// 返り値 -> boolean
let check: (arg1: string, arag2: (arg3: string) => number) => boolean;

function sort(a: string[], conv: (value: string) => string) {

    const entries = a.map((value) => [value, conv(value)]);
    entries.sort((a, b) => {
        if (a[1] > b[1]) {
            return 1;
        } else if (a[1] < b[1]) {
            return -1;
        }
        return 0;
    });
    return entries.map((entry) => entry[0]);
}

const arr: string[] = ["a", "B", "D", "C"];
console.log(sort(arr, s => s.toLowerCase()));
