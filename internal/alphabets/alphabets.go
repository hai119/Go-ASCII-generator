package alphabets

// 不同语言的字符集定义
var (
    GENERAL = map[string]string{
        "standard": "@%#*+=-:. ",
        "complex":  "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ",
    }

    ENGLISH = map[string]string{
        "standard": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
    }

    CHINESE = map[string]string{
        "standard": "永和九年，岁在癸丑。暮春之初，会于会稽山阴之兰亭，修禊事也。",
    }

    JAPANESE = map[string]string{
        "standard": "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん",
    }

    KOREAN = map[string]string{
        "standard": "ㄱㄴㄷㄹㅁㅂㅅㅇㅈㅊㅋㅌㅍㅎㅏㅑㅓㅕㅗㅛㅜㅠㅡㅣ",
    }
) 