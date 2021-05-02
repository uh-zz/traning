numbers = [1, 2, 3, 4]
sum = 0
numbers.each do |n|
  sum += n
end

p sum

a = [1, 2, 3, 1, 2, 3]

# 奇数削除の方法
# わかりやすいけどrubocopさんに注意される
#
# a.delete_if do |n|
#   n.odd?
# end

# こう書いてほしいみたい
# まあまあまあ。。。
a.delete_if(&:odd?)

p a

# 一行でブロックを書きたいときは{}をつかう
numbers.each { |n| sum += n }
p sum

# ブロックを使うメソッド

# map
# 新しい配列を返してくれる
new_numbers = numbers.map { |n| n * 10 }
p new_numbers

# select
# フィルタリングした配列を返す
#
# 例によってこちらも注意
# even_numbers = numbers.select { |n| n.even? }
even_numbers = numbers.select(&:even?)
p even_numbers

# find
# 前方一致の最初の要素を返す

# inject/reduce
# 引数を２つとり、初回のみinject/reduceの引数(例: 0)が入る
# ２回目以降は前回のブロックの戻り値が入る
# javascriptでも使うので説明ありがたい
sum = numbers.reduce(0) { |result, n| result + n }
p sum

# こんなトリッキーなこともできる
p %w[mon tue wed thu fri sat].reduce('sun') { |result, s| "#{result}, #{s}" }
