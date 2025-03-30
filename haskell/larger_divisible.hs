import Prelude hiding (filter)

filter :: (a -> Bool) -> [a] -> [a]
filter _ [] = []
filter p (x : xs)
  | p x = x : filter p xs
  | otherwise = filter p xs

largerDivisible :: Integer
largerDivisible = head (filter p [1000000, 999999 ..])
  where
    p x = x `mod` 3829 == 0
