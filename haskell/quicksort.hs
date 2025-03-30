-- 例えば、リスト[3,1,4,2]でのプロセスを追うと：

-- ピボット3で分割: [1,2]と[4]
-- [1,2]をさらにピボット1で分割: []と[2]
-- [2]をさらにピボット2で分割: []と[]
-- [4]をピボット4で分割: []と[]
-- このようにして、最終的にはすべての再帰呼び出しが空リストに到達し処理が終了します。
-- Haskellの関数型プログラミングの特徴として、明示的な終了条件を記述するのではなく、再帰と基底ケースの設計によって自然に処理が終了するように実装されています。
quicksort :: (Ord a) => [a] -> [a]
quicksort [] = []
quicksort (x : xs) =
  let smallerOrEqual = [a | a <- xs, a <= x]
      larger = [a | a <- xs, a > x]
   in quicksort smallerOrEqual ++ [x] ++ quicksort larger
