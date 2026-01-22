package main

import (
	"fmt"
	"io"
	"os"

	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

// ref https://github.com/ipfs/go-cid

const bufferSize = 64 * 1024 // 64KB

// and return a newly constructed Cid with the resulting multihash.
func SumStream(p cid.Prefix, r io.Reader) (cid.Cid, error) {
	length := p.MhLength
	if p.MhType == mh.IDENTITY {
		length = -1
	}

	if p.Version == 0 && (p.MhType != mh.SHA2_256 ||
		(p.MhLength != 32 && p.MhLength != -1)) {

		return cid.Undef, cid.ErrInvalidCid{fmt.Errorf("invalid v0 prefix")}
	}

	hash, err := mh.SumStream(r, p.MhType, length)
	if err != nil {
		return cid.Undef, cid.ErrInvalidCid{err}
	}

	switch p.Version {
	case 0:
		return cid.NewCidV0(hash), nil
	case 1:
		return cid.NewCidV1(p.Codec, hash), nil
	default:
		return cid.Undef, cid.ErrInvalidCid{fmt.Errorf("invalid cid version")}
	}
}

func main() {

	// get arg 1 as file name. and read file content for do pref.Sum(fileContent)
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		fmt.Println("Example: go run main.go test.txt")
		os.Exit(1) // Exit with non-zero code to indicate error
	}

	// Step 2: Get the filename from the first command-line argument (os.Args[1])
	filename := os.Args[1]

	// Step 3: Read file content into []byte
	// fileContent, err := os.ReadFile(filename)
	// if err != nil {
	// 	fmt.Printf("Error reading file '%s': %v\n", filename, err)
	// 	os.Exit(1)
	// }

	// Step 2: 打开文件（流式处理，不立即读取内容）
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	// 确保文件最终关闭（即使发生错误）
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Warning: failed to close file: %v\n", err)
		}
	}()

	// Create a cid manually by specifying the 'prefix' parameters
	pref := cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.Raw),
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	// And then feed it some data
	// c, err := pref.Sum(file)
	c, err := SumStream(pref, file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created CID : ", c)

}
