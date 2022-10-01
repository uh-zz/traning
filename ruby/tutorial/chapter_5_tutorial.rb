# def convert_hash(source, from, to)
#   length_units = { 'm' => 1.0, 'ft' => 3.28, 'in' => 39.37 }
#   (source / length_units[from]) * length_units[to]
# end

# p convert_hash(35_000, 'ft', 'm')

# シンボルを使うようにする
# シンボルをキーにするときは=>を使わずに、シンボル: 値という書き方にする
#
# 定数は外で保持するようにする
LENGTH_UNITS = { m: 1.0, ft: 3.28, in: 39.37 }
def convert_hash(source, from: :from, to: :to)
  ((source / LENGTH_UNITS[from]) * LENGTH_UNITS[to]).round(2)
end

p convert_hash(35_000, from: :ft, to: :m)

# &.
# lonely operator
# ぼっち演算子
a = 'foo'
p a&.upcase
a = nil
# nilの場合はnilを返す
p a&.upcase

# 真偽値の型変換
#
# これが
# def user_exists?
#     user = find_user
#     if user
#         true
#     else
#         false
#     end
# end

# こう
# rubyがfalseかnil以外は全て真という特性を持ってるから
# find_userの戻り値がnilのときは!!nil -> !true -> falseというように変換できる
# def user_exists?
#     !!find_user
# end
