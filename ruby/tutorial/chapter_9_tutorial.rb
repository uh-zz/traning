# rubyの例外処理
begin
# 1 / 0
# 'hoge'.foo
rescue ZeroDivisionError
  puts 'divided by 0'
rescue NoMethodError
  puts 'no method called'
  # すべての例外クラスはExceptionクラスを継承している
  # 通常のプログラムで捕捉するのはStandardErrorクラスかそのサブクラス
  # そのため、むやみにExceptionクラスを使うべきではない
  #   rescue Exception
  #
  #
end

begin
# また例外クラスを複数連ねるときは、サブクラスから書いていく
# superクラスから書いてしまうと、後続の例外でキャッチできないから
# そのため、StandardErrorクラスなどは最後似キャッチするように書く
rescue NoMethodError
  p 'no method'
rescue NameError
  p 'no name'
  # StandardErrorクラスだと、そもそもクラス指定しなくても良いが、rubocopで自動補完される
rescue StandardError
  p 'other exception'
end

# 意図的に例外を発生させる
def hoge
  if 'hoge'.size > 5
  else
    raise ArgumentError, 'invalid hoge'
  end
end

# ruby では rails に例外処理を任せる
# 自分で例外処理みたいなあぶなっかしいことすんな
# ただし、調査のためのresuceくらいはしといてくれ
# それと、例外処理するときは範囲をしぼってね
# それと、例外処理よりは条件分岐で処理したほうが可読性とかパフォーマンスいいからそっち使って
users = %w[hoge fuga]
users.each do |user|
  send_mail_to(user)
rescue StandardError => e
  # 詳細情報だけは出力しとく
  puts "#{e.class}: #{e.message}"
  puts e.backtrace
end

# ファイル操作は、begin/rescue/ensure(finally)でもいけるけど
# ブロック使ったほうがいいよ
File.open('hoge.txt', 'w') do |file|
  file << 'hogee'
  # 例外発生しても自動で閉じてくれるぜ
  # 1/0
end

# 細かく例外分けなくていいなら
def to_date(string)
  # パースできなかったらnilを返す
  # Date.parse(string) rescue nil
end
