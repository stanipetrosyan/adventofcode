import Data.Char ( isDigit )

type Input = [String]

parse :: String -> Input
parse = lines

extractNumber :: String -> String
extractNumber s = [x | x <- s, isDigit x]

stringToInt :: String -> Int
stringToInt s = read (head s : [last s])


calibrationValue :: String -> Int
calibrationValue = stringToInt . extractNumber

replaceDigit :: String -> String
replaceDigit ('o':'n':'e':xs) = '1':replaceDigit ('e': xs)
replaceDigit ('t':'w':'o':xs) = '2':replaceDigit ('o': xs)
replaceDigit ('t':'h':'r':'e':'e':xs) = '3':replaceDigit ('e':xs)
replaceDigit ('f':'o':'u':'r':xs) ='4':replaceDigit ('r':xs)
replaceDigit ('f':'i':'v':'e':xs) = '5':replaceDigit ('e':xs)
replaceDigit ('s':'i':'x':xs) = '6':replaceDigit ('x': xs)
replaceDigit ('s':'e':'v':'e':'n':xs) = '7':replaceDigit ('n':xs)
replaceDigit ('e':'i':'g':'h':'t':xs) = '8':replaceDigit ('t':xs)
replaceDigit ('n':'i':'n':'e':xs) = '9':replaceDigit ('e':xs)
replaceDigit (x:xs) = x:replaceDigit xs
replaceDigit "" = ""


solve1 :: Input -> Int
solve1 = sum . map calibrationValue

solve2 :: Input -> Int 
solve2 = sum . map (calibrationValue . replaceDigit)


main :: IO ()
main = do 
    contents <- readFile "input.txt"
    let input = parse contents
    
    print (solve1 input)
    print (solve2 input)
