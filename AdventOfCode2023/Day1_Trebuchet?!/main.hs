import Data.Char

type Input = [String]

parse :: String -> Input
parse = lines

extractNumber :: String -> String
extractNumber s = [x | x <- s, isDigit x]

stringToInt :: String -> Int
stringToInt s = read (head s : [last s])


calibrationValue :: String -> Int
calibrationValue = stringToInt . extractNumber


solve1 :: Input -> Int
solve1 = sum . map calibrationValue


main :: IO ()
main = do 
    contents <- readFile "input.txt"
    let input = parse contents

    print (solve1 input)
