applyTwice :: (a -> a) -> a -> a
applyTwice f x = f (f x)

-- example
-- applyTwice (+3) 10
-- applyTwice (++ " HAHA") "HA"
