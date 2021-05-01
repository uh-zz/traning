puts 1 + 2

a = 'hello'
puts a

b = 'ほげ'
puts b

# オブジェクトなのでこんなことができる。。
puts true.to_s
puts nil.to_s

# 変数展開
puts "#{a}, world"

# 数値の区切りができる。。
puts 1_000_000_000

# 小数使いたいときはどちらかに0をつける
puts 1.0 / 2

# 丸め誤差の対処
# これだとうまく行かない
puts 0.1 * 3.0 == 0.3

# 小数の後にrをつける
# rはRational(有理数)クラスの数値にするためのシンタックス
puts 0.1r * 3.0r == 0.3

# 普通の小数に戻したいとき
puts (0.1r * 3.0r).to_f


# インクリメントの仕方
# ++は使えない
# n++
n = 0
n += 1
puts n


# rubyのifは最後に評価された式を戻り値として返す。。
# ええ。。
res =
    if n > 0
        "hoge"
    else
        "fuga"
    end

puts res

# 後置ifというテクニックもある
n *= 5 if n == 1
puts n

# thenを入れると１行にまとまる
# あまりつかわないらしい
# 。。。
if n > 0 then "hgoe"
else "fuga"
end