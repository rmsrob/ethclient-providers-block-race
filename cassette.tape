Output race.gif
Set Padding 10
Set Theme "MaterialDarker"
Set FontSize 14
Set TypingSpeed 75ms
Set Width 1000
Set Height 350

Sleep 1s
Type@23ms "Call ethclientRace.go with 2 params to return the best provider"
Sleep 3s
Backspace@1ms 64
Sleep 500ms

Type@32ms "perl -ne 'print if 20..37' ethclientRace.go"
Enter
Sleep 5s

Type@1ms "clear"
Enter
Sleep 1s

Type@23ms "go run ethclientRace.go"
Tab
Enter
Sleep 50s
