import re


def get_char(s):
    if len(s) != 1:
        raise ValueError('unexpected string "{}"'.format(s))
    # num
    num = [
        "\u0030-\u0039",
        "\u00b2-\u00b3",
        "\u00b9",
        "\u00bc-\u00be",
        "\u0660-\u0669",
        "\u06f0-\u06f9",
        "\u07c0-\u07c9",
        "\u0966-\u096f",
        "\u09e6-\u09ef",
        "\u09f4-\u09f9",
        "\u0a66-\u0a6f",
        "\u0ae6-\u0aef",
        "\u0b66-\u0b6f",
        "\u0b72-\u0b77",
        "\u0be6-\u0bf2",
        "\u0c66-\u0c6f",
        "\u0c78-\u0c7e",
        "\u0ce6-\u0cef",
        "\u0d58-\u0d5e",
        "\u0d66-\u0d78",
        "\u0de6-\u0def",
        "\u0e50-\u0e59",
        "\u0ed0-\u0ed9",
        "\u0f20-\u0f33",
        "\u1040-\u1049",
        "\u1090-\u1099",
        "\u1369-\u137c",
        "\u16ee-\u16f0",
        "\u17e0-\u17e9",
        "\u17f0-\u17f9",
        "\u1810-\u1819",
        "\u1946-\u194f",
        "\u19d0-\u19da",
        "\u1a80-\u1a89",
        "\u1a90-\u1a99",
        "\u1b50-\u1b59",
        "\u1bb0-\u1bb9",
        "\u1c40-\u1c49",
        "\u1c50-\u1c59",
        "\u2070",
        "\u2074-\u2079",
        "\u2080-\u2089",
        "\u2150-\u2182",
        "\u2185-\u2189",
        "\u2460-\u249b",
        "\u24ea-\u24ff",
        "\u2776-\u2793",
        "\u2cfd",
        "\u3007",
        "\u3021-\u3029",
        "\u3038-\u303a",
        "\u3192-\u3195",
        "\u3220-\u3229",
        "\u3248-\u324f",
        "\u3251-\u325f",
        "\u3280-\u3289",
        "\u32b1-\u32bf",
        "\ua620-\ua629",
        "\ua6e6-\ua6ef",
        "\ua830-\ua835",
        "\ua8d0-\ua8d9",
        "\ua900-\ua909",
        "\ua9d0-\ua9d9",
        "\ua9f0-\ua9f9",
        "\uaa50-\uaa59",
        "\uabf0-\uabf9",
        "\uff10-\uff19",
    ]
    alpha = [
        "\u0041-\u005a",
        "\u0061-\u007a",
        "\u00aa-\u00aa",
        "\u00ba-\u00ba",
        "\u00c0-\u00d6",
        "\u00d8-\u00f6",
        "\u00f8-\u02b8",
        "\u02e0-\u02e4",
        "\u1d00-\u1d25",
        "\u1d2c-\u1d5c",
        "\u1d62-\u1d65",
        "\u1d6b-\u1d77",
        "\u1d79-\u1dbe",
        "\u1e00-\u1eff",
        "\u2071-\u2071",
        "\u207f-\u207f",
        "\u2090-\u209c",
        "\u212a-\u212b",
        "\u2132-\u2132",
        "\u214e-\u214e",
        "\u2160-\u2188",
        "\u2c60-\u2c7f",
        "\ua722-\ua787",
        "\ua78b-\ua7ae",
        "\ua7b0-\ua7b7",
        "\ua7f7-\ua7ff",
        "\uab30-\uab5a",
        "\uab5c-\uab64",
        "\ufb00-\ufb06",
        "\uff21-\uff3a",
        "\uff41-\uff5a",
    ]
    hiragana = [
        "\u3041-\u3096",
        "\u309d-\u309f",
    ]
    katakana = [
        "\u30a1-\u30fa",
        "\u30fd-\u30ff",
        "\u31f0-\u31ff",
        "\u32d0-\u32fe",
        "\u3300-\u3357",
        "\uff66-\uff6f",
        "\uff71-\uff9d",
    ]
    if re.compile('[{}]'.format(''.join(num))).match(s):
        s = 'number'
    elif re.compile('[{}]'.format(''.join(alpha))).match(s):
        s = 'alphabet'
    elif re.compile('[{}]'.format(''.join(hiragana))).match(s):
        s = 'ひらがな'
    elif re.compile('[{}]'.format(''.join(katakana))).match(s):
        s = 'カタカナ'
    return s
