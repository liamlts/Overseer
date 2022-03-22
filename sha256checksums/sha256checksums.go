package sha256checksums

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func GenSums() []string {
	var sums []string
	var maindir string = "/bin"
	dir, err := os.ReadDir(maindir)
	if err != nil {
		log.Fatal(err)
	}
	for _, curFile := range dir {
		binary := maindir + "/" + curFile.Name()
		checkDir, err := os.Stat(binary)
		if err != nil {
			log.Fatal(err)
		}

		if !checkDir.IsDir() {
			bin, err := os.Open(binary)
			if err != nil {
				log.Fatal(err)
			}
			hasher := sha256.New()

			if _, err := io.Copy(hasher, bin); err != nil {
				log.Fatal(err)
			}
			encodeString := hex.EncodeToString(hasher.Sum(nil))
			sums = append(sums, binary+" : "+encodeString)
		}

	}
	return sums
}

func exists() bool {
	_, err := os.Stat("checksums/sha256.data")
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func GenSumFile(sums []string) {
	if !exists() {
		binarys := GenSums()
		err := os.Mkdir("checksums", 0744)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.OpenFile("checksums/sha256.data",
			os.O_CREATE|os.O_WRONLY, 0744)

		defer file.Close()

		for i := range binarys {
			if _, err := file.WriteString(binarys[i] + "\n"); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func CheckSums() {
	sumfile, err := os.Open("checksums/sha256.data")
	if err != nil {
		log.Fatal(err)
	}

	defer sumfile.Close()
	var ogsums []string

	scanner := bufio.NewScanner(sumfile)

	for scanner.Scan() {
		ogsums = append(ogsums, scanner.Text())
	}

	cursums := GenSums()

	for i := range ogsums {
		if ogsums[i] != cursums[i] {
			fmt.Println("Possbile file manipulation on ", cursums[i])
		}
	}
}
