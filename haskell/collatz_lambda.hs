chain :: Integer -> [Integer]
chain 1 = [1]
chain n
    | even n    = n : chain (n `div` 2)
    | odd n    = n : chain (3 * n + 1)

-- example
-- 1から100までの数のうち、長さ15以上のコラッツ列の開始数になる数はいくつか
numLongChains :: Int
numLongChains = length (filter (\xs -> length xs > 15) (map chain [1..100]))
