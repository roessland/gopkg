package mathutil

func GCD(r0, r1 int64) int64 {
    if r1 > r0 {
        r0, r1 = r1, r0
    }

    for r1 != 0 {
        r0, r1 = r1, r0 % r1
    }

    return r0
}
