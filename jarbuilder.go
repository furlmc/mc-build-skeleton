package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var archiveName, source, assets string
var version, mcversion string

func main() {
	// Flags
	flag.StringVar(&archiveName, "filename", "build.zip", "name of output file")
	flag.StringVar(&source, "ns", "furl", "source package")
	flag.StringVar(&assets, "assets", "mcmod.info", "list of assets")
	flag.StringVar(&version, "v", "1.0.0", "mod version")
	flag.StringVar(&mcversion, "mc", "1.6.2", "minecraft version")
	flag.Parse()

	// Called after reobfuscate
	// mcp = forge/mcp
	// jar path = reobf/minecraft/* + src/{assets...}
	files := make(map[string]*os.File)

	reobfPath := "forge/mcp/reobf/minecraft"
	assetsPath := "src"

	reobf, err := os.Stat(reobfPath)
	if err != nil {
		panic(err)
	}

	if !reobf.IsDir() {
		panic("reobf folder not found")
	}

	var traverse func(fi os.FileInfo, dirPath, zipPath string)
	traverse = func(fi os.FileInfo, dirPath, zipPath string) {
		fd, err := os.Open(dirPath)
		if err != nil {
			panic(err)
		}
		fis, err := fd.Readdir(0)
		if err != nil {
			panic(err)
		}
		for _, fi := range fis {
			if fi.IsDir() {
				traverse(fi,
					path.Join(dirPath, fi.Name()),
					path.Join(zipPath, fi.Name()),
				)
			} else {
				fileName := path.Join(dirPath, fi.Name())
				zipName := path.Join(zipPath, fi.Name())
				f, err := os.Open(fileName)
				if err != nil {
					panic(err)
				}
				files[zipName] = f
			}
		}
	}
	traverse(reobf, reobfPath, "")

	for _, asset := range strings.Split(assets, " ") {
		assetPath := path.Join(assetsPath, asset)
		stat, err := os.Stat(assetPath)
		if err != nil {
			panic(err)
		}
		if stat.IsDir() {
			traverse(stat, assetPath, asset)
		} else {
			f, err := os.Open(assetPath)
			if err != nil {
				panic(err)
			}
			files[asset] = f
		}
	}

	// add manifest if it exists
	{
		manifest, err := os.Open(path.Join(assetsPath, "MANIFEST.MF"))
		if err == nil {
			fmt.Println("[-]adding manifest to jar")
			files["META-INF/MANIFEST.MF"] = manifest
		}
	}

	// ... zip files
	outZip, err := os.Create(archiveName)
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(outZip)
	for zipPath, file := range files {
		of, err := zipWriter.Create(zipPath)
		if err != nil {
			panic(err)
		}
		if zipPath == "mcmod.info" {
			bs, err := ioutil.ReadAll(file)
			mcmod := strings.Replace(string(bs), "@VERSION@", version, -1)
			mcmod = strings.Replace(mcmod, "@MCVERSION@", mcversion, -1)
			_, err = of.Write([]byte(mcmod))
			if err != nil {
				panic(err)
			}
		} else {
			_, err = io.Copy(of, file)
			if err != nil {
				panic(err)
			}
		}
	}
	err = zipWriter.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[=]%d files written to %s\n", len(files), archiveName)
}
