import java.io.*;

public class ErrorHandling {
    private static void doSomething() {
        File newdir = new File("newdir");
        try {
            // 作ったディレクトリはfinallyで消す
            newdir.mkdir();

            FileWriter fw = new FileWriter(new File("newdir/newfile"));

            // --------------------------------
            // ここで例外が発生!!
            // --------------------------------
        
            // 最後にFileWriterを閉じる
            fw.close();
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        finally {
            // FileWriterがディレクトリ内でファイルハンドルを握っているので
            // 以下のディレクトリ、ファイル削除は行われない
            for (String n: newdir.list()) {
                new File(newdir.getPath(), n).delete();
            }

            newdir.delete();


            // ディレクトリの作成とファイルの作成に順序が必須であるように、削除も順序がある
            // 1. FileWriterを閉じる
            // 2. ファイルやディクトりを消す

            // つまり開発者はFileWriterの生存範囲をよく考えながらプログラミングしなければいけない
            
            // try-with-resourceを使ったとしてもソースコードがネストされるので手順が多い場合は
            // 可読性が下がる
        }
    }

    public static void main(String[] args) {
        doSomething();
    }
}