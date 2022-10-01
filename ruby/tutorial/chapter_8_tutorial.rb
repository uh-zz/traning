# モジュールからはインスタンスを生成できない
# 継承はできない
module Loggable
  def log(text)
    p "[LOG] #{text}"
  end

  # 特異メソッドとして定義してあげれば、直接呼べる
  def self.singular_log(text)
    p "[ABNORMAL] #{text}"
  end

  # これだとミックスインと特異メソッドのどちらでも使える。。。
  module_function :log
end

class Product
  # utilクラス的な扱い
  # インスタンスメソッドとして使うならこっち
  include Loggable

  # クラスメソッドとして使うならこっち
  # extend Loggable

  def title
    log 'title is called'
    'great movie'
  end
end

class User
  include Loggable

  def name
    log 'name is called'
    'alice'
  end
end

product = Product.new
product.title

# 必要なければprivateにしておく
# product.log 'public?'

# クラスメソッドであればこんなふうに呼べたり
# Product.log 'hoge'

user = User.new
user.name

# ひえっ。。
s = 'hoge'
# 特異メソッド
s.extend(Loggable)
# s.log 'fugafuga'

# モジュールを利用すると名前空間をつくれる
module BaseBall
  class Second
    attr_accessor :player

    def initialize(name)
      @player = name
    end
  end
end

# ネストさせずにこういう書き方もできる
# こっちのがスマート
class BaseBall::Third
  attr_accessor :player

  def initialize(name)
    @player = name
  end
end

module Clock
  class Second
    attr_accessor :digits

    def initialize(digits)
      @digits = digits
    end
  end
end

# 同名のクラス名だけどおｋ
hoge = BaseBall::Second.new('pipipi')
fuga = Clock::Second.new(5)

p hoge.player
p fuga.digits

third = BaseBall::Third.new('pupupu')
p third.player

# モジュールだけで呼べるんかい！
# 特異メソッドに限る
Loggable.singular_log('hogehoge')

# module_functionを使うと特異メソッド的にもいける
# こういうのをmodule関数という
Loggable.log('hogehoge')
