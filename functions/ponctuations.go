package functions

import "strings"

func Ponctuations(s []string) []string {
    var res []string
    par := ""

    // Première étape : regrouper ponctuations entre elles
    for i := 0; i < len(s); i++ {
        if CheckPonc(s[i]) {
            par += s[i]
        } else {
            if par != "" {
                res = append(res, par)
                par = ""
            }
            res = append(res, s[i])
        }
    }
    if par != "" {
        res = append(res, par)
    }

    // Deuxième étape : concaténer ponctuations au mot précédent
    var r []string
    for _, tok := range res {
        if len(r) > 0 && strings.ContainsAny(tok, ".,!?:;") {
            r[len(r)-1] += tok
        } else {
            r = append(r, tok)
        }
    }

    return r
}
