import Data.List.Split ( splitOneOf )
import Data.Char(isSpace)

type Input = [String]

parse :: String -> Input
parse = lines


redValid :: String -> Bool
redValid actual = (read actual::Integer) <= 12

greenValid :: String -> Bool
greenValid actual = (read actual::Integer) <= 13

blueValid :: String -> Bool
blueValid actual = (read actual::Integer) <= 14


gameValid :: [String] -> Bool
gameValid [value, "red"] = redValid value
gameValid [value, "green"] = greenValid value
gameValid [value, "blue"] = blueValid value
gameValid ["Game", value] = True


cutWhitespace :: [String] -> [String]
cutWhitespace = map (dropWhile isSpace)

fixInput :: String -> [String]  
fixInput = cutWhitespace . splitOneOf ";:,"

checkValid :: [String] -> String
checkValid set = if and sub then head set else "0"
    where sub = map (gameValid . splitOneOf " ") set

calcScore :: String -> Int
calcScore ('G':'a':'m':'e':xs)  = read xs
calcScore "0" = 0

main :: IO ()
main = do 
    contents <- readFile "input.txt"
    let input = parse contents

    let score = sum $ map (calcScore . checkValid . fixInput) input
    print score