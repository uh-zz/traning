# yieldを使うと、メソッド呼び出し時にブロックの処理を差し込めるぜ！
# ちょっと玄人感ある
def greeting
  p 'おっはー'
  yield
  p 'おやすー'
end

greeting do
  p 'こんちはー'
end

# yieldあるメソッド呼び出しでブロック使わないとエラー
# greeting

# ブロックなしでもできるように
def hoge
  p 1
  yield if block_given?
  p 3
end

hoge do
  p 2
end

hoge

# yieldに引数も渡せる
def fuga
  p 'おっはー'
  yield 'hogeee'
  p 'おやすー'
end

# ブロックでyieldの引数を受け取る
fuga do |text|
  p "#{text} is fun"
end

def shouted(text)
  "the cat suddenly shouted. 「 #{text * 3} 」 "
end

# こんなこととか
fuga do |text|
  p shouted(text)
end

# これはおもしろいかも
# 抽象メソッド的な感じ
# &blockで明示的にブロックを受け取る
def greeting_ja(&block)
  texts = %w[おっはー こんちはー おやすー]
  greeting_common(texts, &block)
end

def greeting_en(&block)
  texts = ['good morning', 'hello', 'good night']
  greeting_common(texts, &block)
end

def greeting_common(texts, &block)
  p texts[0]
  # この要素だけブロックの引数として渡す
  block.call(texts[1])
  p texts[2]
end

greeting_ja do |text|
  p shouted(text)
end

greeting_en do |text|
  p "#{text}, japanese say ういー"
end

# Procはブロックをオブジェクト化するためのクラス
# lambda的な
add_proc = proc { |a, b| a + b }
p add_proc.call(1, 2)

# なるほど、Procっていうブロックを関数とみたてて&blockで渡してるんだ
# こんなことができる。。玄人っぽい
repeat_proc = proc { |text| p text.upcase * 2 }
greeting_en(&repeat_proc)

# もうラムダって紹介されてた
add_lambda = ->(a, b) { a + b }
p add_lambda.call(1, 2)

# procとlambdaの違い
# lambdaのほうが引数チェックが厳密
# returnとbreakの挙動がちがう

# おしゃれなやつ
reverse_proc = proc { |s| s.reverse }
p %w[ruby java python].map(&reverse_proc)

# クロージャとしてprocをつかう
def generate_proc(array)
  counter = 0

  proc do
    counter += 10
    array << counter
  end
end

values = []

# 宣言時にローカル変数と引数を保存する
sample_proc = generate_proc(values)
p values
sample_proc.call
p values
sample_proc.call
p values
