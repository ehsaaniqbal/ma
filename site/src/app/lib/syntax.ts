import Prism, { languages } from "prismjs";

export const maSyntax = languages.extend("clike", {
  string: {
    pattern: /(["`])(?:\\[\s\S]|(?!\1)[^\\])*\1/,
    greedy: true,
  },
  keyword: /\b(?:ma|initiative|super dream|regular|god bless you)\b/,
  boolean: /\b(?:true|false)\b/,
  number: /(?:(?:\b\d+(?:\.\d*)?|\B\.\d+)(?:e[-+]?\d+)?)i?/i,
  operator:
    /[*/%^!=]=?|~|\+[=+]?|-[=-]?|\|[=|]?|&(?:=|&|\^=?)?|>(?:>=?|=)?|<(?:<=?|=|-)?|:=|\.\.\./,
});

Prism.languages.ma = maSyntax;
