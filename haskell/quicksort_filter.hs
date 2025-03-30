import Prelude hiding (filter)
filter :: (a -> Bool) -> [a] -> [a]
filter _ [] = []
filter p (x:xs)
  | p x       = x : filter p xs
  | otherwise = filter p xs

quicksort :: (Ord a) => [a] -> [a]
quicksort [] = []
quicksort (x:xs) =
    let smallerOrEqual = (<= x) `filter` xs
        larger = (> x) `filter` xs
    in quicksort smallerOrEqual ++ [x] ++ quicksort larger
