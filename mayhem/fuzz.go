package fuzz

import "strconv"
import "github.com/88250/gulu"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            var test gulu.GuluFile
            test.IsValidFilename(content)
            return 0

        case 1:
            content := string(bytes)
            var test gulu.GuluGo
            test.IsAPI(content)
            return 0

        case 2:
            content := string(bytes)
            var test gulu.GuluGo
            test.GetExecutableInGOBIN(content)
            return 0

        case 3:
            var i interface{}
            var test gulu.GuluJSON
            test.UnmarshalJSON(bytes, i)
            return 0

        case 4:
            content := string(bytes)
            var test gulu.GuluRune
            test.ContainChinese(content)
            return 0

        case 5:
            var test gulu.GuluStr
            test.FromBytes(bytes)
            return 0

        default:
            content := string(bytes)
            var test gulu.GuluStr
            test.ToBytes(content)
            return 0
    
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}