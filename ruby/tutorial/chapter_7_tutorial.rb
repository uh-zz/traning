class Product
  # rubyでは配列はミュータブルなオブジェクトなので、
  # 定数として持っていても破壊的に変更しかねない
  # SOME_NAMES = %i[Foo Bar Baz]
  #
  # 以下のようにfreezeすることでほんとに定数として扱える（ほんとにとは
  SOME_NAMES = %i[Foo Bar Baz].freeze
  def self.names_without_foo(names = SOME_NAMES)
    # 定数なのでここで削除できないようになってる
    names.delete(:Foo)
    names
  end
end

# freezeしてるのでこれはエラー
# p Product.names_without_foo

# ただし、配列の要素が「文字列」の場合は破壊的に変更できる！
# シンボルだと非破壊的！（うれぴ
# p Product::SOME_NAMES[0].upcase!
# p Product::SOME_NAMES

# 特異メソッド
#
# rubyのapiドキュメントでは、クラスメソッドという用語は登場せず、
# 特異メソッドとして説明されている。。。
#
alice = "i'm alice"
class << alice
  def shuffle
    chars.shuffle.join
  end
end

# これでも行ける
# Goのレシーバに似てるが、、
# def alice.shuffle
#     chars.shuffle.join
# end
p alice.shuffle
