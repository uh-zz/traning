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

# 配列の作りかた

# ふむふむ
p (1..5).to_a

# ?!
# トリッキー
p [*1..5]

# rgb変換
# それぞれサンプル見ずにトライ（えらい
#  10 -> 16
def to_hex(r, g, b)
  format('#%#<red>x%#<green>x%#<blue>x', red: r, green: g, blue: b).gsub(/0x/, '')
end

# 16 -> 10
# やりたいことはこっちのがわかりやすい
def to_ints(hex_str)
  hex_str.delete_prefix('#').scan(/.{1,2}/).map(&:hex)
end

# サンプルのコード
# \wで#をフィルタリングできてる（ぐぬぬ
def to_ints_sample(hex)
  hex.scan(/\w\w/).map(&:hex)
end

p to_hex(255, 255, 255)
p to_ints('#ffffff')
p to_ints('#043c78')
p to_ints_sample('#043c78')

# 同じ要素の配列を作る
same_factor = []
same_factor << 1 while same_factor.size < 5
p same_factor

# rubyでもforで書けるが、ローカル変数のスコープがなくなるのでつかわない
# for local_n in numbers
#   sum += local_n
# end
#
# これはエラーじゃない
# p local_n

# よくあるやつ
question = [
  'わたしをたすけてくれ',
  'まおうをたおしてくれ!',
  'きみが「はい」というまでききつづける'
]
question.each do |n|
  print "#{n} => "
  answer = %w[yes no].sample
  p answer

  # はいと答えなければ聞き返す
  redo unless answer == 'yes'
end
