import Prelude hiding (map)
map :: (a -> b) -> [a] -> [b]
map _ [] = []
map f (x:xs) = f x : map f xs
-- example
-- map (+1) [1,2,3,4]
-- [2,3,4,5]
-- map (map (^2)) [[1,2],[3,4,5],[7,8]]
-- [[1,4],[9,16,25],[49,64]]
