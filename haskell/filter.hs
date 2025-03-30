import Prelude hiding (filter)
filter :: (a -> Bool) -> [a] -> [a]
filter _ [] = []
filter p (x:xs)
  | p x       = x : filter p xs
  | otherwise = filter p xs
-- example
-- filter even [1..10]
-- [2,4,6,8,10]
