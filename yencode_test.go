package yencode

import (
    "bytes"
    "io"
    "io/ioutil"
    "os"
    "testing"
)

func TestYencodeText(t *testing.T) {
    // open and read the input file
    infile, err := os.Open("test1.in")
    if err != nil {
        t.Fatalf("couldn't open test1.in: %s", err)
    }

    inbuf, err := ioutil.ReadAll(infile)
    if err != nil {
        t.Fatalf("couldn't read test1.in: %s", err)
    }

    // open and read the yencode output file
    testfile, err := os.Open("test1.ync")
    if err != nil {
        t.Fatalf("couldn't open test1.ync: %s", err)
    }

    testbuf, err := ioutil.ReadAll(testfile)
    if err != nil {
        t.Fatalf("couldn't read test1.ync: %s", err)
    }

    // generate a dodgy message
    out := new(bytes.Buffer)

    io.WriteString(out, "=ybegin line=128 size=858 name=test1.in\r\n")

    err = Encode(inbuf, out)
    if err != nil {
        t.Fatalf("encode error: %s", err)
    }

    io.WriteString(out, "=yend size=858 crc32=3274F3F7\r\n")

    // compare
    if bytes.Compare(testbuf, out.Bytes()) != 0 {
        t.Fatalf("data mismatch")
    }
}

func TestYencodeBinary(t *testing.T) {
    // open and read the input file
    infile, err := os.Open("test2.in")
    if err != nil {
        t.Fatalf("couldn't open test2.in: %s", err)
    }

    inbuf, err := ioutil.ReadAll(infile)
    if err != nil {
        t.Fatalf("couldn't read test2.in: %s", err)
    }

    // open and read the yencode output file
    testfile, err := os.Open("test2.ync")
    if err != nil {
        t.Fatalf("couldn't open test2.ync: %s", err)
    }

    testbuf, err := ioutil.ReadAll(testfile)
    if err != nil {
        t.Fatalf("couldn't read test2.ync: %s", err)
    }

    // generate a dodgy message
    out := new(bytes.Buffer)

    io.WriteString(out, "=ybegin line=128 size=76800 name=test2.in\r\n")

    err = Encode(inbuf, out)
    if err != nil {
        t.Fatalf("encode error: %s", err)
    }

    io.WriteString(out, "=yend size=76800 crc32=12AAC2CF\r\n")

    // compare
    if bytes.Compare(testbuf, out.Bytes()) != 0 {
        t.Fatalf("data mismatch")
    }
}

func bench(b *testing.B, n int) {
    inbuf := makeInBuf(n)
    out := new(bytes.Buffer)
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        out.Reset()
        _ = Encode(inbuf, out)
    }

    b.SetBytes(int64(len(inbuf)))
}

func BenchmarkEncode10(b *testing.B) {
    bench(b, 10)
}

func BenchmarkEncode100(b *testing.B) {
    bench(b, 100)
}

func BenchmarkEncode1000(b *testing.B) {
    bench(b, 1000)
}

func makeInBuf(length int) []byte {
    chars := length * 256 * 132
    pos := 0

    in := make([]byte, chars)
    for i := 0; i < length; i++ {
        for j := 0; j < 256; j++ {
            for k := 0; k < 132; k++ {
                in[pos] = byte(j)
                pos++
            }
        }
    }

    return in
}
